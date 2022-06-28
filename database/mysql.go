package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := "root:kzy0RV0lte@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
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
