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
	roleID    = "id"
	roleTitle = "title"
)

type RoleHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewRoleHandler(service *service.Service,
	loggers *logrus_log.Logger) *RoleHandler {
	return &RoleHandler{service: service, loggers: loggers}
}

// CreateRole
// @Description Create Role
// @Summary Create Role
// @Tags Role
// @Accept json
// @Produce json
// @Param create body model.CreateRole true "Create Role"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/role/create [post]
// @Security ApiKeyAuth
func (h *RoleHandler) CreateRole(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateRole
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	_, err = h.service.AuthService.CreateRole(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateRole
// @Description Update Role
// @Summary Update Role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateRole true "Update Role"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/role/update/{id} [put]
// @Security ApiKeyAuth
func (h *RoleHandler) UpdateRole(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateRole
	)
	id, err := handler_func.GetUUIDParam(ctx, roleID)
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.ID = id
	err = h.service.AuthService.UpdateRole(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteRole
// @Description Delete Role
// @Summary Delete Role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/role/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *RoleHandler) DeleteRole(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, roleID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.AuthService.DeleteRole(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetRoleList
// @Description Get Role List
// @Summary Get Role List
// @Tags Role
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/role/list [get]
// @Security ApiKeyAuth
func (h *RoleHandler) GetRoleList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	roleList, err := h.service.AuthService.GetRoleList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(roleList) == 0 {
		totalData = 0
	} else {
		totalData = roleList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: roleList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetRoleByID
// @Description Get Role
// @Summary Get Role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/role/{id} [get]
// @Security ApiKeyAuth
func (h *RoleHandler) GetRoleByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, roleID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	roleList, err := h.service.AuthService.GetRoleByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, roleList)
}

// GetRoleTitleByID
// @Description Get Role
// @Summary Get Role
// @Tags Role
// @Accept json
// @Produce json
// @Param title query string true "Role title"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/role/title [get]
// @Security ApiKeyAuth
func (h *RoleHandler) GetRoleTitleByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetStringQuery(ctx, roleTitle)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	roleTitle, err := h.service.AuthService.GetRoleTitleByID(id)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, roleTitle)
}
