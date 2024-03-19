package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type PermissionReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewPermissionReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *PermissionReaderService {
	return &PermissionReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *PermissionReaderService) GetPermissionList(pagination model.
	Pagination) (
	roleList []model.Permission,
	err error) {
	roleList, err = s.repo.AuthRepository.GetPermissionList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return roleList, nil
}
func (s *PermissionReaderService) GetPermissionByID(id string) (role model.
	Permission,
	err error) {
	role, err = s.repo.AuthRepository.GetPermissionByID(id)
	if err != nil {
		return role, response.ServiceError(err, codes.Internal)
	}
	return role, nil
}
