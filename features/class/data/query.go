package data

import (
	"dashboardq-be/features/class"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Create implements class.ClassData
func (*classQuery) Create(userID uint, newClass class.Core) (class.Core, error) {
	panic("unimplemented")
}

// Delete implements class.ClassData
func (*classQuery) Delete(userID uint, classID uint) error {
	panic("unimplemented")
}

// Show implements class.ClassData
func (*classQuery) Show(classID uint) (class.Core, error) {
	panic("unimplemented")
}

// ShowAll implements class.ClassData
func (*classQuery) ShowAll() ([]class.Core, error) {
	panic("unimplemented")
}

// Update implements class.ClassData
func (*classQuery) Update(classID uint, newClass class.Core) (class.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) class.ClassData {
	return &classQuery{
		db: db,
	}
}
