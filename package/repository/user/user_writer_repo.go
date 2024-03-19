package user

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

// NewUserWriterDB returns a new instance of UserWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `UserWriterDB` struct.
func NewUserWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *UserWriterDB {
	return &UserWriterDB{db: db, loggers: loggers}
}
func (repo *UserWriterDB) CreateUser(user model.CreateUser) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateUserQuery, user.FullName, user.BirthdayDate,
		user.AddedDate, user.PhoneNumber, user.RoleID, user.Password,
		user.ExtraDataJSON, user.Photo)
	if err != nil {
		loggers.Error(err)
		return id, err
	}
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			loggers.Error(err)
			return id, err
		}
	}
	return id, nil
}
func (repo *UserWriterDB) UpdateUser(user model.UpdateUser) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateUserQuery, user.FullName, user.BirthdayDate,
		user.AddedDate, user.PhoneNumber, user.RoleID, user.ExtraDataJSON, user.Photo, user.ID)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *UserWriterDB) DeleteUser(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteUserQuery, id)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *UserWriterDB) UpdateUserPassword(id, password string) error {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateUserPasswordQuery, password, id)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
