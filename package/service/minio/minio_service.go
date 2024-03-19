package minio

import (
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"io"
)

type MinioService struct {
	store   *store.Store
	loggers *logrus_log.Logger
}

func NewMinioService(store *store.Store,
	loggers *logrus_log.Logger) *MinioService {
	return &MinioService{store: store, loggers: loggers}
}

func (m *MinioService) UploadImage(imageFile io.Reader, imageSize int64,
	contextType string) (string, error) {
	imageName, err := m.store.UploadStore.UploadImage(imageFile, imageSize, contextType)
	if err != nil {
		return "", err
	}
	return imageName, nil
}
func (m *MinioService) UploadDoc(docFile io.Reader, docSize int64,
	contextType string) (string, error) {
	imageName, err := m.store.UploadStore.UploadDoc(docFile, docSize, contextType)
	if err != nil {
		return "", err
	}
	return imageName, nil
}

func (m *MinioService) GetImageLink(imageName string) (string, error) {
	imageLink, err := m.store.FileLinkStore.GetImageUrl(imageName)
	if err != nil {
		return "", err
	}
	return imageLink, nil
}
