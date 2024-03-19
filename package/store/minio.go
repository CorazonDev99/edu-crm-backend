package store

import (
	"context"

	"EduCRM/config"
	"EduCRM/util/logrus_log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection(config *config.Configuration, logrus *logrus_log.Logger) (*minio.Client, error) {

	ctx := context.Background()

	// Initialize minio client object.
	minioClient, errInit := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKeyID, config.MinioSecretKey, ""),
		Secure: config.MinioUseSSL,
	})
	if errInit != nil {
		logrus.Fatalln(errInit)
	}

	err := minioClient.MakeBucket(ctx, config.MinioBucketName, minio.MakeBucketOptions{Region: config.MinioLocation})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, config.MinioBucketName)
		if errBucketExists == nil && exists {
			logrus.Infof("We already own %s\n", config.MinioBucketName)
		} else {
			logrus.Fatal(err)
		}
	} else {
		logrus.Infof("Successfully created %s\n", config.MinioBucketName)
	}
	return minioClient, errInit
}
