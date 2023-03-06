package handler

import "dashboardq-be/features/users"

type RegResp struct {
}

func ToRegResp(data users.Core) RegResp {
	return RegResp{}
}

type UserReponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Nip    string `json:"nip"`
	Gender string `json:"gender"`
	Team   string `json:"team"`
	Role   string `json:"role"`
}
