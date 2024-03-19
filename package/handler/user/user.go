package user

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/tools/jwt"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math"
	"strings"
)

var (
	userID     = "id"
	superAdmin = "super$007Admin"
)

type UserEndPointHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewUserEndPointHandler(service *service.Service,
	loggers *logrus_log.Logger) *UserEndPointHandler {
	return &UserEndPointHandler{service: service, loggers: loggers}
}

// CreateUser
// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body model.CreateUser true "Create User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/create [post]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) CreateUser(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.CreateUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	accountID, err := h.service.UserService.CreateUser(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	err = h.service.AuthService.CreateAuthAccount(model.CreateAuthAccount{
		RoleID:    body.RoleID,
		AccountID: accountID,
	})
	if err != nil {
		loggers.Error(err)
		//response.ServiceErrorConvert(ctx, err)
		//return
	}
	response.HandleResponse(ctx, response.Created, nil,
		"created")
}

// UpdateUser
// @Description Update User
// @Summary Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param update body model.UpdateUser true "Update User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/update/{id} [put]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) UpdateUser(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.UpdateUser
	)
	userID, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil)
		return
	}
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.ID = userID
	err = h.service.UserService.UpdateUser(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	err = h.service.AuthService.UpdateAuthAccount(model.UpdateAuthAccount{
		RoleID:    body.RoleID,
		AccountID: body.ID,
	})
	if err != nil {
		loggers.Error(err)
		//response.ServiceErrorConvert(ctx, err)
		//return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// DeleteUser
// @Description Delete User
// @Summary Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) DeleteUser(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.UserService.DeleteUser(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, errors.New("deleted"))
}

// GetUserList
// @Description Get User List
// @Summary Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Param role  query string true "Role : all.  get all user"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/list [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserList(ctx *gin.Context) {
	loggers := h.loggers
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	role, err := handler_func.GetStringQuery(ctx, "role")
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	userList, err := h.service.UserService.GetUserList(role, pagination)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	var totalData int64
	if len(userList) == 0 {
		totalData = 0
	} else {
		totalData = userList[0].Total
	}
	response.HandleResponse(ctx, response.OK, nil, model.DataList{
		List: userList,
		Pagination: model.Pagination{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
			PageTotal: int64(math.Ceil(float64(totalData) / float64(pagination.
				PageSize))),
			ItemTotal: totalData,
		},
	})
}

// GetUserByID
// @Description Get User
// @Summary Get User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/{id} [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserByID(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil)
		return
	}
	userList, err := h.service.UserService.GetUserByID(id.String())
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, userList)
}

// SignInUser
// @Description Admin Sign In  User.
// @Summary Admin Sign In User
// @Tags User
// @Accept json
// @Produce json
// @Param signup body model.SignInUser true "Sign In"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/sign-in [post]
func (h *UserEndPointHandler) SignInUser(ctx *gin.Context) {
	loggers := h.loggers
	var (
		body model.SignInUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
	}
	body.PhoneNumber = strings.TrimSpace(body.PhoneNumber)
	body.Password = strings.TrimSpace(body.Password)
	id, role, roleTitle, err := h.service.UserService.SignInUser(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	if id == uuid.Nil || role == uuid.Nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.BadRequest,
			errors.New("username or password is incorrect"), nil)
		return
	}
	tokens, err := jwt.GenerateNewTokens(id.String(), role.String(), roleTitle)
	if err != nil {
		response.HandleResponse(ctx, response.InternalServerError,
			err, nil)
		return
	}
	err = h.service.AuthService.UpdateAuthAccount(model.UpdateAuthAccount{
		AccountID:    id,
		RoleID:       role,
		RefreshToken: tokens.Refresh,
		AccessToken:  tokens.Access,
	})
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, tokens)
}

// UpdateUserPassword
// @Description Update User Password
// @Summary Update User Password
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param password body model.UserPassword true "password"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/update-password/{id} [put]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) UpdateUserPassword(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	var user model.UserPassword
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	err = h.service.UserService.UpdateUserPassword(id.String(), user.Password)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated")
}

// GetTeacherGroupList
// // @Description Get Teacher Course
// // @Summary Get Teacher Course
// // @Tags User
// // @Accept json
// // @Produce json
// // @Param teacherId  query string true "User ID"
// // @Param  courseId query string true "course ID"
// // @Success 200 {object} response.ResponseModel
// // @Failure 400 {object} response.ResponseModel
// // @Failure 404 {object} response.ResponseModel
// // @Failure 500 {object} response.ResponseModel
// // @Router /api/v1/teacher/course/list [GET]
// // @Security ApiKeyAuth
//
//	response_func (h *UserEndPointHandler) GetTeacherCourseList(ctx *gin.
//	Context) {
//		loggers := response.loggers
//		teacherID, err := response.GetUUIDQuery(ctx, userID)
//		if err != nil {
//			response.handleResponse(ctx, response.BadRequest, err.Error())
//			return
//		}
//		data, err := response.service.UserService.GetTeacherCourseList(
//		teacherID)
//		if err != nil {
//			loggers.Error(err)
//			response.ServiceErrorConvert(ctx, err)
//			return
//		}
//		response.handleResponse(ctx, response.OK, data)
//	}

// GetUserMe
// @Description User Me
// @Summary User Me
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/me [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserMe(ctx *gin.Context) {
	loggers := h.loggers
	id, err := handler_func.GetUserId(ctx)
	if err != nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.NotFound, err, nil)
		return
	}
	user, err := h.service.UserService.GetUserByID(id)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, user)
}

// CreateSuperAdmin
// @Description Create Super Admin
// @Summary Create Super Admin
// @Tags SuperAdmin
// @Accept json
// @Produce json
// @Param token query string true "Create Super Admin"
// @Param create body model.CreateUser true "Create Super Admin"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/super-admin/create [post]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) CreateSuperAdmin(ctx *gin.Context) {
	loggers := h.loggers
	token, err := handler_func.GetStringQuery(ctx, "token")
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	if token != superAdmin {
		response.HandleResponse(ctx, response.PermissionDenied, err, nil)
		return
	}
	roleID, err := h.service.AuthService.CreateRole(model.CreateRole{Title: "superAdmin", Description: "", Document: ""})
	var (
		body model.CreateUser
	)
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil)
		return
	}
	body.RoleID = roleID
	accountID, err := h.service.UserService.CreateUser(body)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	err = h.service.AuthService.CreateAuthAccount(model.CreateAuthAccount{
		RoleID:    body.RoleID,
		AccountID: accountID,
	})
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil,
		"created")
}
