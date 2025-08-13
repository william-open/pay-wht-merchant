package biz

import (
	"github.com/shopspring/decimal"
	"likeadmin/util"
)

// CollectionAddressAmount 归集钱包账户实体
type CollectionAddressAmount struct {
	Id             int `gorm:"primaryKey;autoIncrement"`
	MId            int
	CurrencySymbol string
	Address        string
	Amount         decimal.Decimal
	CreateBy       string
	CreateTime     util.Datetime `gorm:"autoCreateTime"`
	UpdateBy       string
	UpdateTime     util.Datetime `gorm:"autoUpdateTime"`
	Remark         string
}
