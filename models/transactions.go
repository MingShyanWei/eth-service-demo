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
	if err = db.GetDb().First(&transactions, "tx_hash = ?", txHash).Error; err != nil {
		return
	}
	return
}
