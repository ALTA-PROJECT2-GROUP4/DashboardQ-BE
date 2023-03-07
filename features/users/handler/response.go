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

type ProfileResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	DateBirth string `json:"date_birth"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Team      string `json:"team"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func ToProfileResponse(data users.Core) ProfileResponse {
	return ProfileResponse{
		ID:        data.ID,
		Name:      data.Name,
		DateBirth: data.DateBirth,
		Email:     data.Email,
		Role:      data.Role,
		Gender:    data.Gender,
		Team:      data.Team,
		Phone:     data.Phone,
		Address:   data.Address,
	}
}

type UpdateResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" form:"name"`
	DateBirth string `json:"date_birth" form:"date_birth"`
	Email     string `json:"email" form:"email"`
	Gender    string `json:"gender" form:"gender"`
	Team      string `json:"team" form:"team"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Password  string `json:"password" form:"password"`
}

func ToUpdateResponse(data users.Core) UpdateResponse {
	return UpdateResponse{
		ID:        data.ID,
		Name:      data.Name,
		DateBirth: data.DateBirth,
		Email:     data.Email,
		Gender:    data.Gender,
		Team:      data.Team,
		Phone:     data.Phone,
		Address:   data.Address,
		Password:  data.Password,
	}
}

type UpdateResponseUser struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password" form:"password"`
}

func ToUpdateResponseUser(data users.Core) UpdateResponseUser {
	return UpdateResponseUser{
		ID:       data.ID,
		Email:    data.Email,
		Password: data.Password,
	}
}

type ShowAllUser struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Team  string `json:"team"`
	Role  string `json:"role"`
}

func ShowAllUserJson(data users.Core) ShowAllUser {
	return ShowAllUser{
		ID:    data.ID,
		Email: data.Email,
		Name:  data.Name,
		Team:  data.Team,
		Role:  data.Role,
	}
}
