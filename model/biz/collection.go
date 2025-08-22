package biz

import (
	"github.com/shopspring/decimal"
	"mwhtpay/util"
)

// Collection 归集实体
type Collection struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	MId         int
	Symbol      string
	BizNo       string
	BlockNum    string
	TxId        string
	FromAddress string
	ToAddress   string
	Amount      decimal.Decimal
	Status      string `gorm:"default:0"`
	CreateBy    string
	CreateTime  util.Datetime `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  util.Datetime `gorm:"autoUpdateTime"`
	Remark      string
}
