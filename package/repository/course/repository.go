package course

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"

	"github.com/jmoiron/sqlx"
)

type CourseRepo struct {
	CourseReader
	CourseWriter
	TeacherCourseEnrollmentWriter
}
type CourseReader interface {
	GetCourseList(course model.Pagination) (courseList []model.Course, err error)
	GetCourseByID(id string) (course model.Course, err error)
	CheckCourseByID(id string) (err error)
	GetTeacherCourseList(teacherID string, pagination model.Pagination) (courseList []model.Course, err error)
	GetStudentCourseList(studentID string, pagination model.Pagination) (courseList []model.Course, err error)
}
type CourseWriter interface {
	CreateCourse(course model.CreateCourse) (id string, err error)
	UpdateCourse(course model.UpdateCourse) (err error)
	DeleteCourse(id string) (teacherID string, err error)
}
type TeacherCourseEnrollmentWriter interface {
	CreateTeacherCourseEnrollment(teacherID, courseID string) (err error)
	DeleteTeacherCourseEnrollment(teacherID, courseID string) (err error)
	UpdateTeacherCourseEnrollment(teacherID, newTeacherID, courseID string) (err error)
}

func NewCourseRepo(db *sqlx.DB, logrus *logrus_log.Logger) *CourseRepo {
	return &CourseRepo{
		CourseWriter:                  NewCourseWriterDB(db, logrus),
		CourseReader:                  NewCourseReaderDB(db, logrus),
		TeacherCourseEnrollmentWriter: NewTeacherCourseEnrollmentWriterDB(db, logrus),
	}
}
