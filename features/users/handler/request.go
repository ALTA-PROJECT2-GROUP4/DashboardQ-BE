package handler

import "dashboardq-be/features/users"

type RegisterReq struct {
	Name      string `json:"name" form:"name"`
	DateBirth string `json:"date_birth" form:"date_birth"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Address   string `json:"address" form:"address"`
	Gender    string `json:"gender" form:"gender"`
	Team      string `json:"team" form:"team"`
	Status    string `json:"status" form:"status"`
	Phone     string `json:"phone" form:"phone"`
}

type LoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ReqToCore(data interface{}) *users.Core {
	res := users.Core{}

	switch data.(type) {
	case RegisterReq:
		cnv := data.(RegisterReq)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Phone = cnv.Phone
		res.Gender = cnv.Gender
		res.Address = cnv.Address
		res.Password = cnv.Password
		res.DateBirth = cnv.DateBirth
		res.Status = cnv.Status
		res.Team = cnv.Team
	case LoginReq:
		cnv := data.(LoginReq)
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
