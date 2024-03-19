package group

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type GroupReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewGroupReaderDB returns a new instance of GroupReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `GroupReaderDB` struct.
func NewGroupReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *GroupReaderDB {
	return &GroupReaderDB{db: db, loggers: loggers}
}
func (repo *GroupReaderDB) GetGroupList(pagination model.Pagination) (groupList []model.Group, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&groupList, GetGroupListQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return groupList, err
		}
		loggers.Error(err)
		return groupList, err
	}
	return groupList, err
}
func (repo *GroupReaderDB) GetGroupByID(id string) (group model.Group, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&group, GetGroupByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return group, err
		}
		loggers.Error(err)
		return group, err
	}
	return group, err
}
func (repo *GroupReaderDB) GetTeacherGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&groupList, GetTeacherGroupListQuery, teacherID, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return groupList, err
		}
		loggers.Error(err)
		return groupList, err
	}
	return groupList, err
}
func (repo *GroupReaderDB) GetStudentGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&groupList, GetStudentGroupListQuery, teacherID, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return groupList, err
		}
		loggers.Error(err)
		return groupList, err
	}
	return groupList, err
}
func (repo *GroupReaderDB) GetCourseGroupList(courseID string, pagination model.Pagination) (groupList []model.Group, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&groupList, GetCourseGroupListQuery, courseID, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return groupList, err
		}
		loggers.Error(err)
		return groupList, err
	}
	return groupList, err
}
func (repo *GroupReaderDB) CheckGroupByID(id string) (groupID string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&groupID, CheckGroupByIDQuery, id)
	if err != nil {
		loggers.Error(err)
		return "", err
	}
	return groupID, nil
}
