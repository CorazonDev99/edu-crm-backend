package store

import (
	"context"
	"net/url"
	"time"

	"EduCRM/config"
	"EduCRM/util/logrus_log"

	"github.com/minio/minio-go/v7"
)

type FileLinkMinio struct {
	minio   *minio.Client
	config  *config.Configuration
	loggers *logrus_log.Logger
}

func NewFileLinkMinio(minio *minio.Client, config *config.Configuration, logrus *logrus_log.Logger) *FileLinkMinio {
	return &FileLinkMinio{minio: minio, config: config, loggers: logrus}
}

func (flm *FileLinkMinio) GetImageUrl(imageName string) (string, error) {

	expiry := time.Second * 24 * 60 * 60 * 7
	loggers := flm.loggers
	client := flm.minio
	configuration := flm.config
	presignedURL, err := client.PresignedGetObject(context.Background(), configuration.MinioBucketName, imageName, expiry, url.Values{})
	if err != nil {
		loggers.Error("Error while getting object URL: ", err.Error())
		return "", err
	}

	return presignedURL.String(), nil
}
