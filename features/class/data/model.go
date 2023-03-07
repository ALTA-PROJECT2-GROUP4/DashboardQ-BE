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

// mengubah dari struct model ke struct core
func ModelToCore(dataModel Class) class.Core {
	return class.Core{
		ID:         dataModel.ID,
		Name:       dataModel.Name,
		StartClass: dataModel.StartClass,
		EndClass:   dataModel.EndClass,
		IdUser:     dataModel.User.ID,
	}
}

// mengubah dari list struct model ke struct core
func ListModelToCore(dataModel []Class) []class.Core {
	var dataCore []class.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
