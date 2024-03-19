package admin

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
)

type AdminService struct {
	RoomReader
	RoomWriter
	Settings
}
type RoomReader interface {
	GetRoomList(pagination model.Pagination) (Room []model.Room, err error)
	GetRoomByID(id string) (room model.Room, err error)
	GetRoomGroupById(id string, pagination model.Pagination) (roomGroup []model.Group, err error)
}
type RoomWriter interface {
	CreateRoom(room model.CreateRoom) (err error)
	UpdateRoom(room model.UpdateRoom) (err error)
	DeleteRoom(id string) (err error)
}
type Settings interface {
	UpsertSettings(settings model.CreateSettings) (err error)
	GetSettings() (settings model.Settings, err error)
}

func NewAdminService(repos *repository.Repository, store *store.Store,
	loggers *logrus_log.Logger) *AdminService {
	return &AdminService{
		Settings:   NewSettingsService(repos, store, loggers),
		RoomReader: NewRoomReaderService(repos, store, loggers),
		RoomWriter: NewRoomWriterService(repos, store, loggers),
	}
}
