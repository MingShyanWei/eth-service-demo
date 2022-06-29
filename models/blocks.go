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

type BlockDetail struct {
	Num          uint64 `gorm:"primary_key;autoIncrement:false"`
	Hash         string `gorm:"type:varchar(66);unique_index"`
	ParentHash   string `gorm:"type:varchar(66);unique_index"`
	Time         uint64
	Transactions []*Transaction
}

func (block *Block) ListBlocks(limit int) (blocks []Block, err error) {
	if err = db.GetDb().Order("Num DESC").Limit(limit).Find(&blocks).Error; err != nil {
		return
	}
	return
}

func (blockDetail *BlockDetail) GetBlockDetail(num int64) (blockDetails BlockDetail, err error) {
	// Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	// .Where("credit_cards.number = ?", "411111111111").Find(&user)
	//

	if err = db.GetDb().Table("blocks").Select("num, hash, parent_hash, time, transactions.tx_hash").Joins("join transactions on transactions.num = blocks.num").Where("blocks.num = ?", num).Find(&blockDetails).Error; err != nil {
		return
	}
	return
}
