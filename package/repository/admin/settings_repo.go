package admin

import (
	"EduCRM/model"
	"EduCRM/util/logrus_log"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type SettingsDB struct {
	db      *sqlx.DB
	loggers *logrus_log.Logger
}

func NewSettingsDB(db *sqlx.DB,
	loggers *logrus_log.Logger) *SettingsDB {
	return &SettingsDB{db: db, loggers: loggers}
}
func (repo *SettingsDB) GetSettings() (settings model.Settings, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&settings, GetSettingsQuery)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return settings, err
		}
		loggers.Error(err)
		return settings, err
	}
	return settings, err
}

func (repo *SettingsDB) UpsertSettings(settings model.CreateSettings) (err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Exec(UpsertSettingsQuery, settings.CompanyTitle,
		settings.CompanyLogo, settings.SiteEnterLogo, settings.OpenDate,
		settings.CompanyPhone, settings.SiteColor, settings.InstructionFile)
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
