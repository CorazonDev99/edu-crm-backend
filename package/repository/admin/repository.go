package admin

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	Settings
	Room
}
type Settings struct {
	SettingsRepository
}
type SettingsRepository interface {
	GetSettings() (settings model.Settings, err error)
	UpsertSettings(settings model.CreateSettings) (err error)
}
type Room struct {
	RoomReader
	RoomWriter
}
type RoomReader interface {
	GetRoomList(pagination model.Pagination) (Room []model.Room, err error)
	GetRoomByID(id string) (room model.Room, err error)
	CheckRoomByID(id string) (err error)
	GroupRoomByID(id string, pagination model.Pagination) (roomGroup []model.Group, err error)
}
type RoomWriter interface {
	CreateRoom(room model.CreateRoom) (err error)
	UpdateRoom(room model.UpdateRoom) (err error)
	DeleteRoom(id string) (err error)
}

func NewAdminRepo(db *sqlx.DB, loggers *logrus_log.Logger) *Repo {
	return &Repo{
		Settings: Settings{
			SettingsRepository: NewSettingsDB(db, loggers),
		},
		Room: Room{
			RoomReader: NewRoomReaderDB(db, loggers),
			RoomWriter: NewRoomWriterDB(db, loggers),
		},
	}
}
