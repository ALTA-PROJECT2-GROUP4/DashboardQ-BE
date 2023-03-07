package handler

import (
	"dashboardq-be/features/users"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv users.UserService
}

func New(srv users.UserService) users.UserHandler {
	return &userHandler{
		srv: srv,
	}
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
			"data":    ToResponse(res),
			"message": "login success",
			"token":   token,
		})
	}
}

// Profile implements users.UserHandler
func (uh *userHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// eID := c.Param("id")
		// userID, _ := strconv.Atoi(eID)
		res, err := uh.srv.Profile(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToProfileResponse(res),
			"message": "success show profile",
		})
	}
}

// ProfileAdm implements users.UserHandler
func (uh *userHandler) ProfileAdm() echo.HandlerFunc {
	return func(c echo.Context) error {
		eID := c.Param("id")
		userID, _ := strconv.Atoi(eID)
		res, err := uh.srv.ProfileAdm(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToProfileResponse(res),
			"message": "success show profile",
		})
	}
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
func (uh *userHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uh.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		result := []ShowAllEmployee{}
		for _, val := range res {
			result = append(result, ShowAllEmployeeJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all users",
		})
	}
}

// ShowAllAdm implements users.UserHandler
func (*userHandler) ShowAllAdm() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements users.UserHandler
func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterReq{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		res, err := uh.srv.Update(c.Get("user"), *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already used"})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "access denied") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "access denied"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "not registered") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})
			}
		}

		result, err := ConvertEmployeeUpdateResponse(res)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			// log.Println(res)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    result,
				"message": "success update user profile",
			})
		}

	}
}

// UpdateAdm implements users.UserHandler
func (*userHandler) UpdateAdm() echo.HandlerFunc {
	panic("unimplemented")
}
