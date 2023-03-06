package users

import (
	// "github.com/labstack/echo/v4"
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

type UserHandler interface {}
type UserService interface {}
type UserData interface {}