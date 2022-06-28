package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Block struct {
	Num        uint64 `gorm:"primary_key;autoIncrement:false"`
	Hash       string `gorm:"type:varchar(66);unique_index"`
	ParentHash string `gorm:"type:varchar(66);unique_index"`
	Time       uint64
}

type Transaction struct {
	TxHash string `gorm:"primary_key;type:varchar(67)`
	Num    uint64 `gorm:"index"` // TODO: index_type as hash
	From   string `gorm:"type:varchar(43)"`
	To     string `gorm:"type:varchar(43)"`
	Nonce  uint64
	Data   string
	value  string
}

func main() {
	dsn := "root:kzy0RV0lte@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Block{}, &Transaction{})
}
