package lid

import (
	"EduCRM/model"
	"EduCRM/util/response"
	"math"

	"EduCRM/package/service"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

var boardID = "id"

type BoardHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewBoardHandler(service *service.Service,
	loggers *logrus_log.Logger) *BoardHandler {
	return &BoardHandler{service: service, loggers: loggers}
}

// CreateBoard
// @Description Create Board
// @Summary Create Board
// @Tags Board
// @Accept json
// @Produce json
// @Param create body model.CreateBoard true "Create Board"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/board/create [post]
// @Security ApiKeyAuth
func (h *BoardHandler) CreateBoard(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateBoard
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.LidService.CreateBoard(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created")
}

// UpdateBoard
// @Description Update Board
// @Summary Update Board
// @Tags Board
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param update body model.UpdateBoard true "Update Board"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/board/update/{id} [put]
// @Security ApiKeyAuth
func (h *BoardHandler) UpdateBoard(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateBoard
	)
	id, err := handler_func.GetUUIDParam(ctx, boardID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil)
		return
	}
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.ID = id
	err = h.service.LidService.UpdateBoard(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteBoard
// @Description Delete Board
// @Summary Delete Board
// @Tags Board
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/board/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *BoardHandler) DeleteBoard(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, boardID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil)
		return
	}
	err = h.service.LidService.DeleteBoard(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "deleted")
}

// GetBoardList
// @Description Get Board List
// @Summary Get Board List
// @Tags Board
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/board/list [get]
// @Security ApiKeyAuth
func (h *BoardHandler) GetBoardList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	boardList, err := h.service.LidService.GetBoardList(pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(boardList) == 0 {
		totalData = 0
	} else {
		totalData = boardList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: boardList,
		Pagination: model.Pagination{
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
			Page:      pagination.Page,
			PageSize:  pagination.PageSize,
		},
	})
}

// GetBoardByID
// @Description Get Board
// @Summary Get Board
// @Tags Board
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/board/{id} [get]
// @Security ApiKeyAuth
func (h *BoardHandler) GetBoardByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, boardID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil)
		return
	}
	board, err := h.service.LidService.GetBoardByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, board)
}
