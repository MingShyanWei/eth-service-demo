package models

type Transaction struct {
	TxHash string `gorm:"primary_key;type:varchar(67)`
	Num    uint64 `gorm:"index"` // TODO: index_type as hash
	From   string `gorm:"type:varchar(43)"`
	To     string `gorm:"type:varchar(43)"`
	Nonce  uint64
	Data   string
	Value  string
}
