package auth

import (
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type ReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewAuthReaderDB returns a new instance of AuthReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthReaderDB` struct.
func NewAuthReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *ReaderDB {
	return &ReaderDB{db: db, loggers: loggers}
}
