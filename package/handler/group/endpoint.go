package group

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func GroupRouter(api *gin.Engine, handler *GroupHandler) {
	group := api.Group("/api/v1/group", middleware.AuthRequestHandler)
	{
		group.POST("/create", handler.GroupEndPoint.CreateGroup)
		group.PUT("/update/:id", handler.GroupEndPoint.UpdateGroup)
		group.DELETE("/delete/:id", handler.GroupEndPoint.DeleteGroup)
		group.GET("/list", handler.GroupEndPoint.GetGroupList)
		group.GET("/:id", handler.GroupEndPoint.GetGroupByID)
		group.DELETE("/delete-learner", handler.GroupEndPoint.DeleteUserFromGroup)
		group.GET("/student/list/:id", handler.GroupEndPoint.GetGroupStudentList)
	}
}
