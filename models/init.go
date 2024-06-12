package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:12345678@tcp(localhost:3306)/iot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	err = db.AutoMigrate(&UserBasic{}, &ProductBasic{}, &DeviceBasic{})
	if err != nil {
		log.Fatalln("[DB ERROR] :", err)
	}
	DB = db
}
