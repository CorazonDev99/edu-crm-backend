package auth

import (
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/tools/jwt"
	"EduCRM/util/handler_func"
	"EduCRM/util/logrus_log"
	"EduCRM/util/response"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	service *service.Service
	loggers *logrus_log.Logger
}

func NewTokenHandler(service *service.Service,
	loggers *logrus_log.Logger) *TokenHandler {
	return &TokenHandler{service: service, loggers: loggers}
}

// RefreshToken
// @Description RefreshToken Auth Refresh Token
// @Summary RefreshToken
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/refresh-token [get]
// @Security ApiKeyAuth
func (h *TokenHandler) RefreshToken(ctx *gin.Context) {
	loggers := h.loggers
	userID, err := handler_func.GetUserId(ctx)
	if err != nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.Internal, err, nil)
		return
	}
	account, roleTitle, err := h.service.AuthService.GetAuthAccountByID(userID)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
	}
	tokens, err := jwt.GenerateNewTokens(account.AccountID.String(), account.RoleID.String(), roleTitle)
	if err != nil {
		response.HandleResponse(ctx, response.InternalServerError,
			err, nil)
		return
	}
	err = h.service.AuthService.UpdateAuthAccount(model.UpdateAuthAccount{
		AccountID:    account.AccountID,
		RoleID:       account.RoleID,
		RefreshToken: account.RefreshToken,
		AccessToken:  tokens.Access,
	})
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, tokens.Access)
}

// Logout
// @Description Log Out Auth Refresh Token
// @Summary Log Out
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/auth/logout [get]
// @Security ApiKeyAuth
func (h *TokenHandler) Logout(ctx *gin.Context) {
	loggers := h.loggers
	userID, err := handler_func.GetUserId(ctx)
	if err != nil {
		loggers.Error(err)
		response.HandleResponse(ctx, response.Internal, err, nil)
		return
	}
	account, _, err := h.service.AuthService.GetAuthAccountByID(userID)
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
	}
	err = h.service.AuthService.UpdateAuthAccount(model.UpdateAuthAccount{
		AccountID:    account.AccountID,
		RoleID:       account.RoleID,
		RefreshToken: "",
		AccessToken:  "",
	})
	if err != nil {
		loggers.Error(err)
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "logout")
}
