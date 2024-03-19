package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type LidWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewLidWriterDB returns a new instance of LidWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `LidWriterDB` struct.
func NewLidWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *LidWriterDB {
	return &LidWriterDB{db: db, loggers: loggers}
}
func (repo *LidWriterDB) CreateLid(board model.CreateLid) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(CreateLidQuery, board.ListID, board.FullName,
		board.PhoneNumber, board.Location, board.Comment)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *LidWriterDB) UpdateLid(board model.UpdateLid) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateLidQuery, board.ListID, board.FullName,
		board.PhoneNumber, board.Location, board.Comment, board.ID)
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
func (repo *LidWriterDB) DeleteLid(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteLidQuery, id)
	if err != nil {
		loggers.Error(err)
		return nil
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return nil
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}

func (repo *LidWriterDB) LidMove(id, from, to string) (err error) {
	loggers := repo.loggers
	loggers.Error(id, from, to)
	db := repo.db
	row, err := db.Exec(LidMoveQuery, id, from, to)
	if err != nil {
		loggers.Error(err)
		return nil
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return nil
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}

func (repo *LidWriterDB) LidReplace(id, from, to string) (err error) {
	loggers := repo.loggers
	loggers.Error(id, from, to)
	db := repo.db
	row, err := db.Exec(LidReplaceQuery, id, from, to)
	if err != nil {
		loggers.Error(err)
		return nil
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return nil
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
