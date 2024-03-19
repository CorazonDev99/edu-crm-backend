package auth

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoleRouter(api *gin.Engine, handler *AuthHandler) {
	auth := api.Group("/api/v1/auth", middleware.AuthRequestHandler)
	{
		authRole := auth.Group("/role")
		{
			authRole.POST("/create", handler.RoleEndPoint.CreateRole)
			authRole.PUT("/update/:id", handler.RoleEndPoint.UpdateRole)
			authRole.DELETE("/delete/:id", handler.RoleEndPoint.DeleteRole)
			authRole.GET("/list", handler.RoleEndPoint.GetRoleList)
			authRole.GET("/:id", handler.RoleEndPoint.GetRoleByID)
			authRole.GET("/title", handler.RoleEndPoint.GetRoleTitleByID)
		}
	}
	authToken := api.Group("/api/v1/auth", middleware.AuthRefreshTokenRequestHandler)
	{
		authToken.GET("/refresh-token", handler.RefreshToken)
		authToken.GET("/logout", handler.Logout)
	}
}

func AuthPermissionRouter(api *gin.Engine, handler *AuthHandler) {
	auth := api.Group("/api/v1/auth", middleware.AuthRequestHandler)
	{
		authPermission := auth.Group("/permission")
		{
			authPermission.POST("/create", handler.PermissionEndPoint.
				CreatePermission)
			authPermission.PUT("/update/:id", handler.PermissionEndPoint.
				UpdatePermission)
			authPermission.DELETE("/delete/:id", handler.PermissionEndPoint.
				DeletePermission)
			authPermission.GET("/list", handler.PermissionEndPoint.
				GetPermissionList)
			authPermission.GET("/:id", handler.PermissionEndPoint.
				GetPermissionByID)
		}
	}
}
