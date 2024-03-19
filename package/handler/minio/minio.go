package minio

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
)

const (
	contentType = "Content-Type"
)

type MinIOEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewMinIOEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *MinIOEndPointHandler {
	return &MinIOEndPointHandler{service: service, loggers: loggers}
}

var (
	fileError = errors.New("Error when close file multipart:")
)

// UploadImage
// @Description Upload Image
// @Tags General
// @Accept       json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/minio/upload-image [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) UploadImage(ctx *gin.Context) {
	// ctx.Request.ParseMultipartForm(1 << 25)
	logger := h.loggers
	file, err := ctx.FormFile("file")
	imageContentType := file.Header[contentType][0]
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var fileIO io.Reader
	fileMultipart, err := file.Open()
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	fileIO = fileMultipart
	imageFileName, err := h.service.MinioService.UploadImage(fileIO,
		file.Size,
		imageContentType)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	defer func(fileMultipart multipart.File) {
		err := fileMultipart.Close()
		if err != nil {
			logger.Error(fileError, err)
		}
	}(fileMultipart)
	response.HandleResponse(ctx, response.OK, nil, imageFileName)
}

// UploadImages
// @Description Upload Images
// @Tags General
// @Accept       json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param files formData file true "files"
// @Accept multipart/form-data
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/minio/upload-images [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) UploadImages(ctx *gin.Context) {
	// ctx.Request.ParseMultipartForm(1 << 25)
	var uploadedFiles []model.Files
	form, err := ctx.MultipartForm()
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	files := form.File["files"]
	for _, file := range files {
		imageContentType := file.Header[contentType][0]
		var fileIO io.Reader
		fileMultipart, err := file.Open()
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil)
			return
		}
		fileIO = fileMultipart
		imageFileName, err := h.service.MinioService.UploadImage(fileIO,
			file.Size,
			imageContentType)
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil)
			return
		}
		imageLink, err := h.service.MinioService.GetImageLink(
			imageFileName)
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil)
			return
		}
		uploadedFiles = append(uploadedFiles, model.Files{Link: imageLink, Name: imageFileName})
		defer func(fileMultipart multipart.File) {
			err := fileMultipart.Close()
			if err != nil {
				h.loggers.Error(fileError,
					err)
			}
		}(fileMultipart)
	}
	response.HandleResponse(ctx, response.OK, nil, uploadedFiles)
}

// UploadDoc
// @Description Upload doc
// @Tags General
// @Accept       json
// @Produce application/octet-stream
// @Produce application/msword
// @Produce application/vnd.openxmlformats-officedocument.wordprocessingml.document
// @Produce image/jpg
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/minio/upload-doc [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) UploadDoc(ctx *gin.Context) {
	// ctx.Request.ParseMultipartForm(1 << 25)
	logger := h.loggers
	file, err := ctx.FormFile("file")
	docContentType := file.Header[contentType][0]

	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var fileIO io.Reader
	fileMultipart, err := file.Open()
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	fileIO = fileMultipart
	imageFileName, err := h.service.MinioService.UploadDoc(fileIO, file.Size, docContentType)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	defer func(fileMultipart multipart.File) {
		err := fileMultipart.Close()
		if err != nil {
			logger.Error(fileError, err)
		}
	}(fileMultipart)
	response.HandleResponse(ctx, response.OK, nil, imageFileName)
}
