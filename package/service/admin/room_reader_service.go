package admin

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type RoomReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewRoomReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *RoomReaderService {
	return &RoomReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *RoomReaderService) GetRoomList(pagination model.Pagination) (roomList []model.Room,
	err error) {
	roomList, err = s.repo.AdminRepository.GetRoomList(pagination)
	if err != nil {
		return roomList, response.ServiceError(err, codes.Internal)
	}
	return roomList, nil
}
func (s *RoomReaderService) GetRoomByID(id string) (room model.Room, err error) {
	room, err = s.repo.AdminRepository.GetRoomByID(id)
	if err != nil {
		return room, response.ServiceError(err, codes.Internal)
	}
	return room, nil
}

func (s *RoomReaderService) GetRoomGroupById(id string, pagination model.Pagination) (
	roomGroup []model.Group, err error) {
	err = s.repo.AdminRepository.CheckRoomByID(id)
	if err != nil {
		return roomGroup, response.ServiceError(err, codes.Internal)
	}
	roomGroup, err = s.repo.AdminRepository.GroupRoomByID(id, pagination)
	if err != nil {
		return roomGroup, response.ServiceError(err, codes.Internal)
	}
	return roomGroup, nil
}
