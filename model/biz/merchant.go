package biz

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"likeadmin/util"
)

// Merchant 商户实体
type Merchant struct {
	MId               int `gorm:"primaryKey;autoIncrement"`
	Username          string
	Password          string
	Nickname          string
	CallbackSecretKey string
	NotifyUrl         string
	AesSecretKey      string
	PublicKey         string
	PrivateKey        string
	AppId             string
	ApiKey            string
	Balance           decimal.Decimal
	Status            string `gorm:"default:0"`
	CreateBy          string
	CreateTime        util.Datetime `gorm:"autoCreateTime"`
	UpdateBy          string
	UpdateTime        util.Datetime `gorm:"autoUpdateTime"`
	DeleteTime        gorm.DeletedAt
	Remark            string
}

func (Merchant) TableName() string {
	return "w_merchant"
}
