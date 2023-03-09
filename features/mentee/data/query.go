package data

import (
	"dashboardq-be/features/mentee"
	"errors"
	"log"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeData {
	return &menteeQuery{
		db: db,
	}
}

// AddMentee implements mentee.MenteeData
func (mq *menteeQuery) AddMentee(userID uint, newMentee mentee.Core) (mentee.Core, error) {
	newMentee.Role = "mentee"

	cnv := CoreToModel(newMentee)
	err := mq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return mentee.Core{}, errors.New("server error")
	}

	newMentee.ID = cnv.ID
	return newMentee, nil
}

// ShowAll implements mentee.MenteeData
func (mq *menteeQuery) ShowAll() ([]mentee.Core, error) {
	getall := []Mentee{}
	err := mq.db.Where("role = ?", "mentee").Find(&getall).Error
	if err != nil {
		log.Println("data not found")
		return []mentee.Core{}, errors.New("data not found")
	}
	result := []mentee.Core{}
	for _, val := range getall {
		result = append(result, ModelToCore(val))
	}
	return result, nil
}
