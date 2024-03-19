package admin

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"errors"

	"github.com/jmoiron/sqlx"
)

type RoomWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

func NewRoomWriterDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *RoomWriterDB {
	return &RoomWriterDB{db: db, loggers: loggers}
}

func (repo *RoomWriterDB) CreateRoom(room model.CreateRoom) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(CreateRoomQuery, room.Title, room.Description,
		room.RoomNumber, room.CloseTime, room.OpenTime)
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
func (repo *RoomWriterDB) UpdateRoom(room model.UpdateRoom) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateRoomQuery, room.Title, room.Description,
		room.RoomNumber, room.CloseTime, room.OpenTime, room.ID)
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
func (repo *RoomWriterDB) DeleteRoom(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteRoomQuery, id)
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
