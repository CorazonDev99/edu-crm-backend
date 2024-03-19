package group

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"errors"

	"github.com/jmoiron/sqlx"
)

type GroupWriterDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

// NewGroupWriterDB returns a new instance of GroupWriterDB.
// It takes a `*sqlx.DB` and a `*loggers_log.Logger` as parameters.
// It returns a pointer to a `GroupWriterDB` struct.
func NewGroupWriterDB(db *sqlx.DB, loggers *logrus_log.Logger) *GroupWriterDB {
	return &GroupWriterDB{db: db, loggers: loggers}
}
func (repo *GroupWriterDB) CreateGroup(group model.CreateGroup) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(CreateGroupQuery, group.Title, group.CourseID,
		group.TeacherID, group.EduDays, group.RoomID.String(), group.Price,
		group.LessonStartTime, group.Status, group.StartDate, group.EndDate, group.Comment)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *GroupWriterDB) UpdateGroup(group model.UpdateGroup) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpdateGroupQuery, group.Title, group.CourseID,
		group.TeacherID, group.EduDays, group.RoomID.String(), group.Price,
		group.LessonStartTime, group.Status, group.StartDate, group.EndDate, group.Comment, group.ID)
	if err != nil {
		loggers.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		loggers.Error(err)
		return err
	}
	if rowAffected == 0 {
		loggers.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
func (repo *GroupWriterDB) DeleteGroup(id string) (err error) {
	logrus := repo.loggers
	db := repo.db
	row, err := db.Exec(DeleteGroupQuery, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		logrus.Error(ErrorNoRowsAffected)
		return ErrorNoRowsAffected
	}
	return nil
}
