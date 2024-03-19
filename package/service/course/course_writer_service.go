package course

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type CourseWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewCourseWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *CourseWriterService {
	return &CourseWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *CourseWriterService) CreateCourse(course model.CreateCourse) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, course)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	_, err = s.repo.CourseRepository.CreateCourse(course)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *CourseWriterService) UpdateCourse(course model.UpdateCourse) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, course)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.CourseRepository.UpdateCourse(course)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *CourseWriterService) DeleteCourse(id string) (err error) {
	_, err = s.repo.CourseRepository.DeleteCourse(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
