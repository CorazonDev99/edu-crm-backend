package lid

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/gin-gonic/gin"
	"math"
)

type LidHandlerMethod struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewLidHandlerMethod(service *service.Service,
	loggers *logrus_log.Logger) *LidHandlerMethod {
	return &LidHandlerMethod{service: service, loggers: loggers}
}

var lidID = "id"

// CreateLid
// @Description Create Lid
// @Summary Create Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param create body model.CreateLid true "Create Lid"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/create [post]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) CreateLid(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateLid
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.CreateLid(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateLid
// @Description Update Lid
// @Summary Update Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateLid true "Update Lid"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/update/{id} [put]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) UpdateLid(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, lidID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var (
		body model.UpdateLid
	)
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.ID = id
	err = h.service.LidService.UpdateLid(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteLid
// @Description Delete Lid
// @Summary Delete Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) DeleteLid(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, lidID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.DeleteLid(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetLidList
// @Description Get Lid List
// @Summary Get Lid List
// @Tags Lid
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/list [get]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) GetLidList(ctx *gin.Context) {
	loggers := h.loggers

	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	lidList, err := h.service.LidService.GetLidList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(lidList) == 0 {
		totalData = 0
	} else {
		totalData = lidList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: lidList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetLidByID
// @Description Get Lid
// @Summary Get Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/{id} [get]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) GetLidByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, lidID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	lid, err := h.service.LidService.GetLidByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, lid)
}

// MoveLid
// @Description Move Lid
// @Summary Move Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param lid body model.MoveLid true "lid"
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/move/{id} [put]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) MoveLid(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, lidID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var lid model.MoveLid
	err = ctx.ShouldBindJSON(&lid)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.LidMove(id.String(), lid.From.String(), lid.To.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "ok")
}

// ReplaceLid
// @Description Replace Lid
// @Summary Replace Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param lid body model.MoveLid true "lid"
// @Param lid body model.MoveLid true "Create Lid"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/lid/replace [get]
// @Security ApiKeyAuth
func (h *LidHandlerMethod) ReplaceLid(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, lidID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var list model.MoveLid
	err = ctx.ShouldBindJSON(&list)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.LidReplace(id.String(), list.From.String(), list.To.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "ok")
}
