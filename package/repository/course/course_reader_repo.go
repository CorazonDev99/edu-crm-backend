package course

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewCourseReaderDB returns a new instance of CourseReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `CourseReaderDB` struct.
func NewCourseReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *ReaderDB {
	return &ReaderDB{db: db, loggers: loggers}
}
func (repo *ReaderDB) GetCourseList(pagination model.Pagination) (courseList []model.Course, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&courseList, GetCourseListQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return courseList, err
		}
		loggers.Error(err)
		return courseList, err
	}
	return courseList, err
}
func (repo *ReaderDB) GetCourseByID(id string) (course model.Course, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&course, GetCourseByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return course, err
		}
		loggers.Error(err)
		return course, err
	}
	return course, err
}
func (repo *ReaderDB) CheckCourseByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckCourseByID, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}
func (repo *ReaderDB) GetTeacherCourseList(teacherID string, pagination model.Pagination) (courseList []model.Course, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&courseList, GetTeacherCourseListQuery, teacherID, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return courseList, err
		}
		loggers.Error(err)
		return courseList, err
	}
	return courseList, err
}
func (repo *ReaderDB) GetStudentCourseList(teacherID string, pagination model.Pagination) (courseList []model.Course, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&courseList, GetStudentCourseListQuery, teacherID, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return courseList, err
		}
		loggers.Error(err)
		return courseList, err
	}
	return courseList, err
}
