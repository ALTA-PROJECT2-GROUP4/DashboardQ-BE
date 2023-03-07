package router

import (
	"dashboardq-be/app/config"
	userData "dashboardq-be/features/users/data"
	userHdl "dashboardq-be/features/users/handler"
	userServ "dashboardq-be/features/users/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := userData.New(db)
	userService := userServ.New(userData)
	userHandler := userHdl.New(userService)

	// LOGIN
	e.POST("/login", userHandler.Login())
	// ADMIN
	e.POST("/register", userHandler.Register(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/users", userHandler.ProfileAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/users", userHandler.ShowAllAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/users/:user_id", userHandler.UpdateAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users/:user_id", userHandler.Deactive(), middleware.JWT([]byte(config.JWTKey)))
	// USER
	e.GET("/profile", userHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/profile/:user_id", userHandler.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/profile/:user_id", userHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
}
