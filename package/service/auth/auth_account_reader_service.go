package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type AuthAccountReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewAuthAccountReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *AuthAccountReaderService {
	return &AuthAccountReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *AuthAccountReaderService) GetAuthAccountList(pagination model.Pagination) (roleList []model.AuthAccount,
	err error) {
	roleList, err = s.repo.AuthRepository.GetAuthAccountList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return roleList, nil
}
func (s *AuthAccountReaderService) GetAuthAccountByID(id string) (role model.AuthAccount, roleTitle string, err error) {
	role, err = s.repo.AuthRepository.GetAuthAccountByID(id)
	if err != nil {
		return role, roleTitle, response.ServiceError(err, codes.Internal)
	}
	roleTitle, err = s.repo.AuthRepository.GetRoleTitleByID(role.RoleID)
	if err != nil {
		return model.AuthAccount{}, "", err
	}
	return role, roleTitle, nil
}
