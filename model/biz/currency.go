package biz

import "mwhtpay/util"

// Currency 货币实体
type Currency struct {
	CurrencyId      uint   `gorm:"primaryKey;autoIncrement;comment:'主键'"`
	PId             uint   `gorm:"not null;default:0;comment:'父级ID'"`
	Currency        string `gorm:"not null;default:'';comment:'货币名称'"`
	Symbol          string `gorm:"not null;default:'';comment:'货币符号'"`
	Logo            string `gorm:"not null;default:'';comment:'货币Logo'"`
	ContractAddress string `gorm:"not null;default:'';comment:'智能合约地址'"`
	CurrencyType    string `gorm:"not null;default:'';comment:'货币类型'"`
	Protocol        string `gorm:"not null;default:'';comment:'协议'"`
	Decimals        uint
	Status          uint8 `gorm:"not null;default:0;comment:'字典状态: 0=正常, 1=禁用'"`
	CreateBy        string
	UpdateBy        string
	CreateTime      util.Datetime `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime      util.Datetime `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime      util.Datetime `gorm:"not null;default:0;comment:'删除时间'"`
	Remark          string        `gorm:"not null;default:'';comment:'备注信息'"`
}
