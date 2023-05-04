package cmd

import (
	"govportal/migration"
	"govportal/pkg/handlers"
	"govportal/pkg/logger"
	"govportal/pkg/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var success string = "Successfully Initiated the"

func Run() error {
	logger, err := logger.NewFileLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug(success, "Loger")
	//////////////////////////////////
	configDb := migration.NewConfDb(logger)
	db := configDb.InitDatabase(logger)
	//////////////////////////////////
	defer db.Close()
	logger.Debug(success, "Data Base")
	//////////////////////////////////
	configDb.CreateTables(db, logger)
	logger.Debug(success, "tables")

	// repo := repository.NewRepository(db)
	// logger.Debug("Successfully Initiated the Repository")

	// service := service.NewService(repo)

	router := gin.New()

	router.Use(gin.Recovery(), gindump.Dump(), middlewares.AuthMiddleware())

	//////////////////////////////////////////////////////

	userGroup := router.Group("/main")
	userGroup.Any("/home" /* middlewares.AuthMiddleware()*/, handlers.MainPage)
	userGroup.Any("/sign-up", handlers.SignUp)
	userGroup.Any("/sign-in", handlers.SignIn)
	userGroup.Any("/profile", handlers.Profile)

	//////////////////////////////////////////////////////

	userStatusAdmin := router.Group("/main/admin")
	userStatusAdmin.Any("/redact")
	userStatusAdmin.Any("/delete")

	//////////////////////////////////////////////////////

	userStatusStudent := router.Group("/main/student")
	userStatusStudent.Any("/profile/:{id}")

	//////////////////////////////////////////////////////

	router.Run()
	defer db.Close()

	//////////////////////////////////////////////////////

	return nil
}
