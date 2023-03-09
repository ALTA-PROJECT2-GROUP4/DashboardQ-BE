package handler

import (
	"dashboardq-be/features/class"
)

type UserClassResp struct {
	UserName string `json:"user_name" form:"user_name"`
}
type ShowClassResp struct {
	ID         uint          `json:"id" form:"id"`
	Name       string        `json:"name" form:"name"`
	StartClass string        `json:"start_class" form:"start_class"`
	EndClass   string        `json:"end_class" form:"end_class"`
	IdUser     uint          `json:"id_user" form:"id_user"`
	User       UserClassResp `json:"user"`
}

func CoreToShowAllClassResp(data []class.Core) []ShowClassResp {
	res := []ShowClassResp{}
	for _, val := range data {
		res = append(res, CoreToShowClassResp(val))
	}
	return res
}

func CoreToShowClassResp(data class.Core) ShowClassResp {
	return ShowClassResp{
		ID:         data.ID,
		Name:       data.Name,
		StartClass: data.StartClass,
		EndClass:   data.EndClass,
		IdUser:     data.IdUser,
		User:       UserClassResp{UserName: data.User.Name},
	}
}
