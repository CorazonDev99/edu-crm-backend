package group

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type GroupReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewGroupReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *GroupReaderService {
	return &GroupReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *GroupReaderService) GetGroupList(pagination model.Pagination) (
	groupList []model.Group, err error) {
	groupList, err = s.repo.GroupRepository.GetGroupList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return groupList, nil
}
func (s *GroupReaderService) GetGroupByID(id string) (group model.Group, err error) {
	group, err = s.repo.GroupRepository.GetGroupByID(id)
	if err != nil {
		return group, response.ServiceError(err, codes.Internal)
	}
	return group, nil
}

func (s *GroupReaderService) GetGroupStudentList(groupID string) (studentList []model.User, err error) {
	studentIDList, err := s.repo.GroupRepository.GetGroupStudentList(groupID)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	studentList, err = s.repo.UserRepository.GetGroupStudentList(studentIDList)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return studentList, nil
}
