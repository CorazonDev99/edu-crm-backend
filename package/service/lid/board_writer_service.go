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

type BoardWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewBoardWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *BoardWriterService {
	return &BoardWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *BoardWriterService) CreateBoard(board model.CreateBoard) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, board)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.CreateBoard(board)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *BoardWriterService) UpdateBoard(board model.UpdateBoard) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, board)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.UpdateBoard(board)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *BoardWriterService) DeleteBoard(id string) (err error) {
	err = s.repo.LidRepository.DeleteBoard(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
