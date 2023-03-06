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
	Gender    string
	Role      string
	Team      string
	Status    string
	Phone     string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	ShowAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(token interface{}, newUser Core) (Core, error)
	Login(email, password string) (string, Core, error)
	ShowAll()
}
type UserData interface {}