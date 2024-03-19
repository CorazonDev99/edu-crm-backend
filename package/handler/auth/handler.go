package auth

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	RoleEndPoint
	PermissionEndPoint
	TokenEndpoint
}

type RoleEndPoint interface {
	CreateRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	GetRoleList(ctx *gin.Context)
	GetRoleByID(ctx *gin.Context)
	GetRoleTitleByID(ctx *gin.Context)
}
type PermissionEndPoint interface {
	CreatePermission(ctx *gin.Context)
	UpdatePermission(ctx *gin.Context)
	DeletePermission(ctx *gin.Context)
	GetPermissionList(ctx *gin.Context)
	GetPermissionByID(ctx *gin.Context)
}
type TokenEndpoint interface {
	RefreshToken(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func NewAuthHandler(service *service.Service,
	loggers *logrus_log.Logger) *AuthHandler {
	return &AuthHandler{
		RoleEndPoint:       NewRoleHandler(service, loggers),
		PermissionEndPoint: NewPermissionHandler(service, loggers),
		TokenEndpoint:      NewTokenHandler(service, loggers),
	}
}
