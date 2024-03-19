package middleware

import (
	"EduCRM/tools/jwt"
	"EduCRM/util/response"
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	userRoleCtx         = "userRole"
)

func AuthRequestHandler(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		HandleResponse(ctx, response.Unauthorized,
			errors.New("unauthorized. Header empty"), nil)
		return
	}
	user, err := jwt.ExtractTokenMetadata(ctx)
	if err != nil {
		HandleResponse(ctx, response.Unauthorized,
			errors.New("unauthorized. "+err.Error()), nil)
		return
	}
	if user == nil {
		// Abort the request with the appropriate error code
		HandleResponse(ctx, response.Unauthorized,
			errors.New("unauthorized. User not found"), nil)
		return
	}
	ctx.Set(userIdCtx, user.UserID)
	ctx.Set(userRoleCtx, user.UserRole)
	// Continue down the chain to handler etc
	ctx.Next()
}

func AuthRefreshTokenRequestHandler(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		HandleResponse(ctx, response.Unauthorized,
			errors.New("unauthorized. Header empty"), nil)
		return
	}
	user, err := jwt.ExtractRefreshTokenMetadata(ctx)
	if err != nil {
		HandleResponse(ctx, response.Unauthorized,
			errors.New("unauthorized. "+err.Error()), nil)
		return
	}
	if user == nil {
		// Abort the request with the appropriate error code
		HandleResponse(ctx, response.Unauthorized,
			errors.New("unauthorized. User not found"), nil)
		return
	}
	ctx.Set(userIdCtx, user.UserID)
	//Continue down the chain to handler etc
	ctx.Next()
}
func HandleResponse(ctx *gin.Context, status response.Status,
	err error, data interface{}) {
	ctx.AbortWithStatusJSON(status.Code, response.ResponseModel{
		Status:       status.Status,
		Description:  status.Description,
		SnapData:     data,
		ErrorMessage: err.Error(),
	})
}
