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
func (*userService) Deactive(token interface{}, userID uint) error {
	panic("unimplemented")
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
func (*userService) Profile(token interface{}) (users.Core, error) {
	panic("unimplemented")
}

// ProfileAdm implements users.UserService
func (*userService) ProfileAdm(userID uint) (users.Core, error) {
	panic("unimplemented")
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
func (*userService) ShowAll() ([]users.Core, error) {
	panic("unimplemented")
}

// ShowAllAdm implements users.UserService
func (*userService) ShowAllAdm() ([]users.Core, error) {
	panic("unimplemented")
}

// Update implements users.UserService
func (*userService) Update(token interface{}, newUpdate users.Core) (users.Core, error) {
	panic("unimplemented")
}

// UpdateAdm implements users.UserService
func (*userService) UpdateAdm(token interface{}, userID uint, newUpdate users.Core) (users.Core, error) {
	panic("unimplemented")
}
