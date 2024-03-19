package lid

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type LidHandler struct {
	ListEndPoint
	BoardEndPoint
	LidEndPoint
}

type ListEndPoint interface {
	CreateList(ctx *gin.Context)
	UpdateList(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
	GetListList(ctx *gin.Context)
	GetListByID(ctx *gin.Context)
	MoveList(ctx *gin.Context)
}

type BoardEndPoint interface {
	CreateBoard(ctx *gin.Context)
	UpdateBoard(ctx *gin.Context)
	DeleteBoard(ctx *gin.Context)
	GetBoardList(ctx *gin.Context)
	GetBoardByID(ctx *gin.Context)
}

type LidEndPoint interface {
	CreateLid(ctx *gin.Context)
	UpdateLid(ctx *gin.Context)
	DeleteLid(ctx *gin.Context)
	GetLidList(ctx *gin.Context)
	GetLidByID(ctx *gin.Context)
	MoveLid(ctx *gin.Context)
	ReplaceLid(ctx *gin.Context)
}

func NewLidHandler(service *service.Service, loggers *logrus_log.Logger) *LidHandler {
	return &LidHandler{
		ListEndPoint:  NewListHandler(service, loggers),
		BoardEndPoint: NewBoardHandler(service, loggers),
		LidEndPoint:   NewLidHandlerMethod(service, loggers),
	}
}
