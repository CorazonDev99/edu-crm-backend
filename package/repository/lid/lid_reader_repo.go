package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type LidReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

// NewLidReaderDB  returns a new instance of LidLidReaderDB.
//
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `LidLidReaderDB` struct.
func NewLidReaderDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *LidReaderDB {
	return &LidReaderDB{db: db, loggers: loggers}
}
func (repo *LidReaderDB) GetLidList(pagination model.Pagination) (lidList []model.Lid, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&lidList, GetLidListQuery, pagination.Limit,
		pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return lidList, err
		}
		loggers.Error(err)
		return lidList, err
	}
	return lidList, err
}
func (repo *LidReaderDB) GetLidByID(id string) (lid model.Lid,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&lid, GetLidByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return lid, err
		}
		loggers.Error(err)
		return lid, err
	}
	return lid, err
}
func (repo *LidReaderDB) CheckLidByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var lidID string
	err = db.Get(&lidID, CheckLidByID, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}

func (repo *LidReaderDB) GetLidByListID(id string) (lidList []model.Lid,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&lidList, GetLidBYListIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return lidList, err
		}
		loggers.Error(err)
		return lidList, err
	}
	return lidList, err
}
