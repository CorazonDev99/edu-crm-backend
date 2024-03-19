package store

import (
	"EduCRM/config"
	"EduCRM/util/logrus_log"
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"strings"
)

type UploadMinio struct {
	minio   *minio.Client
	config  *config.Configuration
	loggers *logrus_log.Logger
}

var (
	docContentType  = "msword"
	docxContentType = "vnd.openxmlformats-officedocument.wordprocessingml.document"
)

func NewUploadMinio(minio *minio.Client, config *config.Configuration, loggers *logrus_log.Logger) *UploadMinio {
	return &UploadMinio{minio: minio, config: config, loggers: loggers}
}
func (um *UploadMinio) UploadImage(file io.Reader, imageSize int64, contextType string) (string, error) {
	loggers := um.loggers
	fileName := uuid.New()
	fileExtension := strings.Split(contextType, "/")[1]
	imageName := fileName.String() + "." + fileExtension
	_, err := um.minio.PutObject(context.Background(), um.config.MinioBucketName, imageName, file, imageSize, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		loggers.Info("Internal Server Error: ", err.Error())
		return "", err
	}
	return imageName, nil
}
func (um *UploadMinio) UploadDoc(file io.Reader, docSize int64,
	contentType string) (string, error) {
	loggers := um.loggers
	fileName := uuid.New()
	fileExtension := "docx"
	loggers.Error(contentType)
	if strings.Contains(contentType, docContentType) {
		fileExtension = "doc"
	}
	if strings.Contains(contentType, docxContentType) {
		fileExtension = "docx"
	}
	docFileName := fileName.String() + "." + fileExtension
	_, err := um.minio.PutObject(context.Background(),
		um.config.MinioBucketName, docFileName, file, docSize,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		loggers.Info("Internal Server Error: ", err.Error())
		return "", err
	}
	return docFileName, nil
}
