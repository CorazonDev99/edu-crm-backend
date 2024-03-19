package group

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type GroupWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewGroupWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *GroupWriterService {
	return &GroupWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *GroupWriterService) CreateGroup(group model.CreateGroup) (err error) {
	err = validation.ValidationStructTag(s.loggers, group)
	if err != nil {
		return err
	}
	err = s.repo.CourseRepository.CheckCourseByID(group.CourseID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.UserRepository.CheckUserByID(group.TeacherID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.AdminRepository.CheckRoomByID(group.RoomID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.GroupRepository.CreateGroup(group)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *GroupWriterService) UpdateGroup(group model.UpdateGroup) (err error) {
	err = validation.ValidationStructTag(s.loggers, group)
	if err != nil {
		return err
	}
	err = s.repo.CourseRepository.CheckCourseByID(group.CourseID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.UserRepository.CheckUserByID(group.TeacherID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.AdminRepository.CheckRoomByID(group.RoomID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.GroupRepository.UpdateGroup(group)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *GroupWriterService) DeleteGroup(id string) (err error) {
	err = s.repo.GroupRepository.DeleteGroup(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *GroupWriterService) DeleteUserFromGroup(learnerID,
	groupID string) (err error) {
	err = s.repo.UserRepository.CheckUserByID(learnerID)
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.GroupRepository.DeleteUserFromGroup(learnerID, groupID)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
