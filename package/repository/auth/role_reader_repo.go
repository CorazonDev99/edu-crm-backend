package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoleReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewRoleReaderDB NewAuthRoleReaderDB returns a new instance of AuthRoleReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthRoleReaderDB` struct.
func NewRoleReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *RoleReaderDB {
	return &RoleReaderDB{db: db, loggers: loggers}
}

func (repo *RoleReaderDB) GetRoleList(pagination model.Pagination) (
	roleList []model.
		Role, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&roleList, GetRoleListByRoleQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return roleList, err
		}
		loggers.Error(err)
		return roleList, err
	}
	return roleList, err
}
func (repo *RoleReaderDB) GetRoleByID(id string) (role model.Role, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&role, GetRoleByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return role, err
		}
		loggers.Error(err)
		return role, err
	}
	return role, err
}
func (repo *RoleReaderDB) GetRoleTitleByID(id uuid.UUID) (role string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&role, GetRoleTitleByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return role, err
		}
		loggers.Error(err)
		return role, err
	}
	return role, err
}
func (repo *RoleReaderDB) GetRoleIDByTitle(title string) (id uuid.UUID, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&id, GetRoleIDbyTitleQuery, title)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, err
	}
	return id, err
}
func (repo *RoleReaderDB) CheckRoleByID(id uuid.UUID) (err error) {
	loggers := repo.loggers
	db := repo.db
	var roleID string
	err = db.Get(&roleID, CheckRoleByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}
