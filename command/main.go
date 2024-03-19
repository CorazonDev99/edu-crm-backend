package main

import (
	"EduCRM/config"
	"EduCRM/package/handler"
	"EduCRM/package/repository"

	"EduCRM/package/service"
	"EduCRM/package/store"
	"EduCRM/util/logrus_log"
	"fmt"
)

// @title EduCRM
// @version 1.0
// @description API Server for EduCRM Application
// @termsOfService gitlab.com/edu-crm
// @host gitlab.com/edu-crm
// @BasePath
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	loggers := logrus_log.GetLogger()
	cfg := config.Config()
	// Migration Up
	err := repository.MigratePsql(cfg, loggers, true)
	if err != nil {
		loggers.Fatal("error while migrate up", err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Username: cfg.PostgresUser,
		DBName:   cfg.PostgresDatabase,
		SSLMode:  cfg.PostgresSSLMode,
		Password: cfg.PostgresPassword,
	}, loggers)
	if err != nil {
		loggers.Fatalf("failed to initialize db: %s", err.Error())
		panic(err)
	}
	minio, err := store.MinioConnection(cfg, loggers)
	if err != nil {
		loggers.Fatal("error while connect to minio server", err)
		panic(err)
	}

	repos := repository.NewRepository(db, loggers)
	newStore := store.NewStore(minio, cfg, loggers)
	newService := service.NewService(repos, newStore, loggers)
	handlers := handler.NewHandler(newService, loggers)
	app := handlers.InitRoutes()
	// Start server (with or without graceful shutdown).
	port := fmt.Sprintf(":%d", cfg.ServerPort)
	err = app.Run(port)
	if err != nil {
		loggers.Fatal("error while running server", err)
	}
}
