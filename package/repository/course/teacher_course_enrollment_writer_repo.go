package course

import (
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type TeacherCourseEnrollmentWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewTeacherCourseEnrollmentWriterDB returns a new instance of
// TeacherCourseEnrollmentWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `TeacherCourseEnrollmentWriterDB` struct.
func NewTeacherCourseEnrollmentWriterDB(db *sqlx.DB,
	logrus *logrus_log.Logger) *TeacherCourseEnrollmentWriterDB {
	return &TeacherCourseEnrollmentWriterDB{db: db, loggers: logrus}
}

func (repo *TeacherCourseEnrollmentWriterDB) CreateTeacherCourseEnrollment(
	teacherID,
	courseID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(CreateTeacherCourseEnrollmentQuery, teacherID, courseID)
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
func (repo *TeacherCourseEnrollmentWriterDB) DeleteTeacherCourseEnrollment(teacherID, courseID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteTeacherCourseEnrollmentQuery, teacherID, courseID)
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
func (repo *TeacherCourseEnrollmentWriterDB) UpdateTeacherCourseEnrollment(teacherID, newTeacherID, courseID string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateTeacherCourseEnrollmentQuery, newTeacherID, teacherID, courseID)
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
