package repository

import (
	"EduCRM/package/repository/admin"
	"EduCRM/package/repository/auth"
	"EduCRM/package/repository/course"
	"EduCRM/package/repository/group"
	"EduCRM/package/repository/lid"
	"EduCRM/package/repository/user"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	UserRepository   *user.UserRepo
	CourseRepository *course.CourseRepo
	GroupRepository  *group.GroupRepo
	AuthRepository   *auth.AuthRepo
	AdminRepository  *admin.Repo
	LidRepository    *lid.Repo
}

func NewRepository(db *sqlx.DB, loggers *logrus_log.Logger) *Repository {
	return &Repository{
		UserRepository:   user.NewUserRepo(db, loggers),
		CourseRepository: course.NewCourseRepo(db, loggers),
		GroupRepository:  group.NewGroupRepo(db, loggers),
		AuthRepository:   auth.NewAuthRepo(db, loggers),
		AdminRepository:  admin.NewAdminRepo(db, loggers),
		LidRepository:    lid.NewLidRepo(db, loggers),
	}
}
