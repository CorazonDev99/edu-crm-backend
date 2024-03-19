package user

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/gin-gonic/gin"
	"math"
)

type StudentEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewStudentEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *StudentEndPointHandler {
	return &StudentEndPointHandler{service: service, loggers: loggers}
}

// GetStudentGroupList
// @Description Get Student Group
// @Summary Get Student Group
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/student/group/list/{id} [GET]
// @Security ApiKeyAuth
func (h *StudentEndPointHandler) GetStudentGroupList(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	data, err := h.service.UserService.GetStudentGroupList(id.String(), pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(data) == 0 {
		totalData = 0
	} else {
		totalData = data[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: data,
		Pagination: model.Pagination{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
		},
	})
}

// GetStudentCourseList
// @Description Get Student Group
// @Summary Get Student Group
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/student/course/list/{id} [GET]
// @Security ApiKeyAuth
func (h *StudentEndPointHandler) GetStudentCourseList(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	data, err := h.service.UserService.GetStudentCourseList(id.String(), pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(data) == 0 {
		totalData = 0
	} else {
		totalData = data[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: data,
		Pagination: model.Pagination{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
		},
	})
}
