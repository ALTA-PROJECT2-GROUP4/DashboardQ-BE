package data

import (
	"dashboardq-be/features/users"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}


func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

// Deactive implements users.UserData
func (*userQuery) Deactive(adminID uint, userID uint) error {
	panic("unimplemented")
}

// Login implements users.UserData
func (uq *userQuery) Login(email string) (users.Core, error) {
	if email == "" {
		log.Println("data empty, query error")
		return users.Core{}, errors.New("email not allowed empty")
	}
	res := User{}
	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return users.Core{}, errors.New("data not found")
	}

	return ModelToCore(res), nil
}

// Profile implements users.UserData
func (*userQuery) Profile(userID uint) (users.Core, error) {
	panic("unimplemented")
}

// ProfileAdm implements users.UserData
func (*userQuery) ProfileAdm(userID uint) (users.Core, error) {
	panic("unimplemented")
}

// Register implements users.UserData
func (uq *userQuery) Register(adminID uint, newUser users.Core) (users.Core, error) {
	if adminID != 1 {
		log.Println("access denied")
		return users.Core{}, errors.New("access denied")
	}

	dupEmail := CoreToModel(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return users.Core{}, errors.New("email duplicated")
	}

	newUser.Role = "user"

	cnv := CoreToModel(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return users.Core{}, errors.New("server error")
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

// ShowAll implements users.UserData
func (uq *userQuery) ShowAll() ([]users.Core, error) {
	getall := []User{}
	err := uq.db.Where("role = ?", "users").Find(&getall).Error
	if err != nil {
		log.Println("data not found")
		return []users.Core{}, errors.New("data not found")
	}
	result := []users.Core{}
	for _, val := range getall {
		result = append(result, ModelToCore(val))
	}
	return result, nil
}

// ShowAllAdm implements users.UserData
func (*userQuery) ShowAllAdm() ([]users.Core, error) {
	panic("unimplemented")
}

// Update implements users.UserData
func (*userQuery) Update(userID uint, newUser users.Core) (users.Core, error) {
	panic("unimplemented")
}

// UpdateAdm implements users.UserData
func (*userQuery) UpdateAdm(adminID uint, userID uint, newUpdate users.Core) (users.Core, error) {
	panic("unimplemented")
}

