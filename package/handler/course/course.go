package course

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/gin-gonic/gin"
	"math"
)

const (
	courseId = "id"
)

type CourseEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewCourseEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *CourseEndPointHandler {
	return &CourseEndPointHandler{service: service, loggers: loggers}
}

// CreateCourse
// @Description Create Course
// @Summary Create Course
// @Tags Course
// @Accept json
// @Produce json
// @Param create body model.CreateCourse true "Create Course"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/course/create [post]
// @Security ApiKeyAuth
func (h *CourseEndPointHandler) CreateCourse(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateCourse
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.CourseService.CreateCourse(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateCourse
// @Description Update Course
// @Summary Update Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateCourse true "Update Course"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/course/update/{id} [put]
// @Security ApiKeyAuth
func (h *CourseEndPointHandler) UpdateCourse(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateCourse
	)
	id, err := handler_func.GetUUIDParam(ctx, courseId)
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
	err = h.service.CourseService.UpdateCourse(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteCourse
// @Description Delete Course
// @Summary Delete Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/course/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *CourseEndPointHandler) DeleteCourse(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, courseId)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.CourseService.DeleteCourse(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetCourseList
// @Description Get Course List
// @Summary Get Course List
// @Tags Course
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/course/list [get]
// @Security ApiKeyAuth
func (h *CourseEndPointHandler) GetCourseList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	courseList, err := h.service.CourseService.GetCourseList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(courseList) == 0 {
		totalData = 0
	} else {
		totalData = courseList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: courseList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetCourseByID
// @Description Get Course
// @Summary Get Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/course/{id} [get]
// @Security ApiKeyAuth
func (h *CourseEndPointHandler) GetCourseByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, courseId)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	course, err := h.service.CourseService.GetCourseByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, course)
}

// GetCourseGroupList
// @Description Get Course List
// @Summary Get Course List
// @Tags Course
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/course/group/{id} [get]
// @Security ApiKeyAuth
func (h *CourseEndPointHandler) GetCourseGroupList(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, courseId)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	courseList, err := h.service.CourseService.GetCourseGroupList(id.String(), pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(courseList) == 0 {
		totalData = 0
	} else {
		totalData = courseList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: courseList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}
