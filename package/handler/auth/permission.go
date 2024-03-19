package auth

import (
	"EduCRM/model"
	"EduCRM/util/response"
	"math"

	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

const (
	permissionID = "id"
)

type PermissionHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewPermissionHandler(service *service.Service,
	loggers *logrus_log.Logger) *PermissionHandler {
	return &PermissionHandler{service: service, loggers: loggers}
}

// CreatePermission
// @Description Create Permission
// @Summary Create Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param create body model.CreatePermission true "Create Permission"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/permission/create [post]
// @Security ApiKeyAuth
func (h *PermissionHandler) CreatePermission(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreatePermission
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AuthService.CreatePermission(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdatePermission
// @Description Update Permission
// @Summary Update Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdatePermission true "Update Permission"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/permission/update/{id} [put]
// @Security ApiKeyAuth
func (h *PermissionHandler) UpdatePermission(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, permissionID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var (
		body model.UpdatePermission
	)
	err = ctx.ShouldBindJSON(&body)
	body.ID = id
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AuthService.UpdatePermission(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeletePermission
// @Description Delete Permission
// @Summary Delete Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/permission/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *PermissionHandler) DeletePermission(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, permissionID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AuthService.DeletePermission(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetPermissionList
// @Description Get Permission List
// @Summary Get Permission List
// @Tags Permission
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/permission/list [get]
// @Security ApiKeyAuth
func (h *PermissionHandler) GetPermissionList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	permissionList, err := h.service.AuthService.GetPermissionList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(permissionList) == 0 {
		totalData = 0
	} else {
		totalData = permissionList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: permissionList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetPermissionByID
// @Description Get Permission
// @Summary Get Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/permission/{id} [get]
// @Security ApiKeyAuth
func (h *PermissionHandler) GetPermissionByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, permissionID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	permissionList, err := h.service.AuthService.GetPermissionByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, permissionList)
}
