package admin

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"EduCRM/util/validation"
	"google.golang.org/grpc/codes"
)

type SettingsService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logrus_log.Logger
}

func NewSettingsService(repo *repository.Repository, minio *store.Store,
	loggers *logrus_log.Logger) *SettingsService {
	return &SettingsService{repo: repo, minio: minio, loggers: loggers}
}

func (s *SettingsService) UpsertSettings(settings model.CreateSettings) (
	err error) {
	err = validation.ValidationStructTag(s.loggers, settings)
	if err != nil {
		loggers := s.loggers
		loggers.Error(err)
		return response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.AdminRepository.UpsertSettings(settings)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}

func (s *SettingsService) GetSettings() (settings model.Settings, err error) {
	settings, err = s.repo.AdminRepository.GetSettings()
	if err != nil {
		return settings, response.ServiceError(err, codes.Internal)
	}
	if len(settings.CompanyLogo) != 0 {
		err = s.minio.ObjectStore.ObjectExists(settings.CompanyLogo)
		if err != nil {
			s.loggers.Error(err)
			settings.CompanyLogoLink = ""
			return settings, nil
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(settings.CompanyLogo)
		if err != nil {
			s.loggers.Error(err)
			settings.CompanyLogoLink = ""
			return settings, nil
		}
		settings.CompanyLogoLink = image
	}
	if len(settings.SiteEnterLogo) != 0 {
		err = s.minio.ObjectStore.ObjectExists(settings.SiteEnterLogo)
		if err != nil {
			s.loggers.Error(err)
			settings.SiteEnterLogoLink = ""
			return settings, nil
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(settings.SiteEnterLogo)
		if err != nil {
			s.loggers.Error(err)
			settings.SiteEnterLogoLink = ""
			return settings, nil
		}
		settings.SiteEnterLogoLink = image
	}
	if len(settings.InstructionFile) != 0 {
		err = s.minio.ObjectStore.ObjectExists(settings.InstructionFile)
		if err != nil {
			s.loggers.Error(err)
			settings.InstructionFileLink = ""
			return settings, nil
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(settings.InstructionFile)
		if err != nil {
			s.loggers.Error(err)
			settings.InstructionFileLink = ""
			return settings, nil
		}
		settings.InstructionFileLink = image
	}
	return settings, nil
}
