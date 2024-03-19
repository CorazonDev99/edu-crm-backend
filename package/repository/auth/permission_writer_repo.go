package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type PermissionWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewPermissionWriterDB NewAuthPermissionWriterDB returns a new instance of
// AuthPermissionWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthPermissionWriterDB` struct.
func NewPermissionWriterDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *PermissionWriterDB {
	return &PermissionWriterDB{db: db, loggers: loggers}
}

func (repo *PermissionWriterDB) CreatePermission(permission model.
	CreatePermission) (
	id string,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreatePermissionQuery, permission.Title,
		permission.Description, permission.Tag, permission.URL, permission.Method)
	if err != nil {
		loggers.Error(err)
		return "", err
	}
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			loggers.Error(err)
			return "", err
		}
	}
	return id, nil
}
func (repo *PermissionWriterDB) UpdatePermission(permission model.
	UpdatePermission) (
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdatePermissionQuery, permission.Title,
		permission.Description, permission.Tag, permission.URL,
		permission.Method, permission.ID)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *PermissionWriterDB) DeletePermission(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeletePermissionQuery, id)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
