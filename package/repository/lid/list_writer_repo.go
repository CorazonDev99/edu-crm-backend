package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type ListWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewListWriterDB returns a new instance of ListWriterDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `ListWriterDB` struct.
func NewListWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *ListWriterDB {
	return &ListWriterDB{db: db, loggers: loggers}
}
func (repo *ListWriterDB) CreateList(board model.CreateList) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(CreateListQuery, board.BoardID, board.Title)
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
func (repo *ListWriterDB) UpdateList(board model.UpdateList) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateListQuery, board.Title, board.ID)
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
func (repo *ListWriterDB) DeleteList(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteListQuery, id)
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

func (repo *ListWriterDB) ListMove(id, from, to string) (err error) {
	loggers := repo.loggers
	loggers.Error(id, from, to)
	db := repo.db
	row, err := db.Exec(ListMoveQuery, id, from, to)
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
