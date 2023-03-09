package mentee

import "github.com/labstack/echo/v4"

type Core struct {
	ID              uint
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
	Role            string
	UserID          uint
}

type MenteeHandler interface {
	AddMentee() echo.HandlerFunc
	ShowAll() echo.HandlerFunc
}

type MenteeService interface {
	AddMentee(token interface{}, newMentee Core) (Core, error)
	ShowAll() ([]Core, error)
}

type MenteeData interface {
	AddMentee(userID uint, newMentee Core) (Core, error)
	ShowAll() ([]Core, error)
}
