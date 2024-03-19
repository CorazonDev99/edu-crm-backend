package lid

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type LidWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewLidWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *LidWriterService {
	return &LidWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *LidWriterService) CreateLid(board model.CreateLid) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, board)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.CheckListByID(board.ListID.String())
	if err != nil {
		return response.ServiceError(err, codes.NotFound)
	}
	err = s.repo.LidRepository.CreateLid(board)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *LidWriterService) UpdateLid(board model.UpdateLid) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, board)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.CheckListByID(board.ListID.String())
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.UpdateLid(board)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *LidWriterService) DeleteLid(id string) (err error) {
	err = s.repo.LidRepository.DeleteLid(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}

func (s *LidWriterService) LidMove(id, from, to string) (err error) {
	err = s.repo.LidRepository.LidMove(id, from, to)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}

func (s *LidWriterService) LidReplace(id, from, to string) (err error) {
	err = s.repo.LidRepository.LidReplace(id, from, to)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
