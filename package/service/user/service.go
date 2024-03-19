package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"
)

type UserService struct {
	User
	UserEnrollment
}
type User struct {
	UserReader
	UserWriter
}
type UserWriter interface {
	CreateUser(user model.CreateUser) (id uuid.UUID, err error)
	UpdateUser(user model.UpdateUser) (err error)
	DeleteUser(id string) (err error)
	UpdateUserPassword(id, password string) error
}
type UserReader interface {
	GetUserList(role string, user model.Pagination) (userList []model.User,
		err error)
	GetUserByID(id string) (user model.User, err error)
	SignInUser(user model.SignInUser) (id, role uuid.UUID, roleTitle string, err error)
}
type UserEnrollment interface {
	GetTeacherCourseList(teacherID string, pagination model.Pagination) (courseList []model.Course, err error)
	GetTeacherGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error)
	GetStudentCourseList(studentID string, pagination model.Pagination) (courseList []model.Course, err error)
	GetStudentGroupList(studentID string, pagination model.Pagination) (groupList []model.Group, err error)
}

func NewUserService(repos *repository.Repository, store *store.Store,
	loggers *logrus_log.Logger) *UserService {
	return &UserService{
		User: User{
			UserReader: NewUserReaderService(repos, store, loggers),
			UserWriter: NewUserWriterService(repos, store, loggers),
		},
		UserEnrollment: NewUserEnrollmentService(repos, store, loggers),
	}
}
