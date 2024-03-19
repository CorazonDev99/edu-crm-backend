package store

import (
	"io"

	"EduCRM/config"
	"EduCRM/util/logrus_log"

	"github.com/minio/minio-go/v7"
)

type Store struct {
	UploadStore
	ObjectStore
	FileLinkStore
}

type UploadStore interface {
	UploadImage(imageFile io.Reader, imageSize int64, contextType string) (string, error)
	UploadDoc(docFile io.Reader, docSize int64, contextType string) (string, error)
}

type ObjectStore interface {
	ObjectExists(name string) error
}

type FileLinkStore interface {
	GetImageUrl(imageName string) (string, error)
}

func NewStore(minio *minio.Client, config *config.Configuration, logrus *logrus_log.Logger) *Store {
	return &Store{UploadStore: NewUploadMinio(minio, config, logrus), ObjectStore: NewObjectMinio(minio, config, logrus), FileLinkStore: NewFileLinkMinio(minio, config, logrus)}
}
