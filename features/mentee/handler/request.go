package handler

import "dashboardq-be/features/mentee"

type AddMenteeReq struct {
	Name            string `json:"name" form:"name"`
	Phone           string `json:"phone" form:"phobe"`
	Telegram        string `json:"telegram" form:"telegram"`
	Email           string `json:"email" form:"email"`
	Class           string `json:"class" form:"class"`
	Gender          string `json:"gender" form:"gender"`
	Address         string `json:"address" form:"address"`
	HomeAddress     string `json:"home_address" form:"home_address"`
	DateBirth       string `json:"date_birth" form:"date_birth"`
	Status          string `json:"status" form:"status"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Category        string `json:"category" form:"category"`
	Major           string `json:"major" form:"major"`
	Graduate        string `json:"graduate" form:"graduate"`
}

func ReqToCore(data interface{}) *mentee.Core {
	res := mentee.Core{}
	switch data.(type) {
	case AddMenteeReq:
		cnv := data.(AddMenteeReq)
		res.Name = cnv.Name
		res.Phone = cnv.Phone
		res.Telegram = cnv.Telegram
		res.Email = cnv.Email
		res.Class = cnv.Class
		res.Gender = cnv.Gender
		res.Address = cnv.Address
		res.HomeAddress = cnv.HomeAddress
		res.DateBirth = cnv.DateBirth
		res.Status = cnv.Status
		res.EmergencyName = cnv.EmergencyName
		res.EmergencyPhone = cnv.EmergencyPhone
		res.EmergencyStatus = cnv.EmergencyStatus
		res.Category = cnv.Category
		res.Major = cnv.Major
		res.Graduate = cnv.Graduate
	default:
		return nil
	}
	return &res
}
