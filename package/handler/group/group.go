package group

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/gin-gonic/gin"
	"math"
)

var (
	groupID   = "id"
	learnerID = "learnerId"
	userID    = "id"
)

type GroupEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewGroupEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *GroupEndPointHandler {
	return &GroupEndPointHandler{service: service, loggers: loggers}
}

// CreateGroup
// @Description Create Group
// @Summary Create Group
// @Tags Group
// @Accept json
// @Produce json
// @Param create body model.CreateGroup true "Create Group"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/create [post]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) CreateGroup(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateGroup
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.GroupService.CreateGroup(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateGroup
// @Description Update Group
// @Summary Update Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateGroup true "Update Group"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/update/{id} [put]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) UpdateGroup(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateGroup
	)
	id, err := handler_func.GetUUIDParam(ctx, groupID)
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
	err = h.service.GroupService.UpdateGroup(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteGroup
// @Description Delete Group
// @Summary Delete Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) DeleteGroup(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, groupID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.GroupService.DeleteGroup(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetGroupList
// @Description Get Group List
// @Summary Get Group List
// @Tags Group
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/list [get]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) GetGroupList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	groupList, err := h.service.GroupService.GetGroupList(pagination)
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

// GetGroupByID
// @Description Get Group
// @Summary Get Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/{id} [get]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) GetGroupByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, groupID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	group, err := h.service.GroupService.GetGroupByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, group)
}

// GetGroupStudentList
// @Description Get Group
// @Summary Get Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/student/list/{id} [get]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) GetGroupStudentList(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, groupID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	group, err := h.service.GroupService.GetGroupStudentList(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, group)
}

// DeleteUserFromGroup
// @Description Delete User From Group
// @Summary Delete User From Group
// @Tags User
// @Accept json
// @Produce json
// @Param learnerId  query string true "Learner ID"
// @Param groupId  query string true "Group ID"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/group/delete-learner [DELETE]
// @Security ApiKeyAuth
func (h *GroupEndPointHandler) DeleteUserFromGroup(ctx *gin.Context) {
	loggers := h.loggers
	learnerID, err := handler_func.GetUUIDQuery(ctx, learnerID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	groupID, err := handler_func.GetUUIDQuery(ctx, groupID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.GroupService.DeleteUserFromGroup(groupID, learnerID)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "ok")
}
