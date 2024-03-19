package service

import (
	"EduCRM/package/repository"
	"EduCRM/package/service/admin"
	"EduCRM/package/service/auth"
	"EduCRM/package/service/course"
	"EduCRM/package/service/group"
	"EduCRM/package/service/lid"
	"EduCRM/package/service/minio"
	"EduCRM/package/service/user"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
)

type Service struct {
	MinioService  *minio.MinioService
	UserService   *user.UserService
	CourseService *course.CourseService
	GroupService  *group.GroupService
	AuthService   *auth.AuthService
	AdminService  *admin.AdminService
	LidService    *lid.LidService
}

func NewService(repos *repository.Repository, store *store.Store, loggers *logrus_log.Logger) *Service {
	return &Service{
		MinioService:  minio.NewMinioService(store, loggers),
		UserService:   user.NewUserService(repos, store, loggers),
		CourseService: course.NewCourseService(repos, store, loggers),
		GroupService:  group.NewGroupService(repos, store, loggers),
		AuthService:   auth.NewAuthService(repos, store, loggers),
		AdminService:  admin.NewAdminService(repos, store, loggers),
		LidService:    lid.NewLidService(repos, store, loggers),
	}
}
