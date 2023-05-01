package cmd

import (
	"govportal/migration"
	"govportal/pkg/handlers"
	"govportal/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func Run() error {
	logger, err := logger.NewFileLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("Successfully Initiated the Loger")

	configDb := migration.NewConfDb(logger)
	db := configDb.InitDatabase(logger)

	defer db.Close()

	logger.Debug("Successfully Initiated the Data Base")

	configDb.CreateTables(db, logger)
	logger.Debug("Successfully created the tables")
	configDb.CreateTables(db, logger)
	logger.Debug("Successfully created the tables")

	// repo := repository.NewRepository(db)
	// logger.Debug("Successfully Initiated the Repository")

	router := gin.New()

	router.Use(gin.Recovery() /*middlewares.BasicAuth()*/, gindump.Dump())

	// router.GET("/", handlers.MainPage)
	// router.OPTIONS("/", handlers.MainPage)
	router.Any("/", handlers.MainPage)
	// router.POST("/", handlers.MainPagePOST)

	router.Any("/sign-up", handlers.SignUp)

	router.Any("/sign-in", handlers.SignIn)

	router.Run(":8080")
	defer db.Close()
	return nil
}
