package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type RoleReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewRoleReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *RoleReaderService {
	return &RoleReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *RoleReaderService) GetRoleList(pagination model.Pagination) (roleList []model.Role,
	err error) {
	roleList, err = s.repo.AuthRepository.GetRoleList(pagination)
	for i, role := range roleList {
		if len(role.Document) != 0 {
			err := s.minio.ObjectStore.ObjectExists(role.Document)
			if err != nil {
				s.loggers.Error(err)
				roleList[i].DocumentLink = ""
			} else {
				image, err := s.minio.FileLinkStore.GetImageUrl(role.Document)
				if err != nil {
					s.loggers.Error(err)
					roleList[i].DocumentLink = ""
				}
				roleList[i].DocumentLink = image
			}
		}
	}
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return roleList, nil
}
func (s *RoleReaderService) GetRoleByID(id string) (role model.Role, err error) {
	role, err = s.repo.AuthRepository.GetRoleByID(id)
	if len(role.Document) != 0 {
		err = s.minio.ObjectStore.ObjectExists(role.Document)
		if err != nil {
			s.loggers.Error(err)
			role.DocumentLink = ""
			return role, nil
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(role.Document)
		if err != nil {
			s.loggers.Error(err)
			role.DocumentLink = ""
			return role, nil
		}
		role.DocumentLink = image
	}
	if err != nil {
		return role, response.ServiceError(err, codes.Internal)
	}
	return role, nil
}

func (s *RoleReaderService) GetRoleTitleByID(roleTitle string) (id uuid.UUID, err error) {
	id, err = s.repo.AuthRepository.GetRoleIDByTitle(roleTitle)
	if err != nil {
		s.loggers.Error(err)
		return id, response.ServiceError(err, codes.InvalidArgument)
	}
	return id, nil
}
