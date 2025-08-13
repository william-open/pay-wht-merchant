package biz

import (
	"github.com/shopspring/decimal"
	"likeadmin/util"
)

// AddressAmount 钱包余额实体
type AddressAmount struct {
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
