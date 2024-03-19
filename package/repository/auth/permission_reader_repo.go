package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type PermissionReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewPermissionReaderDB NewAuthPermissionReaderDB returns a new instance of
// AuthPermissionReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthPermissionReaderDB` struct.
func NewPermissionReaderDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *PermissionReaderDB {
	return &PermissionReaderDB{db: db, loggers: loggers}
}

func (repo *PermissionReaderDB) GetPermissionList(pagination model.
	Pagination) (Permission []model.Permission, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&Permission, GetPermissionListQuery, pagination.Limit,
		pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return Permission, err
		}
		loggers.Error(err)
		return Permission, err
	}
	return Permission, err
}
func (repo *PermissionReaderDB) GetPermissionByID(id string) (permission model.
	Permission,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&permission, GetPermissionByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return permission, err
		}
		loggers.Error(err)
		return permission, err
	}
	return permission, err
}
func (repo *PermissionReaderDB) CheckPermissionByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var permissionID string
	err = db.Get(&permissionID, GetPermissionByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}
