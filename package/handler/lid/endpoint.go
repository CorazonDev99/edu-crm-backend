package lid

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func BoardRouter(api *gin.Engine, handler *LidHandler) {
	board := api.Group("/api/v1/board", middleware.AuthRequestHandler)
	{
		board.POST("/create", handler.BoardEndPoint.CreateBoard)
		board.PUT("/update/:id", handler.BoardEndPoint.UpdateBoard)
		board.DELETE("/delete/:id", handler.BoardEndPoint.DeleteBoard)
		board.GET("/list", handler.BoardEndPoint.GetBoardList)
		board.GET("/:id", handler.BoardEndPoint.GetBoardByID)
		//board.GET("/dashboard", handler.BoardEndPoint.DashboardList)
	}
}
func ListRouter(api *gin.Engine, handler *LidHandler) {
	list := api.Group("/api/v1/list", middleware.AuthRequestHandler)
	{
		list.POST("/create", handler.ListEndPoint.CreateList)
		list.PUT("/update/:id", handler.ListEndPoint.UpdateList)
		list.DELETE("/delete/:id", handler.ListEndPoint.DeleteList)
		list.GET("/list", handler.ListEndPoint.GetListList)
		list.GET("/:id", handler.ListEndPoint.GetListByID)
		list.PUT("/move/:id", handler.ListEndPoint.MoveList)

	}
}

func LidRouter(api *gin.Engine, handler *LidHandler) {
	lid := api.Group("/api/v1/lid", middleware.AuthRequestHandler)
	{
		lid.POST("/create", handler.LidEndPoint.CreateLid)
		lid.PUT("/update/:id", handler.LidEndPoint.UpdateLid)
		lid.DELETE("/delete/:id", handler.LidEndPoint.DeleteLid)
		lid.GET("/list", handler.LidEndPoint.GetLidList)
		lid.GET("/:id", handler.LidEndPoint.GetLidByID)
		lid.PUT("/move/:id", handler.LidEndPoint.MoveLid)
		//lid.GET("/replace/:id", handler.LidEndPoint.ReplaceLid)
	}
}
