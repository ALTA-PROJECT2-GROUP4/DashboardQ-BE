package router

import (
	"dashboardq-be/app/config"
	userData "dashboardq-be/features/users/data"
	userHdl "dashboardq-be/features/users/handler"
	userServ "dashboardq-be/features/users/services"

	classData "dashboardq-be/features/class/data"
	classHndl "dashboardq-be/features/class/handler"
	classServ "dashboardq-be/features/class/services"

	menteeData "dashboardq-be/features/mentee/data"
	menteeHdl "dashboardq-be/features/mentee/handler"
	menteeServ "dashboardq-be/features/mentee/services"

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

	menteeData := menteeData.New(db)
	menteeService := menteeServ.New(menteeData)
	menteeHandler := menteeHdl.New(menteeService)

	// LOGIN
	e.POST("/login", userHandler.Login())

	// ADMIN
	e.POST("/register", userHandler.Register(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/users", userHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/users/:user_id", userHandler.UpdateAdm(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users/:user_id", userHandler.Deactive(), middleware.JWT([]byte(config.JWTKey)))

	// USER
	e.GET("/profile/:user_id", userHandler.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/profile", userHandler.Update(), middleware.JWT([]byte(config.JWTKey)))

	// MENTEE
	e.POST("/mentee", menteeHandler.AddMentee(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/mentee", menteeHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))

	// CLASS
	e.POST("/class", classHandler.Create(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/class/:id_class", classHandler.Show(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/class", classHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/class/:id_class", classHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/class/:id_class", classHandler.Delete(), middleware.JWT([]byte(config.JWTKey)))
}
