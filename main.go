package main

import (
	config "dashboardq-be/app/config"
	database "dashboardq-be/app/database"
	router "dashboardq-be/app/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)

	database.Migrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(db, e)

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
