package data

import (
	"dashboardq-be/features/class"
	"errors"
	"log"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Create implements class.ClassData
func (cl *classQuery) Create(userID uint, newClass class.Core) (class.Core, error) {
	cln := CoreToModel(newClass)
	cln.IdUser = userID

	err := cl.db.Create(&cln).Error
	if err != nil {
		log.Println("\tcreate class query error: ", err.Error())
		return class.Core{}, errors.New("server problem")
	}

	return ModelToCore(cln), nil
}

// Delete implements class.ClassData
func (cl *classQuery) Delete(userID uint, classID uint) error {
	qry := cl.db.Where("user_id = ?", userID).Delete(&Class{}, classID)

	if aff := qry.RowsAffected; aff <= 0 {
		log.Println("\tno rows affected: data not found")
		return errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tdelete query error: ", err.Error())
		return err
	}

	return nil
}

// Show implements class.ClassData
func (cl *classQuery) Show(classID uint) (class.Core, error) {
	res := Class{}
	err := cl.db.Where("id = ?", classID).First(&res).Error
	if err != nil {
		log.Println("query err", err.Error())
		return class.Core{}, errors.New("account not found")
	}
	return ModelToCore(res), nil
}

// ShowAll implements class.ClassData
func (cl *classQuery) ShowAll() ([]class.Core, error) {
	showall := []Class{}
	err := cl.db.Where("role = ?", "user").Find(&showall).Error
	if err != nil {
		log.Println("data not found")
		return []class.Core{}, errors.New("data not found")
	}
	result := []class.Core{}
	for _, val := range showall {
		result = append(result, ModelToCore(val))
	}
	return result, nil
}

// Update implements class.ClassData
func (cl *classQuery) Update(userID, classID uint, newUpdate class.Core) (class.Core, error) {
	cln := CoreToModel(newUpdate)
	cln.IdUser = classID

	qry := cl.db.Where("id = ?", classID).Updates(&cln)
	if qry.RowsAffected <= 0 {
		log.Println("\tupdate class query error: data not found")
		return class.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("\tupdate class query error: ", err.Error())
		return class.Core{}, errors.New("not found")
	}

	return ModelToCore(cln), nil
}

func New(db *gorm.DB) class.ClassData {
	return &classQuery{
		db: db,
	}
}
