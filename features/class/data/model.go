package data

import (
	"dashboardq-be/features/class"
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

// mengubah dari struct core ke struct model
func CoreToModel(dataCore class.Core) Class {
	return Class{
		Model:      gorm.Model{},
		Name:       dataCore.Name,
		StartClass: dataCore.StartClass,
		EndClass:   dataCore.EndClass,
		IdUser:     dataCore.IdUser,
	}
}
