package services

import (
	"dashboardq-be/features/class"
	"dashboardq-be/helper"
	"errors"
	"log"
	"strings"
)

type classService struct {
	qry class.ClassData
}

// Create implements class.ClassService
func (cl *classService) Create(token interface{}, newClass class.Core) (class.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("\t error extract token create class")
		return class.Core{}, errors.New("user not found")
	}

	res, err := cl.qry.Create(uint(userId), newClass)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror create query in service: ", err.Error())
		return class.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements class.ClassService
func (cl *classService) Delete(token interface{}, classID uint) error {
	id := helper.ExtractToken(token)
	err := cl.qry.Delete(uint(id), classID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete class fail")
	}
	return nil
}

// Show implements class.ClassService
func (cl *classService) Show(classID uint) (class.Core, error) {
	res, err := cl.qry.Show(classID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return class.Core{}, errors.New(msg)
	}
	return res, nil
}

// ShowAll implements class.ClassService
func (cl *classService) ShowAll() ([]class.Core, error) {
	res, err := cl.qry.ShowAll()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []class.Core{}, errors.New(msg)
	}
	return res, nil
}

// Update implements class.ClassService
func (cl *classService) Update(token interface{}, classID uint, newUpdate class.Core) (class.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("\t error extract token add product")
		return class.Core{}, errors.New("user not found")
	}

	res, err := cl.qry.Update(uint(userId), classID, newUpdate)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product data not found"
		} else {
			msg = "server problem"
		}
		log.Println("\terror update data in service: ", err.Error())
		return class.Core{}, errors.New(msg)
	}

	return res, nil
}

func New(sr class.ClassData) class.ClassService {
	return &classService{
		qry: sr,
	}
}
