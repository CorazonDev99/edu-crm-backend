package course

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	CourseEndPoint
}

type CourseEndPoint interface {
	CreateCourse(ctx *gin.Context)
	UpdateCourse(ctx *gin.Context)
	DeleteCourse(ctx *gin.Context)
	GetCourseList(ctx *gin.Context)
	GetCourseByID(ctx *gin.Context)
	GetCourseGroupList(ctx *gin.Context)
}

func NewCourseHandler(service *service.Service,
	loggers *logrus_log.Logger) *CourseHandler {
	return &CourseHandler{
		CourseEndPoint: NewCourseEndPointHandler(service, loggers),
	}
}
