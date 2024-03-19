package minio

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func MinIORouter(api *gin.Engine, handler *MinIOHandler) {
	minio := api.Group("/api/v1/minio", middleware.AuthRequestHandler)
	{
		minio.POST("/upload-images", handler.MinIOEndPoint.UploadImages)
		minio.POST("/upload-image", handler.MinIOEndPoint.UploadImage)
		minio.POST("/upload-doc", handler.MinIOEndPoint.UploadDoc)
	}
}
