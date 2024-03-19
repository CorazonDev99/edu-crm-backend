package admin

import (
	"EduCRM/model"
	"EduCRM/util/response"

	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type SettingsEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewSettingsEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *SettingsEndPointHandler {
	return &SettingsEndPointHandler{service: service, loggers: loggers}
}

// UpsertSettings
// @Description Create or Update Settings
// @Summary Create or Update Settings
// @Tags Settings
// @Accept json
// @Produce json
// @Param create body model.CreateSettings true "Create,Update,Delete Settings"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/settings/upsert [post]
// @Security ApiKeyAuth
func (h *SettingsEndPointHandler) UpsertSettings(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateSettings
	)

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AdminService.UpsertSettings(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// GetSettings
// @Description Get Settings
// @Summary Get Settings
// @Tags Settings
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/settings/get [get]
func (h *SettingsEndPointHandler) GetSettings(ctx *gin.Context) {
	loggers := h.loggers
	settings, err := h.service.AdminService.GetSettings()
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, settings)
}
