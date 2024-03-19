package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type AccountWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewAccountWriterDB NewAuthAccountWriterDB returns a new instance of
// AuthAccountWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthAccountWriterDB` struct.
func NewAccountWriterDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *AccountWriterDB {
	return &AccountWriterDB{db: db, loggers: loggers}
}

func (repo *AccountWriterDB) CreateAuthAccount(authAccount model.
	CreateAuthAccount) (id string, err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(CreateAuthAccountQuery, authAccount.AccountID,
		authAccount.RoleID)
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
func (repo *AccountWriterDB) UpdateAuthAccount(authAccount model.
	UpdateAuthAccount) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateAuthAccountQuery, authAccount.RefreshToken, authAccount.ID, authAccount.RoleID, authAccount.AccountID)
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
func (repo *AccountWriterDB) DeleteAuthAccount(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteAuthAccountQuery, id)
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
