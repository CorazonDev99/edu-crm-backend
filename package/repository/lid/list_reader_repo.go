package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ListReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewListReaderDB  returns a new instance of ListListReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `ListListReaderDB` struct.
func NewListReaderDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *ListReaderDB {
	return &ListReaderDB{db: db, loggers: loggers}
}
func (repo *ListReaderDB) GetListList(pagination model.Pagination) (
	boardList []model.List, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&boardList, GetListListQuery, pagination.Limit,
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
func (repo *ListReaderDB) GetListByID(id string) (board model.List,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&board, GetListByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return board, err
		}
		loggers.Error(err)
		return board, err
	}
	return board, err
}
func (repo *ListReaderDB) CheckListByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var boardID string
	err = db.Get(&boardID, CheckListByID, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}

func (repo *ListReaderDB) GetListByBoardID(id string) (boardListList []model.
	List,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&boardListList, GetBoardListListQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return boardListList, err
		}
		loggers.Error(err)
		return boardListList, err
	}
	return boardListList, nil
}
