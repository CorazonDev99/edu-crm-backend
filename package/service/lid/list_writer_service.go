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

type ListWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewListWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *ListWriterService {
	return &ListWriterService{repo: repo, minio: minio, loggers: loggers}
}

func (s *ListWriterService) CreateList(board model.CreateList) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, board)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.CheckBoardByID(board.BoardID.String())
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.CreateList(board)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *ListWriterService) UpdateList(board model.UpdateList) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, board)
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.CheckBoardByID(board.BoardID.String())
	if err != nil {
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.LidRepository.UpdateList(board)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *ListWriterService) DeleteList(id string) (err error) {
	err = s.repo.LidRepository.DeleteList(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
func (s *ListWriterService) ListMove(id, from, to string) (err error) {
	err = s.repo.LidRepository.ListMove(id, from, to)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
