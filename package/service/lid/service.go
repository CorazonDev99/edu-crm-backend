package lid

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
)

type LidService struct {
	BoardMethod
	ListMethod
	LidMethod
}

type BoardMethod struct {
	BoardReaderMethod
	BoardWriterMethod
}

type BoardReaderMethod interface {
	GetBoardList(group model.Pagination) (groupList []model.Board, err error)
	GetBoardByID(id string) (group model.Board, err error)
}
type BoardWriterMethod interface {
	CreateBoard(group model.CreateBoard) (err error)
	UpdateBoard(group model.UpdateBoard) (err error)
	DeleteBoard(id string) (err error)
}
type ListMethod struct {
	ListReaderMethod
	ListWriterMethod
}

type ListReaderMethod interface {
	GetListList(group model.Pagination) (groupList []model.List, err error)
	GetListByID(id string) (group model.List, err error)
}
type ListWriterMethod interface {
	CreateList(group model.CreateList) (err error)
	UpdateList(group model.UpdateList) (err error)
	DeleteList(id string) (err error)
	ListMove(id, from, to string) (err error)
}

type LidMethod struct {
	LidReader
	LidWriter
}
type LidReader interface {
	GetLidList(group model.Pagination) (groupList []model.Lid, err error)
	GetLidByID(id string) (group model.Lid, err error)
}
type LidWriter interface {
	CreateLid(group model.CreateLid) (err error)
	UpdateLid(group model.UpdateLid) (err error)
	DeleteLid(id string) (err error)
	LidMove(id, from, to string) (err error)
	LidReplace(id, from, to string) (err error)
}

func NewLidService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *LidService {
	return &LidService{
		BoardMethod{
			BoardReaderMethod: NewBoardReaderService(repo, minio, loggers),
			BoardWriterMethod: NewBoardWriterService(repo, minio, loggers),
		},
		ListMethod{
			ListReaderMethod: NewListReaderService(repo, minio, loggers),
			ListWriterMethod: NewListWriterService(repo, minio, loggers),
		},
		LidMethod{
			LidReader: NewLidReaderService(repo, minio, loggers),
			LidWriter: NewLidWriterService(repo, minio, loggers),
		},
	}
}
