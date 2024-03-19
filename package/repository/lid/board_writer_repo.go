package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"errors"

	"github.com/jmoiron/sqlx"
)

type BoardWriterDB struct {
	db     *sqlx.DB
	logrus *logrus_log.Logger
}

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

// NewBoardWriterDB returns a new instance of BoardWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `BoardWriterDB` struct.
func NewBoardWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *BoardWriterDB {
	return &BoardWriterDB{db: db, logrus: loggers}
}
func (repo *BoardWriterDB) CreateBoard(board model.CreateBoard) (err error) {
	loggers := repo.logrus
	db := repo.db
	row, err := db.Exec(CreateBoardQuery, board.Title)
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
func (repo *BoardWriterDB) UpdateBoard(board model.UpdateBoard) (err error) {
	loggers := repo.logrus
	db := repo.db
	row, err := db.Exec(UpdateBoardQuery, board.Title, board.ID)
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
func (repo *BoardWriterDB) DeleteBoard(id string) (err error) {
	loggers := repo.logrus
	db := repo.db
	row, err := db.Exec(DeleteBoardQuery, id)
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
