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

type IBizCollectionService interface {
	List(page request.PageReq, listReq req.BizCollectionListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.BizCollectionResp, e error)
}

// NewBizCollectionService 初始化
func NewBizCollectionService() IBizCollectionService {
	// 通过DI获取主数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizCollectionService{db: mainDB}
}

// bizCollectionService 归集数据服务实现类
type bizCollectionService struct {
	db *gorm.DB
}

// List 商户钱包地址数据所有
func (cSrv bizCollectionService) List(page request.PageReq, listReq req.BizCollectionListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	var targets []biz.Collection

	var queryDb = cSrv.db.Where("m_id = ?", listReq.MId)
	if listReq.Keyword != "" {
		queryDb = queryDb.Where("from_address like ?", "%"+listReq.Keyword+"%").Or("to_address like ?", "%"+listReq.Keyword+"%")
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

	var dtResp []resp.BizCollectionResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Detail 商户钱包地址数据详情
func (cSrv bizCollectionService) Detail(id uint) (res resp.BizCollectionResp, e error) {
	var dd biz.Collection
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
