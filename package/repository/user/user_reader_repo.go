package user

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	errorSignInUser = errors.New("  Username or Password is Incorrect")
)

type UserReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

func NewUserReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *UserReaderDB {
	return &UserReaderDB{db: db, loggers: loggers}
}
func (repo *UserReaderDB) GetUserList(role string,
	pagination model.Pagination) (userList []model.User,
	err error) {
	loggers := repo.loggers
	db := repo.db
	if role == "all" {
		err = db.Select(&userList, GetUserListAllQuery, pagination.Limit, pagination.Offset)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return userList, err
			}
			loggers.Error(err)
			return userList, err
		}
	} else {
		err = db.Select(&userList, GetUserListByRoleQuery, role, pagination.Limit, pagination.Offset)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return userList, err
			}
			loggers.Error(err)
			return userList, err
		}
	}
	return userList, err
}
func (repo *UserReaderDB) GetUserByID(id string) (user model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetUserByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}
func (repo *UserReaderDB) SignInUser(user model.SignInUser) (id, role uuid.UUID, err error) {
	var result model.SignInUserResponse
	loggers := repo.loggers
	err = repo.db.Get(&result, SignInUserQuery, user.PhoneNumber, user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, role, err
		}
		loggers.Error(err)
		return id, role, errorSignInUser
	}
	return result.ID, result.Role, nil
}
func (repo *UserReaderDB) CheckUserByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckUserByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	if courseID != id {
		return errors.New("invalid course id")
	}
	return nil
}
func (repo *UserReaderDB) GetGroupStudentList(studentIDList []uuid.UUID) (studentList []model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&studentList, GetGroupStudentListQuery, studentIDList)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return studentList, err
		}
		loggers.Error(err)
		return studentList, err
	}
	return studentList, nil
}
