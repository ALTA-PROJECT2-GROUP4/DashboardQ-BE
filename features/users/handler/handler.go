package handler

import (
	"dashboardq-be/features/users"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv users.UserService
}

// Deactive implements users.UserHandler
func (*userHandler) Deactive() echo.HandlerFunc {
	panic("unimplemented")
}

// Login implements users.UserHandler
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		if input.Email == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Email not allowed empty"})
		} else if input.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not allowed empty"})
		}

		token, res, err := uh.srv.Login(input.Email, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not match"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":             ToResponse(res),
			"message":          "login success",
			"token":            token,
		})
	}
}

// Profile implements users.UserHandler
func (*userHandler) Profile() echo.HandlerFunc {
	panic("unimplemented")
}

// ProfileAdm implements users.UserHandler
func (*userHandler) ProfileAdm() echo.HandlerFunc {
	panic("unimplemented")
}

// Register implements users.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterReq{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		res, err := uh.srv.Register(c.Get("user"), *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "access") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create user account"})
	}
}

// ShowAll implements users.UserHandler
func (*userHandler) ShowAll() echo.HandlerFunc {
	panic("unimplemented")
}

// ShowAllAdm implements users.UserHandler
func (*userHandler) ShowAllAdm() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements users.UserHandler
func (*userHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateAdm implements users.UserHandler
func (*userHandler) UpdateAdm() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv users.UserService) users.UserHandler {
	return &userHandler{
		srv: srv,
	}
}
