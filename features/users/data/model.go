package data

import (
	"dashboardq-be/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	DateBirth string
	Email     string `gorm:"unique"`
	Password  string
	Address   string
	Gender    string
	Role      string
	Team      string
	Status    string
	Phone     string
}

func ModelToCore(data User) users.Core {
	return users.Core{
		ID:        data.ID,
		Name:      data.Name,
		DateBirth: data.DateBirth,
		Email:     data.Email,
		Password:  data.Password,
		Address:   data.Address,
		Gender:    data.Gender,
		Role:      data.Role,
		Team:      data.Team,
		Status:    data.Status,
		Phone:     data.Phone,
	}
}

func CoreToModel(data users.Core) User {
	return User{
		Model:     gorm.Model{ID: data.ID},
		Name:      data.Name,
		DateBirth: data.DateBirth,
		Email:     data.Email,
		Password:  data.Password,
		Address:   data.Address,
		Gender:    data.Gender,
		Role:      data.Role,
		Team:      data.Team,
		Status:    data.Status,
		Phone:     data.Phone,
	}
}
