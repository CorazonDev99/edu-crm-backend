package course

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"errors"

	"github.com/jmoiron/sqlx"
)

type CourseWriterDB struct {
	db     *sqlx.DB
	logrus *logrus_log.Logger
}

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

// NewCourseWriterDB returns a new instance of CourseWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `CourseWriterDB` struct.
func NewCourseWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *CourseWriterDB {
	return &CourseWriterDB{db: db, logrus: loggers}
}
func (repo *CourseWriterDB) CreateCourse(course model.CreateCourse) (id string, err error) {
	loggers := repo.logrus
	db := repo.db
	row, err := db.Query(CreateCourseQuery, course.Title, course.Description, course.Photo, course.Duration, course.Status, course.Price,
		course.LessonDuration)
	if err != nil {
		loggers.Error(err)
		return "", err
	}
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			loggers.Error(err)
			return "", err
		}
	}
	return id, nil
}
func (repo *CourseWriterDB) UpdateCourse(course model.UpdateCourse) (err error) {
	loggers := repo.logrus
	db := repo.db
	//UpdateCourseQuery  = `-- UPDATE course SET title=$1,description=$2,duration=$3,status=$4,price=$5,lesson_duration=$6,	photo= $7,updated_at = NOW() WHERE id=$8 AND deleted_at IS NULL`

	row, err := db.Exec(UpdateCourseQuery, course.Title, course.Description,
		course.Duration, course.Status, course.Price,
		course.LessonDuration, course.Photo, course.ID)
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
func (repo *CourseWriterDB) DeleteCourse(id string) (teacherID string, err error) {
	loggers := repo.logrus
	db := repo.db
	row, err := db.Exec(DeleteCourseQuery, id)
	if err != nil {
		loggers.Error(err)
		return "", nil
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return "", nil
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return "", ErrorNoRowsAffected
	}
	return teacherID, nil
}
func (repo *CourseWriterDB) CreateTeacherCourseEnrollment(teacherID, courseID string) (err error) {
	loggers := repo.logrus
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
func (repo *CourseWriterDB) DeleteTeacherCourseEnrollment(teacherID, courseID string) (err error) {
	logrus := repo.logrus
	db := repo.db
	row, err := db.Exec(DeleteTeacherCourseEnrollmentQuery, teacherID, courseID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		logrus.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *CourseWriterDB) UpdateTeacherCourseEnrollment(teacherID, newTeacherID, courseID string) (err error) {
	logrus := repo.logrus
	db := repo.db
	row, err := db.Exec(UpdateTeacherCourseEnrollmentQuery, newTeacherID, teacherID, courseID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		logrus.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
