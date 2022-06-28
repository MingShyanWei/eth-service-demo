package models

import (
	db "eth-service-demo/database"
)

type Block struct {
	Num        uint64 `gorm:"primary_key;autoIncrement:false"`
	Hash       string `gorm:"type:varchar(66);unique_index"`
	ParentHash string `gorm:"type:varchar(66);unique_index"`
	Time       uint64
}

var Blocks []Block

func (block *Block) ListBlocks(limit int) (blocks []Block, err error) {
	if err = db.GetDb().Order("Num DESC").Limit(limit).Find(&blocks).Error; err != nil {
		return
	}
	return
}

func (block *Block) GetBlock(num int64) (blocks Block, err error) {
	if err = db.GetDb().Joins("join transactions on transactions.num = blocks.num").Find(&blocks, num).Error; err != nil {
		return
	}
	return
}
