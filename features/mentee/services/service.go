package services

import (
	"dashboardq-be/features/mentee"
	"dashboardq-be/helper"
	"errors"
	"log"
	"strings"
)

type menteeService struct {
	qry mentee.MenteeData
}

func New(data mentee.MenteeData) mentee.MenteeService {
	return &menteeService{
		qry: data,
	}
}

// AddMentee implements mentee.MenteeService
func (ms *menteeService) AddMentee(token interface{}, newMentee mentee.Core) (mentee.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("\t error extract token add mentee")
		return mentee.Core{}, errors.New("user not found")
	}

	res, err := ms.qry.AddMentee(uint(userId), newMentee)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror add query in service: ", err.Error())
		return mentee.Core{}, errors.New(msg)
	}

	return res, nil
}

// ShowAll implements mentee.MenteeService
func (ms *menteeService) ShowAll() ([]mentee.Core, error) {
	res, err := ms.qry.ShowAll()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []mentee.Core{}, errors.New(msg)
	}
	return res, nil
}
