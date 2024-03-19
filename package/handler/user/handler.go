package user

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserEndPoint
	TeacherEndPoint
	StudentEndPoint
}

type UserEndPoint interface {
	SignInUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetUserList(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUserPassword(ctx *gin.Context)
	GetUserMe(ctx *gin.Context)
	CreateSuperAdmin(ctx *gin.Context)
}
type TeacherEndPoint interface {
	GetTeacherGroupList(ctx *gin.Context)
	GetTeacherCourseList(ctx *gin.Context)
}
type StudentEndPoint interface {
	GetStudentGroupList(ctx *gin.Context)
	GetStudentCourseList(ctx *gin.Context)
}

func NewUserHandler(service *service.Service,
	loggers *logrus_log.Logger) *UserHandler {
	return &UserHandler{
		UserEndPoint:    NewUserEndPointHandler(service, loggers),
		TeacherEndPoint: NewTeacherEndPointHandler(service, loggers),
		StudentEndPoint: NewStudentEndPointHandler(service, loggers),
	}
}
