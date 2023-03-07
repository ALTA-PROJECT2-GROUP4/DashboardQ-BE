package data

import (
	user "dashboardq-be/features/users/data"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name       string
	StartClass string
	EndClass   string
	IdUser     uint
	UserName   string
	User       user.User
}
