package router

import (
	"dashboardq-be/app/config"
	userData "dashboardq-be/features/users/data"
	userHdl "dashboardq-be/features/users/handler"
	userServ "dashboardq-be/features/users/services"

	classData "dashboardq-be/features/class/data"
	classHndl "dashboardq-be/features/class/handler"
	classServ "dashboardq-be/features/class/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := userData.New(db)
	userService := userServ.New(userData)
	userHandler := userHdl.New(userService)

	classData := classData.New(db)
	classService := classServ.New(classData)
	classHandler := classHndl.New(classService)

	// LOGIN
	e.POST("/login", userHandler.Login())
	// ADMIN
	e.POST("/register", userHandler.Register(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/users/:user_id", userHandler.ProfileAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/users", userHandler.ShowAllAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/users/:user_id", userHandler.UpdateAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users/:user_id", userHandler.Deactive(), middleware.JWT([]byte(config.JWTKey)))
	// USER
	e.GET("/profile", userHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/profile/:user_id", userHandler.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/profile", userHandler.Update(), middleware.JWT([]byte(config.JWTKey)))

	// CLASS
	e.POST("/create", classHandler.Create(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/class:class_id", classHandler.Show(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/class", classHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/class/:class_id", classHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/class/:class_id", classHandler.Delete(), middleware.JWT([]byte(config.JWTKey)))
}
