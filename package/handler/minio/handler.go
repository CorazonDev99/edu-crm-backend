package minio

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type MinIOHandler struct {
	MinIOEndPoint
}

type MinIOEndPoint interface {
	UploadImages(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
	UploadDoc(ctx *gin.Context)
}

func NewMinIOHandler(service *service.Service,
	loggers *logrus_log.Logger) *MinIOHandler {
	return &MinIOHandler{
		MinIOEndPoint: NewMinIOEndPointHandler(service, loggers),
	}
}
