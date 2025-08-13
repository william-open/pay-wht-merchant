package biz

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/biz"
)

type IBizAddressService interface {
	List(page request.PageReq, listReq req.BizAddressListReq) (res response.PageResp, e error)
	Balance(page request.PageReq, listReq req.BizAddressBalanceReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.BizAddressResp, e error)
}

// NewBizAddressService 初始化
func NewBizAddressService(db *gorm.DB) IBizAddressService {
	return &bizAddressService{db: db}
}

// bizAddressService 字典数据服务实现类
type bizAddressService struct {
	db *gorm.DB
}

// List 商户钱包地址数据所有
func (cSrv bizAddressService) List(page request.PageReq, listReq req.BizAddressListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	var targets []biz.Address

	var queryDb = cSrv.db.Where("m_id = ?", listReq.MId)
	if listReq.Keyword != "" {
		queryDb = queryDb.Where("address like ?", "%"+listReq.Keyword+"%")
	}
	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&targets).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询交易下信息失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)

	var dtResp []resp.BizAddressResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Balance 商户钱包余额数据所有
func (cSrv bizAddressService) Balance(page request.PageReq, listReq req.BizAddressBalanceReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	var targets []biz.AddressAmount

	var queryDb = cSrv.db.Where("m_id = ?", listReq.MId)
	if listReq.Address != "" {
		queryDb = queryDb.Where("address = ?", listReq.Address)
	}
	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&targets).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询交易下信息失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)

	var dtResp []resp.BizAddressBalanceResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Detail 商户钱包地址数据详情
func (cSrv bizAddressService) Detail(id uint) (res resp.BizAddressResp, e error) {
	var dd biz.Address
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
