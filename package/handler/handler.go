package handler

import (
	"EduCRM/config"
	"EduCRM/docs"
	"EduCRM/package/handler/admin"
	"EduCRM/package/handler/auth"
	"EduCRM/package/handler/course"
	"EduCRM/package/handler/group"
	"EduCRM/package/handler/lid"
	"EduCRM/package/handler/minio"
	"EduCRM/package/handler/user"
	"EduCRM/package/service"
	"EduCRM/tools/middleware"
	"EduCRM/util/logrus_log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

//
//type Handler struct {
//	service *service.Service
//	loggers *loggers_log.Logger
//	config  *config.Configuration
//}
//
//handler_func NewHandler(service *service.Service, loggers *loggers_log.Logger, config *config.Configuration) *Handler {
//	return &Handler{service: service, loggers: loggers, config: config}
//}

type Handler struct {
	Lid    *lid.LidHandler
	Auth   *auth.AuthHandler
	Course *course.CourseHandler
	Group  *group.GroupHandler
	Admin  *admin.AdminHandler
	User   *user.UserHandler
	MinIO  *minio.MinIOHandler
}

func NewHandler(service *service.Service,
	loggers *logrus_log.Logger) *Handler {
	return &Handler{
		Lid:    lid.NewLidHandler(service, loggers),
		Auth:   auth.NewAuthHandler(service, loggers),
		Course: course.NewCourseHandler(service, loggers),
		Group:  group.NewGroupHandler(service, loggers),
		Admin:  admin.NewAdminHandler(service, loggers),
		User:   user.NewUserHandler(service, loggers),
		MinIO:  minio.NewMinIOHandler(service, loggers),
	}
}

func (handler *Handler) InitRoutes() (route *gin.Engine) {
	cfg := config.Config()
	route = gin.New()
	//gin.SetMode(gin.ReleaseMode)
	route.HandleMethodNotAllowed = true
	middleware.GinMiddleware(route)
	//swagger settings
	docs.SwaggerInfo.Title = cfg.AppName
	docs.SwaggerInfo.Version = cfg.AppVersion
	//docs.SwaggerInfo.Host = cfg.AppURL
	//docs.SwaggerInfo.BasePath = cfg.AppURL + "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	serverHost := ""
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler),
		func(ctx *gin.Context) {
			serverHost = ctx.Request.Host
		})
	docs.SwaggerInfo.Host = serverHost

	//routers
	minio.MinIORouter(route, handler.MinIO)
	user.UserRouter(route, handler.User)
	course.CourseRouter(route, handler.Course)
	group.GroupRouter(route, handler.Group)
	auth.AuthRoleRouter(route, handler.Auth)
	auth.AuthPermissionRouter(route, handler.Auth)
	admin.SettingsRouter(route, handler.Admin)
	admin.RoomRouter(route, handler.Admin)
	lid.ListRouter(route, handler.Lid)
	lid.BoardRouter(route, handler.Lid)
	lid.LidRouter(route, handler.Lid)

	return
}
