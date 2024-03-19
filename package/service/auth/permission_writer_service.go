package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type PermissionWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewPermissionWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *PermissionWriterService {
	return &PermissionWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *PermissionWriterService) CreatePermission(role model.
	CreatePermission) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, role)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	_, err = s.repo.AuthRepository.CreatePermission(role)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *PermissionWriterService) UpdatePermission(role model.
	UpdatePermission) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, role)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.AuthRepository.UpdatePermission(role)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *PermissionWriterService) DeletePermission(id string) (err error) {
	err = s.repo.AuthRepository.DeletePermission(id)
	if err != nil {
		s.loggers.Error(err)
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
