package store

import (
	"context"
	"errors"

	"EduCRM/config"
	"EduCRM/util/logrus_log"

	"github.com/minio/minio-go/v7"
)

type ObjectMinio struct {
	minio  *minio.Client
	config *config.Configuration
	logrus *logrus_log.Logger
}

func NewObjectMinio(minio *minio.Client, config *config.Configuration, logrus *logrus_log.Logger) *ObjectMinio {
	return &ObjectMinio{minio: minio, config: config, logrus: logrus}
}

func (om *ObjectMinio) ObjectExists(imageName string) error {
	loggers := om.logrus
	client := om.minio
	configuration := om.config
	_, err := client.StatObject(context.Background(), configuration.MinioBucketName, imageName, minio.GetObjectOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		if errResponse.Code == "AccessDenied" {
			loggers.Error(errResponse)
			return errors.New("access Denied")
		}
		if errResponse.Code == "NoSuchBucket" {
			loggers.Error(errResponse)
			return errors.New("no Exist Bucket Object")
		}
		if errResponse.Code == "InvalidBucketName" {
			loggers.Error(errResponse)
			return errors.New("invalid Bucket Name")
		}
		if errResponse.Code == "NoSuchKey" {
			loggers.Error(errResponse)
			return errors.New("no Exist Image Object")
		}
		return errors.New("unknown Error")
	}
	return nil
}
