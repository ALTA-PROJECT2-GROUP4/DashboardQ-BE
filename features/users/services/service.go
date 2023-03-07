package services

import (
	"dashboardq-be/app/config"
	"dashboardq-be/features/users"
	"dashboardq-be/helper"
	"errors"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry users.UserData
}

func New(ud users.UserData) users.UserService {
	return &userService{
		qry: ud,
	}
}

// Deactive implements users.UserService
func (us *userService) Deactive(token interface{}, userID uint) error {
	id := helper.ExtractToken(token)
	err := us.qry.Deactive(uint(id), userID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

// Login implements users.UserService
func (us *userService) Login(email string, password string) (string, users.Core, error) {
	res, err := us.qry.Login(email)

	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "not found") {
			errmsg = err.Error()
		} else {
			errmsg = "server problem"
		}
		log.Println("error login query: ", err.Error())
		return "", users.Core{}, errors.New(errmsg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		log.Println("wrong password :", err.Error())
		return "", users.Core{}, errors.New("wrong password")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))

	return useToken, res, nil
}

// Profile implements users.UserService
func (us *userService) Profile(token interface{}) (users.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := us.qry.Profile(uint(userID))
	if err != nil {
		log.Println("data not found")
		return users.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// ProfileAdm implements users.UserService
func (us *userService) ProfileAdm(userID uint) (users.Core, error) {
	res, err := us.qry.ProfileAdm(uint(userID))
	if err != nil {
		log.Println("data not found")
		return users.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// Register implements users.UserService
func (us *userService) Register(token interface{}, newUser users.Core) (users.Core, error) {
	adminID := helper.ExtractToken(token)

	hashed := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)

	res, err := us.qry.Register(uint(adminID), newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else if strings.Contains(err.Error(), "access denied") {
			msg = "access denied"
		} else {
			msg = "server error"
		}
		return users.Core{}, errors.New(msg)
	}

	return res, nil
}

// ShowAll implements users.UserService
func (us *userService) ShowAll() ([]users.Core, error) {
	res, err := us.qry.ShowAll()
	if err != nil {
		log.Println("data not found", err.Error())
		return []users.Core{}, errors.New("data not found")
	}
	return res, nil
}

// ShowAllAdm implements users.UserService
func (us *userService) ShowAllAdm() ([]users.Core, error) {
	res, err := us.qry.ShowAllAdm()
	if err != nil {
		log.Println("data not found", err.Error())
		return []users.Core{}, errors.New("data not found")
	}
	return res, nil
}

// Update implements users.UserService
func (us *userService) Update(token interface{}, newUpdate users.Core) (users.Core, error) {
	userID := helper.ExtractToken(token)

	hashed := helper.GeneratePassword(newUpdate.Password)
	newUpdate.Password = string(hashed)

	res, err := us.qry.Update(uint(userID), newUpdate)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "account not registered"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email duplicated"
		} else if strings.Contains(err.Error(), "access denied") {
			msg = "access denied"
		}
		return users.Core{}, errors.New(msg)
	}
	return res, nil
}

// UpdateAdm implements users.UserService
func (us *userService) UpdateAdm(token interface{}, userID uint, newUpdate users.Core) (users.Core, error) {
	adminID := helper.ExtractToken(token)

	hashed := helper.GeneratePassword(newUpdate.Password)
	newUpdate.Password = hashed

	res, err := us.qry.UpdateAdm(uint(adminID), userID, newUpdate)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "access") {
			msg = "access denied"
		} else if strings.Contains(err.Error(), "not found") {
			msg = "account not registered"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email duplicated"
		} else {
			msg = err.Error()
		}
		return users.Core{}, errors.New(msg)
	}
	return res, nil
}
