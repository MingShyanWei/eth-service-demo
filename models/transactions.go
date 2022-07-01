package models

import db "eth-service-demo/database"

type Transaction struct {
	TxHash string `gorm:"primary_key;type:varchar(67) json:"tx_hash"`
	Num    uint64 `gorm:"index" json:"-"`
	From   string `gorm:"type:varchar(43)" json:"from"`
	To     string `gorm:"type:varchar(43)" json:"to"`
	Nonce  uint64 `json:"nonce"`
	Data   string `json:"data"`
	Value  string `json:"value"`
}

func (transaction *Transaction) GetTransaction(txHash string) (transactions Transaction, err error) {
	// SELECT * FROM `transactions` WHERE tx_hash = '0xd7d9c32699fabd278e9d5f1119d7bfcee07f778c1314940f511b5385e6b30c12'
	err = db.GetDb().Debug().Model(&Transaction{}).Where("tx_hash = ?", txHash).Find(&transactions).Error
	if err != nil {
		return
	}

	// // SELECT * FROM `blocks` WHERE `blocks`.`num` = 20597939
	// err = db.GetDb().Debug().Model(&Block{}).Find(&blocks, num).Error
	// if err != nil {
	// 	return
	// }
	// // SELECT `tx_hash` FROM `transactions` WHERE num = 20597939s
	// err = db.GetDb().Debug().Model(&Transaction{}).Select("tx_hash").Where("num = ?", num).Find(&blocks.TxHashs).Error

	return
}
