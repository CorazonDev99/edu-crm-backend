package lid

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	Board
	List
	Lid
}
type Board struct {
	BoardReader
	BoardWriter
}
type BoardReader interface {
	GetBoardList(board model.Pagination) (boardList []model.Board, err error)
	GetBoardByID(id string) (board model.Board, err error)
	CheckBoardByID(id string) (err error)
}
type BoardWriter interface {
	CreateBoard(board model.CreateBoard) (err error)
	UpdateBoard(board model.UpdateBoard) (err error)
	DeleteBoard(id string) (err error)
}

type List struct {
	ListReader
	ListWriter
}
type ListReader interface {
	GetListList(list model.Pagination) (listList []model.List, err error)
	GetListByID(id string) (list model.List, err error)
	CheckListByID(id string) (err error)
	GetListByBoardID(id string) (listList []model.List, err error)
}
type ListWriter interface {
	CreateList(list model.CreateList) (err error)
	UpdateList(list model.UpdateList) (err error)
	DeleteList(id string) (err error)
	ListMove(id, from, to string) (err error)
}
type Lid struct {
	LidReader
	LidWriter
}
type LidReader interface {
	GetLidList(lid model.Pagination) (lidList []model.Lid, err error)
	GetLidByID(id string) (lid model.Lid, err error)
	CheckLidByID(id string) (err error)
	GetLidByListID(id string) (lidList []model.Lid, err error)
}
type LidWriter interface {
	CreateLid(lid model.CreateLid) (err error)
	UpdateLid(lid model.UpdateLid) (err error)
	DeleteLid(id string) (err error)
	LidMove(id, from, to string) (err error)
	LidReplace(id, from, to string) (err error)
}

func NewLidRepo(db *sqlx.DB, loggers *logrus_log.Logger) *Repo {
	return &Repo{
		Board: Board{
			BoardReader: NewBoardReaderDB(db, loggers),
			BoardWriter: NewBoardWriterDB(db, loggers),
		},
		List: List{
			ListReader: NewListReaderDB(db, loggers),
			ListWriter: NewListWriterDB(db, loggers),
		},
		Lid: Lid{
			LidReader: NewLidReaderDB(db, loggers),
			LidWriter: NewLidWriterDB(db, loggers),
		},
	}
}
