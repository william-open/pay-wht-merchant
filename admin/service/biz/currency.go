package biz

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/biz"
)

type IBizCurrencyService interface {
	All(allReq req.BizCurrencyListReq) (res []resp.BizCurrencyResp, e error)
	List(page request.PageReq, listReq req.BizMCurrencyListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.BizCurrencyResp, e error)
	Add(addReq req.BizMCurrencyAddReq) (e error)
	Del(delReq req.BizMCurrencyDelReq) (e error)
}

// NewBizCurrencyService 初始化
func NewBizCurrencyService() IBizCurrencyService {
	// 通过DI获取主数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizCurrencyService{db: mainDB}
}

// bizCurrencyService 字典数据服务实现类
type bizCurrencyService struct {
	db *gorm.DB
}

// All 货币数据所有
func (cSrv bizCurrencyService) All(allReq req.BizCurrencyListReq) (res []resp.BizCurrencyResp, e error) {

	ddModel := cSrv.db.Where("status = ?", 0)
	if allReq.Currency != "" {
		ddModel = ddModel.Where("currency like ?", "%"+allReq.Currency+"%")
	}
	if allReq.CurrencyType != "" {
		ddModel = ddModel.Where("currency_type = ?", "%"+allReq.CurrencyType+"%")
	}

	var currencyList []biz.Currency
	err := ddModel.Order("currency_id desc").Find(&currencyList).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []resp.BizCurrencyResp{}
	response.Copy(&res, currencyList)
	return
}

// List 商户货币数据列表
func (cSrv bizCurrencyService) List(page request.PageReq, listReq req.BizMCurrencyListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	//表前缀
	var tablePrefix = config.Config.DbTablePrefix
	var targets []biz.Currency

	var queryDb = cSrv.db.Table(tablePrefix+"currency AS a").
		Joins("JOIN "+tablePrefix+"merchant_currency AS b  ON a.currency_id = b.currency_id").Where("b.m_id = ?", listReq.MId)

	if listReq.Currency != "" {
		queryDb = queryDb.Where("a.currency like ?", "%"+listReq.Currency+"%")
	}
	if listReq.CurrencyType != "" {
		queryDb = queryDb.Where("a.currency_type like ?", "%"+listReq.CurrencyType+"%")
	}

	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&targets).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询商户货币信息失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)

	var dtResp []resp.BizCurrencyResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Detail 货币数据详情
func (cSrv bizCurrencyService) Detail(id uint) (res resp.BizCurrencyResp, e error) {
	var dd biz.Currency
	err := cSrv.db.Where("currency_id = ?", id, 0).Limit(1).First(&dd).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, dd)
	return
}

// Add 商户货币数据新增
func (cSrv bizCurrencyService) Add(addReq req.BizMCurrencyAddReq) (e error) {

	for _, id := range addReq.CurrencyId {

		if r := cSrv.db.Where("currency_id = ? AND m_id = ?", id, addReq.Mid).Limit(1).First(&biz.MerchantCurrency{}); r.RowsAffected > 0 {
			return response.AssertArgumentError.Make("数据已存在！")
		}
		var dd biz.MerchantCurrency
		dd.CurrencyId = id
		dd.MId = addReq.Mid
		//response.Copy(&dd, addReq)
		err := cSrv.db.Create(&dd).Error

		if err != nil {
			e = response.CheckErr(err, "Add Create err")
			return
		}
	}

	return nil

}

// Del 商户货币数据删除
func (cSrv bizCurrencyService) Del(delReq req.BizMCurrencyDelReq) (e error) {
	err := cSrv.db.Delete(&biz.Currency{}, delReq.Ids).Error
	return response.CheckErr(err, "Del Data err")
}
