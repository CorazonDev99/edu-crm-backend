package user

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	UserReader
	UserWriter
}
type UserReader interface {
	GetUserList(role string, user model.Pagination) (userList []model.User, err error)
	GetUserByID(id string) (user model.User, err error)
	SignInUser(user model.SignInUser) (id, role uuid.UUID, err error)
	CheckUserByID(id string) (err error)
	GetGroupStudentList(studentIDList []uuid.UUID) (studentList []model.User, err error)
}
type UserWriter interface {
	CreateUser(user model.CreateUser) (id uuid.UUID, err error)
	UpdateUser(user model.UpdateUser) (err error)
	DeleteUser(id string) (err error)
	UpdateUserPassword(id, password string) error
}

func NewUserRepo(db *sqlx.DB, loggers *logrus_log.Logger) *UserRepo {
	return &UserRepo{
		UserWriter: NewUserWriterDB(db, loggers),
		UserReader: NewUserReaderDB(db, loggers),
	}
}
