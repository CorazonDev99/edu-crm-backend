package course

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type CourseReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewCourseReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *CourseReaderService {
	return &CourseReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *CourseReaderService) GetCourseList(pagination model.Pagination) (
	courseList []model.Course, err error) {
	courseList, err = s.repo.CourseRepository.GetCourseList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	for i, user := range courseList {
		if len(user.Photo) != 0 {
			err := s.minio.ObjectStore.ObjectExists(user.Photo)
			if err != nil {
				s.loggers.Error(err)
				courseList[i].PhotoLink = ""
			} else {
				image, err := s.minio.FileLinkStore.GetImageUrl(user.Photo)
				if err != nil {
					s.loggers.Error(err)
					courseList[i].PhotoLink = ""
				}
				courseList[i].PhotoLink = image
			}
		}
	}
	return courseList, nil
}
func (s *CourseReaderService) GetCourseByID(id string) (course model.Course,
	err error) {
	course, err = s.repo.CourseRepository.GetCourseByID(id)
	if err != nil {
		return course, response.ServiceError(err, codes.Internal)
	}
	if len(course.Photo) != 0 {
		err = s.minio.ObjectStore.ObjectExists(course.Photo)
		if err != nil {
			s.loggers.Error(err)
			course.PhotoLink = ""
			return course, nil
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(course.Photo)
		if err != nil {
			s.loggers.Error(err)
			course.PhotoLink = ""
			return course, nil
		}
		course.PhotoLink = image
	}
	return course, nil
}
func (s *CourseReaderService) GetCourseGroupList(courseID string, pagination model.Pagination) (groupList []model.Group, err error) {
	groupList, err = s.repo.GroupRepository.GetCourseGroupList(courseID, pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return groupList, nil
}
