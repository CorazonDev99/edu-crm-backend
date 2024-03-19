package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type AuthAccountWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewAuthAccountWriterService(repo *repository.Repository,
	minio *store.Store,
	loggers *logrus_log.Logger) *AuthAccountWriterService {
	return &AuthAccountWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *AuthAccountWriterService) CreateAuthAccount(role model.CreateAuthAccount) (err error) {
	_, err = s.repo.AuthRepository.CreateAuthAccount(role)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *AuthAccountWriterService) UpdateAuthAccount(role model.UpdateAuthAccount) (err error) {
	err = s.repo.AuthRepository.UpdateAuthAccount(role)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
