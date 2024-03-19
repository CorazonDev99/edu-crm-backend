package group

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	GroupEndPoint
}

type GroupEndPoint interface {
	CreateGroup(ctx *gin.Context)
	UpdateGroup(ctx *gin.Context)
	DeleteGroup(ctx *gin.Context)
	GetGroupList(ctx *gin.Context)
	GetGroupByID(ctx *gin.Context)
	DeleteUserFromGroup(ctx *gin.Context)
	GetGroupStudentList(ctx *gin.Context)
}

func NewGroupHandler(service *service.Service,
	loggers *logrus_log.Logger) *GroupHandler {
	return &GroupHandler{
		GroupEndPoint: NewGroupEndPointHandler(service, loggers),
	}
}
