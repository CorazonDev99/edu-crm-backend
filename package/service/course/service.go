package course

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
)

type CourseService struct {
	CourseReaderMethod
	CourseWriterMethod
}

type CourseReaderMethod interface {
	GetCourseList(course model.Pagination) (courseList []model.Course, err error)
	GetCourseByID(id string) (course model.Course, err error)
	GetCourseGroupList(courseID string, pagination model.Pagination) (groupList []model.Group, err error)
}

type CourseWriterMethod interface {
	CreateCourse(course model.CreateCourse) (err error)
	UpdateCourse(course model.UpdateCourse) (err error)
	DeleteCourse(id string) (err error)
}

func NewCourseService(repo *repository.Repository, minio *store.Store,
	logrus *logrus_log.Logger) *CourseService {
	return &CourseService{
		CourseReaderMethod: NewCourseReaderService(repo, minio, logrus),
		CourseWriterMethod: NewCourseWriterService(repo, minio, logrus),
	}
}
