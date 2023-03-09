package handler

import (
	"dashboardq-be/features/mentee"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type menteeHandler struct {
	srv mentee.MenteeService
}

func New(srv mentee.MenteeService) mentee.MenteeHandler {
	return &menteeHandler{
		srv: srv,
	}
}

// AddMentee implements mentee.MenteeHandler
func (mh *menteeHandler) AddMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddMenteeReq{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		res, err := mh.srv.AddMentee(c.Get("mentee"), *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "access") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create mentee"})
	}
}

// ShowAll implements mentee.MenteeHandler
func (mh *menteeHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := mh.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		result := []ShowAllMentee{}
		for _, val := range res {
			result = append(result, ShowAllMenteeJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all mentees",
		})
	}
}
