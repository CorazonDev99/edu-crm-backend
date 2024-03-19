package user

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.Engine, handler *UserHandler) {
	user := api.Group("/api/v1/user")
	{
		user.POST("/sign-in", handler.UserEndPoint.SignInUser)
		user.POST("/super-admin/create", handler.UserEndPoint.CreateSuperAdmin)
	}
	userAuth := api.Group("/api/v1/user", middleware.AuthRequestHandler)
	{
		userAuth.POST("/create", handler.UserEndPoint.CreateUser)
		userAuth.PUT("/update/:id", handler.UserEndPoint.UpdateUser)
		userAuth.DELETE("/delete/:id", handler.UserEndPoint.DeleteUser)
		userAuth.GET("/list", handler.UserEndPoint.GetUserList)
		userAuth.GET("/:id", handler.UserEndPoint.GetUserByID)
		userAuth.GET("/me", handler.UserEndPoint.GetUserMe)
		userAuth.PUT("/update-password/:id", handler.UserEndPoint.UpdateUserPassword)

	}
	teacherCourseGroupAuth := api.Group("/api/v1/teacher", middleware.AuthRequestHandler)
	{
		teacherCourseAuth := teacherCourseGroupAuth.Group("/course/")
		{
			teacherCourseAuth.GET("/list/:id", handler.TeacherEndPoint.GetTeacherCourseList)

		}
		teacherGroupAuth := teacherCourseGroupAuth.Group("/group/")
		{
			teacherGroupAuth.GET("/list/:id", handler.TeacherEndPoint.GetTeacherGroupList)

		}
	}
	studentAuth := api.Group("/api/v1/student", middleware.AuthRequestHandler)
	{
		studentCourseAuth := studentAuth.Group("/course/")
		{
			studentCourseAuth.GET("/list/:id", handler.StudentEndPoint.GetStudentCourseList)

		}
		studentGroupAuth := studentAuth.Group("/group/")
		{
			studentGroupAuth.GET("/list/:id", handler.StudentEndPoint.GetStudentGroupList)
		}
	}
}
