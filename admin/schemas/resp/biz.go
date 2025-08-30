package resp

import (
	"github.com/shopspring/decimal"
	"mwhtpay/util"
	"time"
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

// BizChannelProductResp 通道产品返回信息
type BizChannelProductResp struct {
	Id           uint   `json:"id" structs:"id"`                     // 主键
	ChannelCode  string `json:"channelCode" structs:"channelCode"`   //通道编码
	ChannelTitle string `json:"channelTitle" structs:"channelTitle"` //通道标题
	OrderRange   string `json:"orderRange" structs:"orderRange"`     //金额范围
	DefaultRate  string `json:"defaultRate" structs:"defaultRate"`   //默认费率
	SingleFee    string `json:"singleFee" structs:"singleFee"`       //单笔费用
	Country      string `json:"country" structs:"country"`           //国家
	Type         string `json:"type" structs:"type"`                 //通道类型
}

// MerchantWhitelistResponse 商户白名单列表
type MerchantWhitelistResponse struct {
	ID         int           `json:"id"`
	MID        uint64        `json:"mId"`
	IPAddress  string        `json:"iPAddress"`
	CanAdmin   uint8         `json:"canAdmin"`
	CanPayout  uint8         `json:"canPayout"`
	CanReceive uint8         `json:"canReceive"`
	CreateTime util.Datetime `json:"createTime"`
	Remark     string        `json:"remark"`
}

// BizCollectOrderResp 收款订单返回信息
type BizCollectOrderResp struct {
	Id           uint   `json:"id" structs:"id"`                     // 主键
	ChannelCode  string `json:"channelCode" structs:"channelCode"`   //通道编码
	ChannelTitle string `json:"channelTitle" structs:"channelTitle"` //通道标题
	OrderRange   string `json:"orderRange" structs:"orderRange"`     //金额范围
	DefaultRate  string `json:"defaultRate" structs:"defaultRate"`   //默认费率
	SingleFee    string `json:"singleFee" structs:"singleFee"`       //单笔费用
	Country      string `json:"country" structs:"country"`           //国家
}

// 代收列表
type OrderReceiveListResponse struct {
	OrderID        string    `json:"orderId"`        // 全局唯一订单ID
	MID            uint64    `json:"mId"`            // 商户ID
	SupplierID     int64     `json:"supplierId"`     // 上游供应商ID
	MOrderId       string    `json:"mOrderId"`       // 商户订单号
	Amount         float64   `json:"amount"`         // 订单金额
	Fees           float64   `json:"fees"`           // 手续费
	PayAmount      float64   `json:"payAmount"`      // 实际支付金额
	RealMoney      float64   `json:"realMoney"`      // 实际到账金额
	FreezeAmount   float64   `json:"freezeAmount"`   // 冻结金额
	Currency       string    `json:"currency"`       // 货币代码
	NotifyURL      string    `json:"notifyUrl"`      // 异步回调通知URL
	ReturnURL      string    `json:"returnUrl"`      // 同步回调URL
	MDomain        string    `json:"mDomain"`        // 下单域名
	MIP            string    `json:"mIp"`            // 下单IP
	Title          string    `json:"title"`          // 订单标题
	MTitle         string    `json:"mTitle"`         // 商户名称
	ChannelCode    string    `json:"channelCode"`    // 通道编码
	ChannelTitle   string    `json:"channelTitle"`   // 通道名称
	UpChannelCode  string    `json:"upChannelCode"`  // 上游通道编码
	UpChannelTitle string    `json:"upChannelTitle"` // 上游通道名称
	MRate          string    `json:"mRate"`          // 商户费率
	UpRate         string    `json:"upRate"`         // 上游商户费率
	UpFixedFee     string    `json:"upFixedFee"`     // 上游通道固定费用
	MFixedFee      string    `json:"mFixedFee"`      // 商户通道固定费用
	Country        string    `json:"country"`        // 国家
	AccountNo      string    `json:"accountNo"`      // 付款人账号
	AccountName    string    `json:"accountName"`    // 付款人姓名
	PayEmail       string    `json:"payEmail"`       // 付款人邮箱
	PayPhone       string    `json:"payPhone"`       // 付款人手机号码
	BankCode       string    `json:"bankCode"`       // 付款人银行编码
	BankName       string    `json:"bankName"`       // 付款人银行名
	Status         int8      `json:"status"`         // 0:待支付,1:成功,2:失败,3:退款
	NotifyStatus   int8      `json:"notifyStatus"`   // 回调通知状态:0表示未回调，1表示回调成功，2回调失败
	UpOrderID      *uint64   `json:"upOrderId"`      // 上游交易订单ID
	ChannelID      int64     `json:"channelId"`      // 系统支付渠道ID
	UpChannelID    int64     `json:"upChannelId"`    // 上游通道ID
	NotifyTime     time.Time `json:"notifyTime"`     // 回调通知时间
	CreateTime     time.Time `json:"createTime"`     // 创建时间
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
	FinishTime     time.Time `json:"finishTime"`     // 完成时间
}

// 代付列表
type OrderPayoutListResponse struct {
	OrderID        string    `json:"orderId"`        // 全局唯一订单ID
	MID            uint64    `json:"mId"`            // 商户ID
	SupplierID     int64     `json:"supplierId"`     // 上游供应商ID
	MOrderId       string    `json:"mOrderId"`       // 商户订单号
	Amount         float64   `json:"amount"`         // 订单金额
	Fees           float64   `json:"fees"`           // 手续费
	PayAmount      float64   `json:"payAmount"`      // 实际支付金额
	RealMoney      float64   `json:"realMoney"`      // 实际到账金额
	FreezeAmount   float64   `json:"freezeAmount"`   // 冻结金额
	Currency       string    `json:"currency"`       // 货币代码
	NotifyURL      string    `json:"notifyUrl"`      // 异步回调通知URL
	ReturnURL      string    `json:"returnUrl"`      // 同步回调URL
	MDomain        string    `json:"mDomain"`        // 下单域名
	MIP            string    `json:"mIp"`            // 下单IP
	Title          string    `json:"title"`          // 订单标题
	MTitle         string    `json:"mTitle"`         // 商户名称
	ChannelCode    string    `json:"channelCode"`    // 通道编码
	ChannelTitle   string    `json:"channelTitle"`   // 通道名称
	UpChannelCode  string    `json:"upChannelCode"`  // 上游通道编码
	UpChannelTitle string    `json:"upChannelTitle"` // 上游通道名称
	MRate          string    `json:"mRate"`          // 商户费率
	UpRate         string    `json:"upRate"`         // 上游商户费率
	UpFixedFee     string    `json:"upFixedFee"`     // 上游通道固定费用
	MFixedFee      string    `json:"mFixedFee"`      // 商户通道固定费用
	Country        string    `json:"country"`        // 国家
	AccountNo      string    `json:"accountNo"`      // 付款人账号
	AccountName    string    `json:"accountName"`    // 付款人姓名
	PayEmail       string    `json:"payEmail"`       // 付款人邮箱
	PayPhone       string    `json:"payPhone"`       // 付款人手机号码
	BankCode       string    `json:"bankCode"`       // 付款人银行编码
	BankName       string    `json:"bankName"`       // 付款人银行名
	Status         int8      `json:"status"`         // 0:待支付,1:成功,2:失败,3:退款
	NotifyStatus   int8      `json:"notifyStatus"`   // 回调通知状态:0表示未回调，1表示回调成功，2回调失败
	UpOrderID      *uint64   `json:"upOrderId"`      // 上游交易订单ID
	ChannelID      int64     `json:"channelId"`      // 系统支付渠道ID
	UpChannelID    int64     `json:"upChannelId"`    // 上游通道ID
	NotifyTime     time.Time `json:"notifyTime"`     // 回调通知时间
	CreateTime     time.Time `json:"createTime"`     // 创建时间
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
	FinishTime     time.Time `json:"finishTime"`     // 完成时间
}

// BizGenGoogleCodeResp 谷歌验证码返回信息
type BizGenGoogleCodeResp struct {
	Secret string `json:"secret" structs:"secret"` //谷歌密钥
	Qrcode string `json:"qrcode" structs:"qrcode"` //谷歌密钥二维码
}

// BizDockingViewKeyResp 对接密钥返回信息
type BizDockingViewKeyResp struct {
	ApiKey string `json:"apiKey" structs:"apiKey"` //对接密钥
}

// BizDockingConfigInfoResp 对接配置信息
type BizDockingConfigInfoResp struct {
	AppId      string `json:"appId" structs:"appId"`           //应用ID/商户ID
	ApiDoc     string `json:"apiDoc" structs:"apiDoc"`         //API对接文档
	ApiGateway string `json:"apiGateway" structs:"apiGateway"` //API网关
}
