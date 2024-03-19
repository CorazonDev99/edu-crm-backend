package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	RoleRepository
	AccountRepository
	PermissionRepository
}
type RoleRepository struct {
	RoleReader
	RoleWriter
}
type RoleReader interface {
	GetRoleList(pagination model.Pagination) (Role []model.Role, err error)
	GetRoleByID(id string) (role model.Role, err error)
	CheckRoleByID(id uuid.UUID) (err error)
	GetRoleTitleByID(id uuid.UUID) (role string, err error)
	GetRoleIDByTitle(title string) (id uuid.UUID, err error)
}
type RoleWriter interface {
	CreateRole(role model.CreateRole) (id uuid.UUID, err error)
	UpdateRole(role model.UpdateRole) (err error)
	DeleteRole(id string) (err error)
}
type AccountRepository struct {
	AccountReader
	AccountWriter
	AccountRoleEnrollmentWriter
}
type AccountReader interface {
	GetAuthAccountList(pagination model.Pagination) (AuthAccount []model.AuthAccount, err error)
	GetAuthAccountByID(id string) (authAccount model.AuthAccount, err error)
}
type AccountWriter interface {
	CreateAuthAccount(authAccount model.CreateAuthAccount) (id string, err error)
	UpdateAuthAccount(authAccount model.UpdateAuthAccount) (err error)
	DeleteAuthAccount(id string) (err error)
}
type AccountRoleEnrollmentWriter interface {
	CreateAccountRoleEnrollment(accountID, roleID string) (id string, err error)
	DeleteAccountRoleEnrollment(accountID, roleID string) (err error)
	UpdateAccountRoleEnrollment(accountID, oldRoleID, newRoleID string) (err error)
}

type PermissionRepository struct {
	PermissionReader
	PermissionWriter
}
type PermissionReader interface {
	GetPermissionList(pagination model.Pagination) (Permission []model.
		Permission,
		err error)
	GetPermissionByID(id string) (router model.Permission, err error)
	CheckPermissionByID(id string) (err error)
}
type PermissionWriter interface {
	CreatePermission(router model.CreatePermission) (id string, err error)
	UpdatePermission(router model.UpdatePermission) (err error)
	DeletePermission(id string) (err error)
}

func NewAuthRepo(db *sqlx.DB, loggers *logrus_log.Logger) *AuthRepo {
	return &AuthRepo{
		RoleRepository: RoleRepository{
			RoleReader: NewRoleReaderDB(db, loggers),
			RoleWriter: NewRoleWriterDB(db, loggers),
		},
		AccountRepository: AccountRepository{
			AccountReader: NewAccountReaderDB(db, loggers),
			AccountWriter: NewAccountWriterDB(db, loggers),
			//AccountRoleEnrollmentWriter: NewAccountRoleEnrollmentWriterDB(db, loggers),
		},

		PermissionRepository: PermissionRepository{
			PermissionReader: NewPermissionReaderDB(db, loggers),
			PermissionWriter: NewPermissionWriterDB(db, loggers),
		},
	}
}
