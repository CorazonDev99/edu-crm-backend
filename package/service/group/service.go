package group

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
)

type GroupService struct {
	GroupReaderMethod
	GroupWriterMethod
}

type GroupReaderMethod interface {
	GetGroupList(group model.Pagination) (groupList []model.Group, err error)
	GetGroupByID(id string) (group model.Group, err error)
	GetGroupStudentList(groupID string) (studentList []model.User, err error)
}
type GroupWriterMethod interface {
	CreateGroup(group model.CreateGroup) (err error)
	UpdateGroup(group model.UpdateGroup) (err error)
	DeleteGroup(id string) (err error)
	DeleteUserFromGroup(learnerID, groupID string) (err error)
}

func NewGroupService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *GroupService {
	return &GroupService{
		GroupReaderMethod: NewGroupReaderService(repo, minio, loggers),
		GroupWriterMethod: NewGroupWriterService(repo, minio, loggers),
	}
}
