package resp

import (
	"github.com/shopspring/decimal"
	"likeadmin/util"
)

// BizCurrencyResp 货币数据返回信息
type BizCurrencyResp struct {
	CurrencyId      uint          `json:"currencyId" structs:"currencyId"`           // 主键
	PId             uint          `json:"pId" structs:"pId"`                         // 上级
	Currency        string        `json:"currency" structs:"currency"`               // 货币名称
	Symbol          string        `json:"symbol" structs:"symbol"`                   // 货币符号
	Logo            string        `json:"logo" structs:"logo"`                       // 货币Logo
	ContractAddress string        `json:"contractAddress" structs:"contractAddress"` // 智能合约
	CurrencyType    string        `json:"currencyType" structs:"currencyType"`       // 货币类型
	Protocol        string        `json:"protocol" structs:"protocol"`               // 货币协议
	ChainSymbol     string        `json:"chainSymbol" structs:"chainSymbol"`         // 货币符号
	Precision       uint          `json:"precision" structs:"precision"`             // 精度
	Status          uint8         `json:"status" structs:"status"`                   // 状态: [0=正常, 1=禁用]
	CreateTime      util.Datetime `json:"createTime" structs:"createTime"`           // 创建时间
	UpdateTime      util.Datetime `json:"updateTime" structs:"updateTime"`           // 更新时间
}

// BizMCurrencyDelReq 商户货币数据删除参数
type BizMCurrencyDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}

// BizTransactionResp 交易数据返回信息
type BizTransactionResp struct {
	Id           uint            `json:"id" structs:"id"`                     // 主键
	MId          uint            `json:"mId" structs:"mId"`                   //商户ID
	CurrencyType string          `json:"currencyType" structs:"currencyType"` //货币类型
	Symbol       string          `json:"symbol" structs:"symbol"`             //货币符号
	BizNo        string          `json:"bizNo" structs:"bizNo"`               //业务流水号
	BlockNum     string          `json:"blockNum" structs:"blockNum"`         //区块高度
	TxId         string          `json:"txId" structs:"txId"`                 //TXID
	FromAddress  string          `json:"fromAddress" structs:"fromAddress"`   //发起地址
	ToAddress    string          `json:"toAddress" structs:"toAddress"`       //接收地址
	Amount       decimal.Decimal `json:"amount" structs:"amount"`             //交易数量
	Status       uint8           `json:"status" structs:"status"`             // 状态: [0=正常, 1=禁用]
	CreateTime   util.Datetime   `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   util.Datetime   `json:"updateTime" structs:"updateTime"`     // 更新时间
}

// BizAddressResp 钱包地址数据返回信息
type BizAddressResp struct {
	AddressId   uint          `json:"id" structs:"id"`                   // 主键
	ChainSymbol string        `json:"chainSymbol" structs:"chainSymbol"` //主链符号
	Address     string        `json:"address" structs:"address"`         //钱包地址
	Status      uint8         `json:"status" structs:"status"`           // 状态: [0=正常, 1=禁用]
	CreateTime  util.Datetime `json:"createTime" structs:"createTime"`   // 创建时间
	UpdateTime  util.Datetime `json:"updateTime" structs:"updateTime"`   // 更新时间
}

// BizAddressBalanceResp 钱包地址余额数据返回信息
type BizAddressBalanceResp struct {
	CurrencySymbol string          `json:"currencySymbol" structs:"currencySymbol"` // 货币符号
	Amount         decimal.Decimal `json:"amount" structs:"amount"`                 //钱包地址
	UpdateTime     util.Datetime   `json:"updateTime" structs:"updateTime"`         // 更新时间
}

// BizCollectionResp 归集数据返回信息
type BizCollectionResp struct {
	Id          uint            `json:"id" structs:"id"`                   // 主键
	Symbol      string          `json:"symbol" structs:"symbol"`           //货币符号
	BizNo       string          `json:"bizNo" structs:"bizNo"`             //流水号
	BlockNum    string          `json:"blockNum" structs:"blockNum"`       //区块高度
	TxId        string          `json:"txId" structs:"txId"`               //TXID
	FromAddress string          `json:"fromAddress" structs:"fromAddress"` //发起地址
	ToAddress   string          `json:"toAddress" structs:"toAddress"`     //接收地址
	Amount      decimal.Decimal `json:"amount" structs:"amount"`           //归集数量
	Status      uint8           `json:"status" structs:"status"`           // 状态: [0=正常, 1=禁用]
	CreateTime  util.Datetime   `json:"createTime" structs:"createTime"`   // 创建时间
	UpdateTime  util.Datetime   `json:"updateTime" structs:"updateTime"`   // 更新时间
}

// BizCollectionAddressResp 归集账户数据返回信息
type BizCollectionAddressResp struct {
	Id          int           `json:"id" structs:"id"`
	MId         uint          `json:"mId" structs:"mId"`                 // 商户ID
	ChainSymbol string        `json:"chainSymbol" structs:"chainSymbol"` // 主链符号
	Address     string        `json:"address" structs:"address"`         // 钱包地址
	MinAmount   string        `json:"minAmount" structs:"minAmount"`     // 最小金额
	MaxAmount   string        `json:"maxAmount" structs:"maxAmount"`     // 最大金额
	Status      uint8         `json:"status" structs:"status"`           // 状态: [1=正常, 0=禁用]
	CreateTime  util.Datetime `json:"createTime" structs:"createTime"`   // 创建时间
	UpdateTime  util.Datetime `json:"updateTime" structs:"updateTime"`   // 更新时间
}

// BizAddressBalanceAmountResp 归集钱包地址余额数据返回信息
type BizAddressBalanceAmountResp struct {
	CurrencySymbol string          `json:"currencySymbol" structs:"currencySymbol"` // 货币符号
	Amount         decimal.Decimal `json:"amount" structs:"amount"`                 //钱包地址
	UpdateTime     util.Datetime   `json:"updateTime" structs:"updateTime"`         // 更新时间
}
