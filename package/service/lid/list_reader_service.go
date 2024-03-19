package lid

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"google.golang.org/grpc/codes"
)

type ListReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewListReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *ListReaderService {
	return &ListReaderService{repo: repo, minio: minio, loggers: loggers}
}

func (s *ListReaderService) GetListList(pagination model.Pagination) (listList []model.List, err error) {
	listList, err = s.repo.LidRepository.GetListList(
		pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	for i := range listList {
		listList[i].Lid, err = s.repo.LidRepository.GetLidByListID(listList[i].ID.String())
		if err != nil {
			s.loggers.Error(err)
			continue
		}
	}
	return listList, nil
}
func (s *ListReaderService) GetListByID(id string) (board model.List,
	err error) {
	board, err = s.repo.LidRepository.GetListByID(id)
	if err != nil {
		return board, response.ServiceError(err, codes.Internal)
	}
	return board, nil
}
