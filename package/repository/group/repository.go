package group

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GroupRepo struct {
	GroupReader
	GroupWriter
	GroupEnrollment
}
type GroupReader interface {
	GetGroupList(group model.Pagination) (groupList []model.Group, err error)
	GetGroupByID(id string) (group model.Group, err error)
	GetTeacherGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error)
	GetCourseGroupList(courseID string, pagination model.Pagination) (groupList []model.Group, err error)
	CheckGroupByID(id string) (groupID string, err error)
	GetStudentGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error)
}
type GroupWriter interface {
	CreateGroup(group model.CreateGroup) (err error)
	UpdateGroup(group model.UpdateGroup) (err error)
	DeleteGroup(id string) (err error)
}
type GroupEnrollment interface {
	CreateGroupEnrollment(groupID, learnerID string) (err error)
	DeleteGroupEnrollment(groupID string) (err error)
	DeleteUserAllGroupEnrollment(learnerID string) (err error)
	DeleteUserFromGroup(learnerID, groupID string) (err error)
	GetGroupStudentList(groupID string) (studentList []uuid.UUID, err error)
}

func NewGroupRepo(db *sqlx.DB, loggers *logrus_log.Logger) *GroupRepo {
	return &GroupRepo{
		GroupWriter:     NewGroupWriterDB(db, loggers),
		GroupReader:     NewGroupReaderDB(db, loggers),
		GroupEnrollment: NewGroupEnrollmentDB(db, loggers),
	}
}
