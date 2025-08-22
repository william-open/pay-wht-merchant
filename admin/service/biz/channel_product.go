package biz

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/biz"
)

type IBizChannelProductService interface {
	List(page request.PageReq, listReq req.BizChannelProductListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.BizChannelProductResp, e error)
}

// NewBizChannelProductService 初始化
func NewBizChannelProductService() IBizChannelProductService {
	// 通过DI获取订单数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizChannelProductService{db: mainDB}
}

// bizChannelProductService 字典数据服务实现类
type bizChannelProductService struct {
	db *gorm.DB
}

// List 通道产品数据所有
func (cSrv bizChannelProductService) List(page request.PageReq, listReq req.BizChannelProductListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	//log.Printf("商户ID:%+v", listReq.MId)
	//var targets []biz.MerchantChannel
	dtResp := make([]resp.BizChannelProductResp, 0)

	var queryDb = cSrv.db.Table("w_merchant_channel AS a").Where("a.m_id = ?", listReq.MId)
	queryDb.
		Joins("JOIN w_pay_way AS b ON a.sys_channel_id = b.id").
		Joins("JOIN w_currency_code AS c ON a.currency = c.`code`").
		Joins("JOIN w_merchant AS d ON a.m_id = d.m_id").
		Select("a.id,a.default_rate,a.single_fee,a.order_range,b.coding as channel_code,b.title as channel_title,c.country")
	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&dtResp).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询商户通道产品息失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)

	//log.Printf("返回值: %+v", dtResp)
	//var dtResp []resp.BizChannelProductResp
	//response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Detail 通道产品数据详情
func (cSrv bizChannelProductService) Detail(id uint) (res resp.BizChannelProductResp, e error) {
	var dd biz.Transaction
	err := cSrv.db.Where("id = ?", id, 0).Limit(1).First(&dd).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, dd)
	return
}
