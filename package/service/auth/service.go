package auth

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"
)

type AuthService struct {
	Role
	Permission
	AuthAccount
}

type Role struct {
	RoleReader
	RoleWriter
}

type RoleReader interface {
	GetRoleList(pagination model.Pagination) (Role []model.Role, err error)
	GetRoleByID(id string) (role model.Role, err error)
	GetRoleTitleByID(title string) (roleID uuid.UUID, err error)
}
type RoleWriter interface {
	CreateRole(role model.CreateRole) (id uuid.UUID, err error)
	UpdateRole(role model.UpdateRole) (err error)
	DeleteRole(id string) (err error)
}

type Permission struct {
	PermissionReader
	PermissionWriter
}
type PermissionReader interface {
	GetPermissionList(pagination model.Pagination) (Permission []model.Permission, err error)
	GetPermissionByID(id string) (role model.Permission, err error)
}
type PermissionWriter interface {
	CreatePermission(role model.CreatePermission) (err error)
	UpdatePermission(role model.UpdatePermission) (err error)
	DeletePermission(id string) (err error)
}
type AuthAccount struct {
	AuthAccountReader
	AuthAccountWriter
}
type AuthAccountReader interface {
	GetAuthAccountList(pagination model.Pagination) (AuthAccount []model.AuthAccount, err error)
	GetAuthAccountByID(id string) (role model.AuthAccount, roleTitle string, err error)
}
type AuthAccountWriter interface {
	CreateAuthAccount(auth model.CreateAuthAccount) (err error)
	UpdateAuthAccount(auth model.UpdateAuthAccount) (err error)
}

func NewAuthService(repo *repository.Repository, minio *store.Store, loggers *logrus_log.Logger) *AuthService {
	return &AuthService{
		Role: Role{
			RoleReader: NewRoleReaderService(repo, minio, loggers),
			RoleWriter: NewRoleWriterService(repo, minio, loggers),
		},

		Permission: Permission{
			PermissionReader: NewPermissionReaderService(repo, minio, loggers),
			PermissionWriter: NewPermissionWriterService(repo, minio, loggers),
		},
		AuthAccount: AuthAccount{
			AuthAccountReader: NewAuthAccountReaderService(repo, minio, loggers),
			AuthAccountWriter: NewAuthAccountWriterService(repo, minio, loggers),
		},
	}
}
