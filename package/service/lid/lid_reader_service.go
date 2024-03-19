package lid

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type LidReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewLidReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *LidReaderService {
	return &LidReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *LidReaderService) GetLidList(pagination model.Pagination) (
	boardLid []model.Lid, err error) {
	boardLid, err = s.repo.LidRepository.GetLidList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	return boardLid, nil
}
func (s *LidReaderService) GetLidByID(id string) (board model.Lid,
	err error) {
	board, err = s.repo.LidRepository.GetLidByID(id)
	if err != nil {
		return board, response.ServiceError(err, codes.Internal)
	}
	return board, nil
}
