package group

import (
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GroupEnrollmentDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewGroupEnrollmentDB returns a new instance of GroupEnrollmentDB.
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `GroupEnrollmentDB` struct.
func NewGroupEnrollmentDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *GroupEnrollmentDB {
	return &GroupEnrollmentDB{db: db, loggers: loggers}
}

func (repo *GroupEnrollmentDB) CreateGroupEnrollment(groupID, learnerID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(CreateGroupEnrollmentQuery, groupID, learnerID)
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
func (repo *GroupEnrollmentDB) DeleteGroupEnrollment(groupID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteGroupEnrollmentQuery, groupID)
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
func (repo *GroupEnrollmentDB) DeleteUserAllGroupEnrollment(
	teacherID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteUserAllGroupEnrollmentQuery, teacherID)
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
func (repo *GroupEnrollmentDB) DeleteUserFromGroup(learnerID, groupID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteUserFromGroupQuery, learnerID, groupID)
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

func (repo *GroupEnrollmentDB) GetGroupStudentList(groupID string) (studentList []uuid.UUID, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&studentList, GetStudentGroupListQuery, groupID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return studentList, err
		}
		loggers.Error(err)
		return studentList, err
	}
	return studentList, err
}
