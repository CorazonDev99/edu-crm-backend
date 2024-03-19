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
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type UserReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewUserReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *UserReaderService {
	return &UserReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *UserReaderService) GetUserList(role string, pagination model.Pagination) (userList []model.User,
	err error) {
	if role != "all" {
		err = validation.UUIDValidation(role)
		if err != nil {
			return nil, response.ServiceError(err, codes.InvalidArgument)
		}
	}
	userList, err = s.repo.UserRepository.GetUserList(role, pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	for i, user := range userList {
		if len(user.Photo) != 0 {
			err := s.minio.ObjectStore.ObjectExists(user.Photo)
			if err != nil {
				s.loggers.Error(err)
			} else {
				image, err := s.minio.FileLinkStore.GetImageUrl(user.Photo)
				if err != nil {
					s.loggers.Error(err)
				}
				userList[i].PhotoLink = image
			}
		}
		if len(user.ExtraDataByte) != 0 {
			err = json.Unmarshal(user.ExtraDataByte, &userList[i].ExtraDataJSON)
			if err != nil {
				s.loggers.Error(err)
			}
		}
		if len(userList[i].RoleID.String()) != 0 {
			userList[i].Role, err = s.repo.AuthRepository.GetRoleTitleByID(
				userList[i].RoleID)
			if err != nil {
				s.loggers.Error(err)
			}
		}
	}
	return userList, nil
}
func (s *UserReaderService) GetUserByID(id string) (user model.User, err error) {
	user, err = s.repo.UserRepository.GetUserByID(id)
	if err != nil {
		return user, response.ServiceError(err, codes.Internal)
	}
	if len(user.Photo) != 0 {
		err = s.minio.ObjectStore.ObjectExists(user.Photo)
		if err != nil {
			s.loggers.Error(err)
			return user, nil
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(user.Photo)
		if err != nil {
			s.loggers.Error(err)
			return user, nil
		}
		user.PhotoLink = image
	}
	if len(user.ExtraDataByte) != 0 {
		err = json.Unmarshal(user.ExtraDataByte, &user.ExtraDataJSON)
		if err != nil {
			s.loggers.Error(err)
		}
	}
	if user.RoleID != uuid.Nil {
		user.Role, err = s.repo.AuthRepository.GetRoleTitleByID(
			user.RoleID)
		if err != nil {
			s.loggers.Error(err)
		}
	}
	return user, nil
}
func (s *UserReaderService) SignInUser(user model.SignInUser) (id, role uuid.UUID, roleTitle string, err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		s.loggers.Error(err)
		return id, role, roleTitle, response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	id, role, err = s.repo.UserRepository.SignInUser(user)
	if err != nil {
		s.loggers.Error(err)
		return id, role, roleTitle, response.ServiceError(err, codes.NotFound)
	}
	if role != uuid.Nil {
		roleTitle, err = s.repo.AuthRepository.GetRoleTitleByID(role)
		if err != nil {
			s.loggers.Error(err)
		}
	}
	return id, role, roleTitle, nil
}
