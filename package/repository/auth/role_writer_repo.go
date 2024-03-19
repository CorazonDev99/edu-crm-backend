package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoleWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewRoleWriterDB NewAuthRoleWriterDB returns a new instance of
// AuthRoleWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthRoleWriterDB` struct.
func NewRoleWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *RoleWriterDB {
	return &RoleWriterDB{db: db, loggers: loggers}
}

func (repo *RoleWriterDB) CreateRole(role model.CreateRole) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateRoleQuery, role.Title, role.Description, role.Document)
	if err != nil {
		loggers.Error(err)
		return id, err
	}
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			loggers.Error(err)
			return id, err
		}
	}
	return id, nil
}
func (repo *RoleWriterDB) UpdateRole(role model.UpdateRole) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateRoleQuery, role.Title, role.Description, role.Document,
		role.ID)
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
func (repo *RoleWriterDB) DeleteRole(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteRoleQuery, id)
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
