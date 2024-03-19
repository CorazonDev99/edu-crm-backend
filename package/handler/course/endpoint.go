package course

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func CourseRouter(api *gin.Engine, handler *CourseHandler) {
	course := api.Group("/api/v1/course", middleware.AuthRequestHandler)
	{
		course.POST("/create", handler.CourseEndPoint.CreateCourse)
		course.PUT("/update/:id", handler.CourseEndPoint.UpdateCourse)
		course.DELETE("/delete/:id", handler.CourseEndPoint.DeleteCourse)
		course.GET("/list", handler.CourseEndPoint.GetCourseList)
		course.GET("/:id", handler.CourseEndPoint.GetCourseByID)
		course.GET("/group/:id", handler.CourseEndPoint.GetCourseGroupList)
	}
}
