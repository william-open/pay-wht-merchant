package req

// BizCurrencyListReq 货币数据列表参数
type BizCurrencyListReq struct {
	CurrencyType string `form:"currencyType" binding:"max=200"` // 货币类型
	Currency     string `form:"currency" binding:"max=100"`     // 货币名称
}

// BizMCurrencyDelReq 商户货币数据删除参数
type BizMCurrencyDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}

// BizMCurrencyListReq 商户货币数据列表参数
type BizMCurrencyListReq struct {
	CurrencyType string `form:"currencyType" binding:"max=200"` // 货币类型
	Currency     string `form:"currency" binding:"max=100"`     // 货币名称
	MId          uint   `form:"mId" binding:"max=100"`          // 商户ID
}

// BizMCurrencyAddReq 商户货币数据新增参数
type BizMCurrencyAddReq struct {
	Mid        uint   `form:"mId" binding:"required,gt=0"` // 商户ID
	CurrencyId []uint `form:"currencyId" binding:"gt=0"`   // 货币ID
}

// BizCurrentDetailReq 货币数据详情参数
type BizCurrentDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizTransactionDetailReq 交易数据详情参数
type BizTransactionDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizTransactionListReq 交易数据列表参数
type BizTransactionListReq struct {
	Keyword string `form:"keyword" binding:"max=200"`   // 关键字
	MId     uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}

// BizAddressDetailReq 钱包数据详情参数
type BizAddressDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizAddressListReq 交易数据列表参数
type BizAddressListReq struct {
	Keyword string `form:"keyword" binding:"max=200"`   // 关键字
	MId     uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}

// BizAddressBalanceReq 钱包余额数据列表参数
type BizAddressBalanceReq struct {
	Address string `form:"address" binding:"max=200"`   // 钱包地址
	MId     uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}

// BizCollectionDetailReq 归集数据详情参数
type BizCollectionDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizCollectionListReq 归集数据列表参数
type BizCollectionListReq struct {
	Keyword string `form:"keyword" binding:"max=200"`   // 关键字
	MId     uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}

// BizCollectionAddressListReq 归集账户数据列表参数
type BizCollectionAddressListReq struct {
	Address string `form:"address"`                // 钱包地址
	MId     uint   `form:"mId" binding:"required"` // 商户ID
}

// BizCollectionAddressDelReq 归集账户数据删除参数
type BizCollectionAddressDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}

// BizCollectionAddressAddReq 归集账户数据新增参数
type BizCollectionAddressAddReq struct {
	MId         uint   `form:"mId" binding:"required,gt=0"`    // 商户ID
	ChainSymbol string `form:"chainSymbol" binding:"required"` // 主链符号
	Address     string `form:"address" binding:"required"`     // 钱包地址
	MinAmount   string `form:"minAmount" binding:"required"`   // 最小金额
	MaxAmount   string `form:"maxAmount" binding:"required"`   // 最大金额
}

// BizCollectionAddressDetailReq 货币数据详情参数
type BizCollectionAddressDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizCollectionAddressStatusReq 状态切换参数
type BizCollectionAddressStatusReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizCollectionAddressBalanceReq 钱包余额数据列表参数
type BizCollectionAddressBalanceReq struct {
	Address string `form:"address" binding:"max=200"`   // 钱包地址
	MId     uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}

// BizChannelProductListReq 通道产品列表参数
type BizChannelProductListReq struct {
	Keyword string `form:"keyword" binding:"max=200"`   // 关键字
	MId     uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}

// BizChannelProductDetailReq 通道产品数据详情参数
type BizChannelProductDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// BizCollectOrderListReq 收款订单列表参数
type BizCollectOrderListReq struct {
	Keyword    string `query:"keyword" form:"keyword"`
	ChannelId  int64  `query:"channelId" form:"channelId"`
	UpChanelId int64  `query:"upChanelId" form:"upChanelId"`
	Status     string `query:"status" form:"status"`
	Currency   string `query:"currency" form:"currency"`
	YearMonth  string `query:"yearMonth" form:"yearMonth"`
	MId        uint   `form:"mId" binding:"required,gt=0"` // 商户ID
}
