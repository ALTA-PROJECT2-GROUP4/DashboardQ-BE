package data

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Status string
	Notes  string
}
