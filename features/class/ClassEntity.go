package class

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Name       string
	StartClass string
	EndClass   string
	IdUser     uint
}

type ClassHandler interface {
	Create() echo.HandlerFunc
	ShowAll() echo.HandlerFunc
	Show() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ClassService interface {
	Create(token interface{}, newClass Core) (Core, error)
	ShowAll() ([]Core, error)
	Show(token interface{}) (Core, error)
	Update(token interface{}, classID uint, newUpdate Core) (Core, error)
	Delete(token interface{}, classID uint) error
}

type ClassData interface {
	Create(userID uint, newClass Core) (Core, error)
	ShowAll() ([]Core, error)
	Show(classID uint) (Core, error)
	Update(classID uint, newClass Core) (Core, error)
	Delete(userID, classID uint) error
}
