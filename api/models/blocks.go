package models

import (
	orm "api/database"
)

type Block struct {
	Num        uint64 `gorm:"primary_key;autoIncrement:false"`
	Hash       string `gorm:"type:varchar(64);unique_index"`
	ParentHash string `gorm:"type:varchar(64);unique_index"`
	Time       uint64
}

var Blocks []Block

func (block *Block) ListBlocks(limit int) (blocks []Block, err error) {
	if err = orm.Db.Order("Num DESC").Limit(limit).Find(&blocks).Error; err != nil {
		return
	}
	return
}

func (block *Block) GetBlock(num int64) (blocks Block, err error) {
	if err = orm.Db.Joins("join transactions on transactions.num = blocks.num").Find(&blocks, num).Error; err != nil {
		return
	}
	return
}
