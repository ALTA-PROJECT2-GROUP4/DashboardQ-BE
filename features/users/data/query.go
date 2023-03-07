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
func (uq *userQuery) Profile(userID uint) (users.Core, error) {
	if userID == 1 {
		log.Println("cannot access admin data")
		return users.Core{}, errors.New("cannot access admin data")
	}
	res := User{}
	err := uq.db.Where("id = ?", userID).First(&res).Error
	if err != nil {
		log.Println("query err", err.Error())
		return users.Core{}, errors.New("account not found")
	}
	return ModelToCore(res), nil
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
func (uq *userQuery) Update(userID uint, newUpdate users.Core) (users.Core, error) {
	if userID == 1 {
		log.Println("access denied")
		return users.Core{}, errors.New("access denied")
	}
	if newUpdate.Email != "" {
		dupEmail := User{}
		err := uq.db.Where("email = ?", newUpdate.Email).First(&dupEmail).Error
		if err == nil {
			log.Println("duplicated")
			return users.Core{}, errors.New("email duplicated")
		}
	}
	cnv := CoreToModel(newUpdate)
	qry := uq.db.Model(&User{}).Where("id = ?", userID).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return users.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return users.Core{}, errors.New("user not found")
	}
	result := ModelToCore(cnv)
	result.ID = userID
	return result, nil
}

// UpdateAdm implements users.UserData
func (*userQuery) UpdateAdm(adminID uint, userID uint, newUpdate users.Core) (users.Core, error) {
	panic("unimplemented")
}

