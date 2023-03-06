package router

import (
	// "dashboardq-be/app/config"
	userData "dashboardq-be/features/users/data"
	userHdl "dashboardq-be/features/users/handler"
	userServ "dashboardq-be/features/users/services"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := userData.New(db)
	userService := userServ.New(userData)
	userHandler := userHdl.New(userService)

	e.POST("/register", userHandler.Register())
	e.POST("/login", userHandler.Login())
	// e.PUT("/users", userHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
}
