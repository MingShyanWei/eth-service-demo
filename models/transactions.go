package models

import db "eth-service-demo/database"

type Transaction struct {
	TxHash          string           `gorm:"primary_key;type:varchar(67)" json:"tx_hash"`
	Num             uint64           `gorm:"index" json:"-"`
	From            string           `gorm:"type:varchar(43)" json:"from"`
	To              string           `gorm:"type:varchar(43)" json:"to"`
	Nonce           uint64           `json:"nonce"`
	Data            string           `json:"data"`
	Value           string           `json:"value"`
	TransactionLogs []TransactionLog `gorm:"-" json:"logs"`
}

type TransactionLog struct {
	TxHash string `gorm:"primary_key;type:varchar(67)" json:"-"`
	Index  uint   `gorm:"primary_key;autoIncrement:false" json:"index"`
	Data   string `json:"data"`
}

func (transaction *Transaction) GetTransaction(txHash string) (transactions Transaction, err error) {

	// SELECT * FROM `transactions` WHERE tx_hash = '0x1d411fbc5cf27d0b50f21daacba3ee70d27fff30772390abdddfae93b2d57e55'
	err = db.GetDb().Debug().Model(&Transaction{}).Where("tx_hash = ?", txHash).Find(&transactions).Error
	if err != nil {
		return
	}
	// SELECT * FROM `transaction_logs` WHERE tx_hash = '0x1d411fbc5cf27d0b50f21daacba3ee70d27fff30772390abdddfae93b2d57e55'
	err = db.GetDb().Debug().Model(&TransactionLog{}).Where("tx_hash = ?", txHash).Find(&transactions.TransactionLogs).Error

	return
}
