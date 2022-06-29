package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	// dsn := "root:kzy0RV0lte@tcp(192.168.17.104:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DB_CONNECTION")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}
}

func GetDb() *gorm.DB {
	return db
}
