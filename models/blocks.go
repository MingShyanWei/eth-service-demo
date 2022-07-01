package models

import (
	db "eth-service-demo/database"
)

type Block struct {
	Num        uint64 `gorm:"primary_key;autoIncrement:false" json:"block_num"`
	Hash       string `gorm:"type:varchar(66);unique_index" json:"block_hash"`
	ParentHash string `gorm:"type:varchar(66);unique_index" json:"parent_hash"`
	Time       uint64 `json:"block_time"`
}

type BlockWithTranscations struct {
	Num        uint64   `gorm:"primary_key;autoIncrement:false" json:"block_num"`
	Hash       string   `gorm:"type:varchar(66);unique_index" json:"block_hash"`
	ParentHash string   `gorm:"type:varchar(66);unique_index" json:"parent_hash"`
	Time       uint64   `json:"block_time"`
	TxHashs    []string `json:"transcations"`
}

func (block *Block) ListBlocks(limit int) (blocks []Block, err error) {
	// SELECT * FROM `blocks` ORDER BY Num DESC LIMIT 10
	if err = db.GetDb().Debug().Order("Num DESC").Limit(limit).Find(&blocks).Error; err != nil {
		return
	}
	return
}

func (blockWithTranscations *BlockWithTranscations) GetBlockDetail(num int64) (blocks BlockWithTranscations, err error) {

	// SELECT * FROM `blocks` WHERE `blocks`.`num` = 20597939
	err = db.GetDb().Debug().Model(&Block{}).Find(&blocks, num).Error
	if err != nil {
		return
	}
	// SELECT `tx_hash` FROM `transactions` WHERE num = 20597939s
	err = db.GetDb().Debug().Model(&Transaction{}).Select("tx_hash").Where("num = ?", num).Find(&blocks.TxHashs).Error

	return
}
