package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type UserEnrollmentService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewUserEnrollmentService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *UserEnrollmentService {
	return &UserEnrollmentService{repo: repo, minio: minio, loggers: loggers}
}

func (s *UserEnrollmentService) GetTeacherCourseList(teacherID string, pagination model.Pagination) (courseList []model.Course, err error) {
	courseList, err = s.repo.CourseRepository.GetTeacherCourseList(teacherID, pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return courseList, nil
}
func (s *UserEnrollmentService) GetTeacherGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error) {
	groupList, err = s.repo.GroupRepository.GetTeacherGroupList(teacherID,
		pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return groupList, nil
}

func (s *UserEnrollmentService) GetStudentCourseList(teacherID string, pagination model.Pagination) (courseList []model.Course, err error) {
	courseList, err = s.repo.CourseRepository.GetStudentCourseList(teacherID, pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return courseList, nil
}
func (s *UserEnrollmentService) GetStudentGroupList(teacherID string, pagination model.Pagination) (groupList []model.Group, err error) {
	groupList, err = s.repo.GroupRepository.GetStudentGroupList(teacherID,
		pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return groupList, nil
}
