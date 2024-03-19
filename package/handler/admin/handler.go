package admin

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	RoomEndPoint
	SettingsEndPoint
}

type RoomEndPoint interface {
	CreateRoom(ctx *gin.Context)
	UpdateRoom(ctx *gin.Context)
	DeleteRoom(ctx *gin.Context)
	GetRoomList(ctx *gin.Context)
	GetRoomByID(ctx *gin.Context)
	GetRoomGroupByID(ctx *gin.Context)
}
type SettingsEndPoint interface {
	UpsertSettings(ctx *gin.Context)
	GetSettings(ctx *gin.Context)
}

func NewAdminHandler(service *service.Service, loggers *logrus_log.Logger) *AdminHandler {
	return &AdminHandler{
		RoomEndPoint:     NewRoomEndPointHandler(service, loggers),
		SettingsEndPoint: NewSettingsEndPointHandler(service, loggers),
	}
}
