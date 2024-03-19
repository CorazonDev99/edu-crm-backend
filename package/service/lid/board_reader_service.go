package lid

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type BoardReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewBoardReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *BoardReaderService {
	return &BoardReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *BoardReaderService) GetBoardList(pagination model.Pagination) (
	boardList []model.Board, err error) {
	boardList, err = s.repo.LidRepository.GetBoardList(
		pagination)
	if err != nil {
		s.loggers.Error(err)
		return nil, response.ServiceError(err, codes.Internal)
	}
	for i := range boardList {
		boardList[i].List, err = s.repo.LidRepository.GetListByBoardID(
			boardList[i].ID.String())
		if err != nil {
			s.loggers.Error(err)
			continue
		}
		for j := range boardList[i].List {
			boardList[i].List[j].Lid, err = s.repo.LidRepository.GetLidByListID(
				boardList[i].List[j].ID.String())
			if err != nil {
				s.loggers.Error(err)
				continue
			}
		}
	}

	return boardList, nil
}
func (s *BoardReaderService) GetBoardByID(id string) (board model.Board,
	err error) {
	board, err = s.repo.LidRepository.GetBoardByID(id)
	if err != nil {
		return board, response.ServiceError(err, codes.Internal)
	}
	return board, nil
}
