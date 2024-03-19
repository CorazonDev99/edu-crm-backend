package admin

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/gin-gonic/gin"
	"math"
)

type RoomEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewRoomEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *RoomEndPointHandler {
	return &RoomEndPointHandler{service: service, loggers: loggers}
}

var roomID = "id"

// CreateRoom
// @Description Create Room
// @Summary Create Room
// @Tags Room
// @Accept json
// @Produce json
// @Param create body model.CreateRoom true "Create Room"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/room/create [post]
// @Security ApiKeyAuth
func (h *RoomEndPointHandler) CreateRoom(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateRoom
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AdminService.CreateRoom(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateRoom
// @Description Update Room
// @Summary Update Room
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateRoom true "Update Room"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/room/update/{id} [put]
// @Security ApiKeyAuth
func (h *RoomEndPointHandler) UpdateRoom(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateRoom
	)
	id, err := handler_func.GetUUIDParam(ctx, roomID)
	if err != nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.ID = id
	err = h.service.AdminService.UpdateRoom(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteRoom
// @Description Delete Room
// @Summary Delete Room
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/room/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *RoomEndPointHandler) DeleteRoom(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, roomID)
	if err != nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AdminService.DeleteRoom(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetRoomList
// @Description Get Room List
// @Summary Get Room List
// @Tags Room
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/room/list [get]
// @Security ApiKeyAuth
func (h *RoomEndPointHandler) GetRoomList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	roomList, err := h.service.AdminService.GetRoomList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(roomList) == 0 {
		totalData = 0
	} else {
		totalData = roomList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: roomList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetRoomByID
// @Description Get Room
// @Summary Get Room
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/room/{id} [get]
// @Security ApiKeyAuth
func (h *RoomEndPointHandler) GetRoomByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, roomID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	room, err := h.service.AdminService.GetRoomByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, room)
}

// GetRoomGroupByID
// @Description Get Room group
// @Summary Get Room group
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/room/group/{id} [get]
// @Security ApiKeyAuth
func (h *RoomEndPointHandler) GetRoomGroupByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, roomID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
	}
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	groupList, err := h.service.AdminService.GetRoomGroupById(id.String(), pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(groupList) == 0 {
		totalData = 0
	} else {
		totalData = groupList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: groupList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}
