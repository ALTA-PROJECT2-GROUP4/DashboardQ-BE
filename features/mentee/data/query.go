package data

import (
	"dashboardq-be/features/mentee"
	"errors"
	"log"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeData {
	return &menteeQuery{
		db: db,
	}
}

// AddMentee implements mentee.MenteeData
func (mq *menteeQuery) AddMentee(userID uint, newMentee mentee.Core) (mentee.Core, error) {
	cnvP := CoreToModel(newMentee)
	cnvP.UserID = userID

	err := mq.db.Create(&cnvP).Error
	if err != nil {
		log.Println("\tadd mentee query error: ", err.Error())
		return mentee.Core{}, errors.New("server problem")
	}

	return ModelToCore(cnvP), nil
}

// ShowAll implements mentee.MenteeData
func (mq *menteeQuery) ShowAll() ([]mentee.Core, error) {
	allMentee := []mentee.Core{}
	err := mq.db.Raw("SELECT m.id, m.user_id, u.name user_name, m.name, name, class, status, category, gender FROM mentees m JOIN users u ON u.id = p.user_id WHERE p.deleted_at IS NULL").Scan(&allMentee).Error
	if err != nil {
		log.Println("\terror query get all mentee: ", err.Error())
		return []mentee.Core{}, err
	}

	return allMentee, nil
}
