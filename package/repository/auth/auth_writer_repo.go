package auth

import (
	"EduCRM/util/logrus_log"
	"errors"

	"github.com/jmoiron/sqlx"
)

type WriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

// NewAuthWriterDB returns a new instance of AuthWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `AuthWriterDB` struct.
func NewAuthWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *WriterDB {
	return &WriterDB{db: db, loggers: loggers}
}
