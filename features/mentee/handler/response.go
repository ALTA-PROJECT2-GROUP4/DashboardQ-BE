package handler

import "dashboardq-be/features/mentee"

type ShowAllMentee struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Status   string `json:"status"`
	Category string `json:"category"`
	Gender   string `json:"gender"`
}

func ShowAllMenteeJson(data mentee.Core) ShowAllMentee {
	return ShowAllMentee{
		ID:       data.ID,
		Name:     data.Name,
		Class:    data.Class,
		Status:   data.Status,
		Category: data.Category,
		Gender:   data.Gender,
	}
}
