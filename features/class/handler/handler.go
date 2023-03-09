package handler

import (
	"dashboardq-be/features/class"
	"dashboardq-be/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type classHandler struct {
	hd class.ClassService
}

// Create implements class.ClassHandler
func (cl *classHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := CreateReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("\tbind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
			// return c.JSON(http.StatusBadRequest, "wrong input")
		}

		res, err := cl.hd.Create(token, *CrtToCore(input))
		if err != nil {
			log.Println("\terror running create class service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			// return c.JSON(http.StatusInternalServerError, "server problem")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    CoreToShowClassResp(res),
			"message": "success create class",
		})
	}
}

// Delete implements class.ClassHandler
func (cl *classHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id_class")
		cln, err := strconv.Atoi(input)
		if err != nil {
			log.Println("\tRead param error: ", err.Error())
			return c.JSON(http.StatusBadRequest, "wrong class id parameter")
		}

		err = cl.hd.Delete(token, uint(cln))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("error calling delete class service: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("class not found"))
			} else {
				log.Println("error calling delete class service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete product",
		})
	}
}

// Show implements class.ClassHandler
func (cl *classHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := c.Param("id_class")
		cln, err := strconv.Atoi(input)
		if err != nil {
			log.Println("\tRead param error: ", err.Error())
			return c.JSON(http.StatusBadRequest, "wrong class id parameter")
		}

		res, err := cl.hd.Show(uint(cln))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("error calling show class service: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("class not found"))
			} else {
				log.Println("error calling show class service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    CoreToShowClassResp(res),
			"message": "success show class",
		})
	}
}

// ShowAll implements class.ClassHandler
func (cl *classHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cl.hd.ShowAll()
		if err != nil {
			log.Println("error running ShowAll Class service: ", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    CoreToShowAllClassResp(res),
			"message": "success show all class",
		})
	}
}

// Update implements class.ClassHandler
func (cl *classHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		classId := c.Param("id_class")
		cClassId, _ := strconv.Atoi(classId)

		input := CreateReq{}
		err := c.Bind(&input)
		if err != nil {
			log.Println("bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		res, err := cl.hd.Update(token, uint(cClassId), *CrtToCore(input))
		if err != nil {
			log.Println("\terror running update post service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    CoreToShowClassResp(res),
			"message": "success update class",
		})
	}
}

func New(hd class.ClassService) class.ClassHandler {
	return &classHandler{
		hd: hd,
	}
}
