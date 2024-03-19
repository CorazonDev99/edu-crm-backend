package admin

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func RoomRouter(api *gin.Engine, handler *AdminHandler) {
	room := api.Group("/api/v1/room", middleware.AuthRequestHandler)
	{
		room.POST("/create", handler.RoomEndPoint.CreateRoom)
		room.PUT("/update/:id", handler.RoomEndPoint.UpdateRoom)
		room.DELETE("/delete/:id", handler.RoomEndPoint.DeleteRoom)
		room.GET("/list", handler.RoomEndPoint.GetRoomList)
		room.GET("/:id", handler.RoomEndPoint.GetRoomByID)
		room.GET("/group/:id", handler.RoomEndPoint.GetRoomGroupByID)
	}
}

func SettingsRouter(api *gin.Engine, handler *AdminHandler) {
	settingsAuth := api.Group("/api/v1/settings", middleware.AuthRequestHandler)
	{
		settingsAuth.POST("/upsert", handler.SettingsEndPoint.UpsertSettings)
	}
	settings := api.Group("/api/v1/settings")
	{
		settings.GET("/get", handler.SettingsEndPoint.GetSettings)
	}
}
