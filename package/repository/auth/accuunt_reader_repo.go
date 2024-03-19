package auth

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AccountReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewAccountReaderDB NewAuthAccountReaderDB returns a new instance of
// AuthAccountReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthAccountReaderDB` struct.
func NewAccountReaderDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *AccountReaderDB {
	return &AccountReaderDB{db: db, loggers: loggers}
}

func (repo *AccountReaderDB) GetAuthAccountList(pagination model.Pagination) (
	AuthAccount []model.AuthAccount, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&AuthAccount, GetAllAuthAccountQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return AuthAccount, err
		}
		loggers.Error(err)
		return AuthAccount, err
	}
	return AuthAccount, err
}
func (repo *AccountReaderDB) GetAuthAccountByID(id string) (authAccount model.
	AuthAccount, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&authAccount, GetAuthAccountByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return authAccount, err
		}
		loggers.Error(err)
		return authAccount, err
	}
	return authAccount, err
}
