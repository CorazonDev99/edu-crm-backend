package admin

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type RoomWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewRoomWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *RoomWriterService {
	return &RoomWriterService{repo: repo, minio: minio, loggers: loggers}
}
func (s RoomWriterService) CreateRoom(room model.CreateRoom) (err error) {
	err = validation.ValidationStructTag(s.loggers, room)
	if err != nil {
		loggers := s.loggers
		loggers.Error(err)
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.AdminRepository.CreateRoom(room)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s RoomWriterService) UpdateRoom(room model.UpdateRoom) (err error) {
	err = validation.ValidationStructTag(s.loggers, room)
	if err != nil {
		loggers := s.loggers
		loggers.Error(err)
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.AdminRepository.UpdateRoom(room)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s RoomWriterService) DeleteRoom(id string) (err error) {
	err = s.repo.AdminRepository.DeleteRoom(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
