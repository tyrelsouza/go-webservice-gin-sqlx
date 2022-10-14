package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Obviously change this to your settings and get securely.
	dsn := "mysql:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	_ = db.AutoMigrate(&Album{})

	DB = db
}
