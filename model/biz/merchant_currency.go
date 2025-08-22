package biz

import "mwhtpay/util"

// MerchantCurrency 商户货币实体
type MerchantCurrency struct {
	CurrencyId uint `gorm:"comment:'货币ID'"`
	MId        uint `gorm:"comment:'商户ID'"`
	CreateBy   string
	UpdateBy   string
	CreateTime util.Datetime `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime util.Datetime `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime util.Datetime `gorm:"not null;default:0;comment:'删除时间'"`
	Remark     string        `gorm:"not null;default:'';comment:'备注信息'"`
}
