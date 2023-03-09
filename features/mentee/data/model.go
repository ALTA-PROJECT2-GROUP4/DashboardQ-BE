package data

import (
	feedback "dashboardq-be/features/feedback/data"
	"dashboardq-be/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name            string
	Phone           string
	Telegram        string
	Email           string
	Class           string
	Gender          string
	Address         string
	HomeAddress     string
	DateBirth       string
	Status          string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	Category        string
	Major           string
	Graduate        string
	UserID          uint
	FeedbackID      feedback.Feedback `gorm:"foreignkey:FeedbackID;association_foreignkey:ID"`
}

func ModelToCore(data Mentee) mentee.Core {
	return mentee.Core{
		ID:              data.ID,
		Name:            data.Name,
		Phone:           data.Phone,
		Telegram:        data.Telegram,
		Email:           data.Email,
		Class:           data.Class,
		Gender:          data.Gender,
		Address:         data.Address,
		HomeAddress:     data.HomeAddress,
		DateBirth:       data.DateBirth,
		Status:          data.Status,
		EmergencyName:   data.EmergencyName,
		EmergencyPhone:  data.EmergencyPhone,
		EmergencyStatus: data.EmergencyStatus,
		Category:        data.Category,
		Major:           data.Major,
		Graduate:        data.Graduate,
		UserID:          data.UserID,
	}
}

func CoreToModel(core mentee.Core) Mentee {
	return Mentee{
		Model:           gorm.Model{ID: core.ID},
		Name:            core.Name,
		Phone:           core.Phone,
		Telegram:        core.Telegram,
		Email:           core.Email,
		Class:           core.Class,
		Gender:          core.Gender,
		Address:         core.Address,
		HomeAddress:     core.HomeAddress,
		DateBirth:       core.DateBirth,
		Status:          core.Status,
		EmergencyName:   core.EmergencyName,
		EmergencyPhone:  core.EmergencyPhone,
		EmergencyStatus: core.EmergencyStatus,
		Category:        core.Category,
		Major:           core.Major,
		Graduate:        core.Graduate,
		UserID:          core.UserID,
	}
}
