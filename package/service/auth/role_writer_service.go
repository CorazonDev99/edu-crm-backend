package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type RoleWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewRoleWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *RoleWriterService {
	return &RoleWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *RoleWriterService) CreateRole(role model.CreateRole) (id uuid.UUID, err error) {
	err = validation.ValidationStructTag(s.loggers, role)
	if err != nil {
		return id, response.ServiceError(err, codes.InvalidArgument)
	}
	id, err = s.repo.AuthRepository.CreateRole(role)
	if err != nil {
		return id, response.ServiceError(err, codes.Internal)
	}
	return id, nil
}
func (s *RoleWriterService) UpdateRole(role model.UpdateRole) (err error) {
	err = validation.ValidationStructTag(s.loggers, role)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.AuthRepository.UpdateRole(role)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *RoleWriterService) DeleteRole(id string) (err error) {
	err = s.repo.AuthRepository.DeleteRole(id)
	if err != nil {
		s.loggers.Error(err)
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
