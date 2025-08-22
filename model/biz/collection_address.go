package biz

import (
	"mwhtpay/util"
)

// CollectionAddress 归集账户实体
type CollectionAddress struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	MId         uint
	ChainSymbol string
	Address     string
	MinAmount   string
	MaxAmount   string
	Status      uint8 `gorm:"default:1"`
	CreateBy    string
	CreateTime  util.Datetime `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  util.Datetime `gorm:"autoUpdateTime"`
	Remark      string
}
