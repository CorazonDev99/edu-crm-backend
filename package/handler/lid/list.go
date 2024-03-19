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

type ListHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewListHandler(service *service.Service, loggers *logrus_log.Logger) *ListHandler {
	return &ListHandler{service: service, loggers: loggers}
}

var listID = "id"

// CreateList
// @Description Create List
// @Summary Create List
// @Tags List
// @Accept json
// @Produce json
// @Param create body model.CreateList true "Create List"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/list/create [post]
// @Security ApiKeyAuth
func (h *ListHandler) CreateList(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateList
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.CreateList(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateList
// @Description Update List
// @Summary Update List
// @Tags List
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateList true "Update List"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/list/update/{id} [put]
// @Security ApiKeyAuth
func (h *ListHandler) UpdateList(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateList
	)
	id, err := handler_func.GetUUIDParam(ctx, "id")
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.ID = id
	err = h.service.LidService.UpdateList(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteList
// @Description Delete List
// @Summary Delete List
// @Tags List
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/list/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *ListHandler) DeleteList(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, listID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.DeleteList(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetListList
// @Description Get List List
// @Summary Get List List
// @Tags List
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/list/list [get]
// @Security ApiKeyAuth
func (h *ListHandler) GetListList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	listList, err := h.service.LidService.GetListList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(listList) == 0 {
		totalData = 0
	} else {
		totalData = listList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: listList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetListByID
// @Description Get List
// @Summary Get List
// @Tags List
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/list/{id} [get]
// @Security ApiKeyAuth
func (h *ListHandler) GetListByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, listID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	list, err := h.service.LidService.GetListByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, list)
}

// MoveList
// @Description Move Lid
// @Summary Move Lid
// @Tags List
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param move body model.MoveList true "Move Lid"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/list/move/{id} [put]
// @Security ApiKeyAuth
func (h *ListHandler) MoveList(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, listID)
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
	err = h.service.LidService.ListMove(id.String(), lid.From.String(), lid.To.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "ok")
}
