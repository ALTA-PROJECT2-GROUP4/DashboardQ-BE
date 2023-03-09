package db

import (
	"fmt"
	"log"

	"dashboardq-be/app/config"
	classData "dashboardq-be/features/class/data"
	userData "dashboardq-be/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(userData.User{})
	db.AutoMigrate(classData.Class{})
}
