package admin

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type RoomReaderDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

func NewRoomReaderDB(db *sqlx.DB, loggers *logrus_log.Logger) *RoomReaderDB {
	return &RoomReaderDB{db: db, loggers: loggers}
}
func (repo *RoomReaderDB) GetSettings() (settings model.Settings, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&settings, GetSettingsQuery)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return settings, err
		}
		loggers.Error(err)
		return settings, err
	}
	return settings, err
}

func (repo *RoomReaderDB) GetRoomList(pagination model.Pagination) (
	roomList []model.Room,
	err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&roomList, GetRoomListQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return roomList, err
		}
		loggers.Error(err)
		return roomList, err
	}
	return roomList, err
}
func (repo *RoomReaderDB) GetRoomByID(id string) (room model.Room, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&room, GetRoomByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return room, err
		}
		loggers.Error(err)
		return room, err
	}
	return room, err
}

func (repo *RoomReaderDB) CheckRoomByID(id string) (err error) {
	loggers := repo.loggers
	var roomID string
	db := repo.db
	err = db.Get(&roomID, CheckRoomByIDQuery, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return err
}

func (repo *RoomReaderDB) GroupRoomByID(id string, pagination model.Pagination) (groupList []model.Group, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&groupList, GetGroupListByRoomIDQuery, id, pagination.Limit, pagination.Offset)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return groupList, err
		}
		loggers.Error(err)
		return groupList, err
	}
	return groupList, nil
}
