package users

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Name      string
	DateBirth string
	Email     string `validate:"omitempty,email"`
	Password  string `validate:"min=6,emitempty"`
	Address   string
	Gender    string
	Role      string
	Team      string
	Status    string
	Phone     string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	ShowAllAdm() echo.HandlerFunc
	ProfileAdm() echo.HandlerFunc
	UpdateAdm() echo.HandlerFunc
	ShowAll() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(token interface{}, newUser Core) (Core, error)
	Login(email, password string) (string, Core, error)
	ShowAllAdm() ([]Core, error)
	ProfileAdm(userID uint) (Core, error)
	ShowAll() ([]Core, error)
	Profile(token interface{}) (Core, error)
	UpdateAdm(token interface{}, userID uint, newUpdate Core) (Core, error)
	Update(token interface{}, newUpdate Core) (Core, error)
	Deactive(token interface{}, userID uint) error
}

type UserData interface {
	Register(adminID uint, newUser Core) (Core, error)
	Login(email string) (Core, error)
	ShowAllAdm() ([]Core, error)
	ProfileAdm(userID uint) (Core, error)
	UpdateAdm(adminID, userID uint, newUpdate Core) (Core, error)
	ShowAll() ([]Core, error)
	Profile(userID uint) (Core, error)
	Update(userID uint, newUser Core) (Core, error)
	Deactive(adminID, userID uint) error
}
