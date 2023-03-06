package handler

import "dashboardq-be/features/users"

type RegResp struct {
}

func ToRegResp(data users.Core) RegResp {
	return RegResp{}
}

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Role   string `json:"role"`
	Team   string `json:"team"`
	Phone  string `json:"phone"`
}

func ToResponse(data users.Core) UserResponse {
	return UserResponse{
		ID:     data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Phone:  data.Phone,
		Gender: data.Gender,
		Team:   data.Team,
		Role:   data.Role,
	}
}

type UpdateUserResp struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Team   string `json:"team"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func ToResponseUpd(data users.Core) UpdateUserResp {
	return UpdateUserResp{
		Name:   data.Name,
		Role:   data.Role,
		Email:  data.Email,
		Team:   data.Team,
		Status: data.Status,
	}
}
