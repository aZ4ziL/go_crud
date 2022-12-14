package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = GetDB()
	db.AutoMigrate(&Data{})
}

func GetDB() *gorm.DB {
	dsn := "root:rafi213fajri@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
