package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/hash"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

var (
	groupID    = "groupID"
	newGroupID = "newGroupID"
	// oldGroupID   = "oldGroupID"
	// teacherID    = "teacherID"
	// newTeacherID = "newTeacherID"
	// oldTeacherID = "oldTeacherID"
)

type UserWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewUserWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *UserWriterService {
	return &UserWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *UserWriterService) CreateUser(user model.CreateUser) (id uuid.UUID,
	err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		return id, response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	user.ExtraDataJSON, err = json.Marshal(user.ExtraData)
	if err != nil {
		return id, response.ServiceError(err, codes.Internal)
	}
	err = s.repo.AuthRepository.CheckRoleByID(user.RoleID)
	if err != nil {
		err := errors.New("role is not exist")
		return id, response.ServiceError(err, codes.NotFound)
	}
	userID, err := s.repo.UserRepository.CreateUser(user)
	if err != nil {
		return id, response.ServiceError(err, codes.Internal)
	}
	groupID, existGroupKey := user.ExtraData[groupID].(string)
	if existGroupKey {
		if len(groupID) != 0 {
			_, err = s.repo.GroupRepository.CheckGroupByID(groupID)
			if err != nil {
				return id, response.ServiceError(err, codes.InvalidArgument)
			}
			err := s.repo.GroupRepository.CreateGroupEnrollment(groupID, userID.String())
			if err != nil {
				return id, err
			}
			delete(user.ExtraData, groupID)
		}
	}
	return userID, nil
}
func (s *UserWriterService) UpdateUser(user model.UpdateUser) (err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	user.ExtraDataJSON, err = json.Marshal(user.ExtraData)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	err = s.repo.AuthRepository.CheckRoleByID(user.RoleID)
	if err != nil {
		err := errors.New("role is not exist")
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.UserRepository.UpdateUser(user)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	groupID, existGroupKey := user.ExtraData[newGroupID].(string)
	if existGroupKey {
		if len(groupID) != 0 {
			_, err = s.repo.GroupRepository.CheckGroupByID(groupID)
			if err != nil {
				return response.ServiceError(err, codes.InvalidArgument)
			}
			err := s.repo.GroupRepository.CreateGroupEnrollment(groupID, user.ID.String())
			if err != nil {
				return err
			}
			delete(user.ExtraData, newGroupID)
		}
	}
	return nil
}
func (s *UserWriterService) DeleteUser(id string) (err error) {
	err = s.repo.UserRepository.DeleteUser(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	// err = s.repo.GroupRepository.DeleteUserAllGroupEnrollment(id)
	// if err != nil {
	// 	return ServiceErrorHandler(err, codes.Internal)
	// }
	//delete from auth_account
	return nil
}

func (s *UserWriterService) UpdateUserPassword(id, password string) error {
	err := validation.ValidatePassword(password, "password")
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	password = hash.GeneratePasswordHash(password)
	err = s.repo.UserRepository.UpdateUserPassword(id, password)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
