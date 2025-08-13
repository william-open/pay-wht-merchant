package biz

import (
	"github.com/shopspring/decimal"
	"likeadmin/util"
)

// Address 交易实体
type Address struct {
	AddressId       int `gorm:"primaryKey;autoIncrement"`
	MId             int
	ChainSymbol     string
	Address         string
	privateKey      string
	TrxAmount       decimal.Decimal
	Trc20UsdtAmount decimal.Decimal
	Status          string `gorm:"default:0"`
	CreateBy        string
	CreateTime      util.Datetime `gorm:"autoCreateTime"`
	UpdateBy        string
	UpdateTime      util.Datetime `gorm:"autoUpdateTime"`
	Remark          string
}
