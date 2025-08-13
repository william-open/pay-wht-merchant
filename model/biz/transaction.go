package biz

import (
	"github.com/shopspring/decimal"
	"likeadmin/util"
)

// Transaction 交易实体
type Transaction struct {
	Id           int `gorm:"primaryKey;autoIncrement"`
	MId          int
	CurrencyType string
	Symbol       string
	BizNo        string
	BlockNum     string
	TxId         string
	FromAddress  string
	ToAddress    string
	Amount       decimal.Decimal
	Status       string `gorm:"default:0"`
	CreateBy     string
	CreateTime   util.Datetime `gorm:"autoCreateTime"`
	UpdateBy     string
	UpdateTime   util.Datetime `gorm:"autoUpdateTime"`
	Remark       string
}
