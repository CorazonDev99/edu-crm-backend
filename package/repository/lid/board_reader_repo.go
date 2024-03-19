package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type BoardReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewBoardReaderDB  returns a new instance of BoardBoardReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `BoardBoardReaderDB` struct.
func NewBoardReaderDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *BoardReaderDB {
	return &BoardReaderDB{db: db, loggers: loggers}
}
func (repo *BoardReaderDB) GetBoardList(pagination model.Pagination) (boardList []model.Board, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&boardList, GetBoardListQuery, pagination.Limit,
		pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return boardList, err
		}
		loggers.Error(err)
		return boardList, err
	}
	return boardList, err
}
func (repo *BoardReaderDB) GetBoardByID(id string) (board model.Board,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&board, GetBoardByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return board, err
		}
		loggers.Error(err)
		return board, err
	}
	return board, err
}
func (repo *BoardReaderDB) CheckBoardByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var boardID string
	err = db.Get(&boardID, CheckBoardByID, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}
