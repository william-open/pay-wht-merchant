package biz

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/biz"
)

type IBizCollectionAddressService interface {
	List(page request.PageReq, listReq req.BizCollectionAddressListReq) (res response.PageResp, e error)
	Balance(page request.PageReq, listReq req.BizCollectionAddressBalanceReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.BizCollectionAddressResp, e error)
	Add(addReq req.BizCollectionAddressAddReq) (e error)
	Del(delReq req.BizCollectionAddressDelReq) (e error)
	Status(c *gin.Context, id uint) (e error)
}

// NewBizCollectionAddressService 初始化
func NewBizCollectionAddressService() IBizCollectionAddressService {
	// 通过DI获取主数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizCollectionAddressService{db: mainDB}
}

// bizCollectionAddressService 字典数据服务实现类
type bizCollectionAddressService struct {
	db *gorm.DB
}

// Status 状态切换
func (cSrv bizCollectionAddressService) Status(c *gin.Context, id uint) (e error) {
	var statusRecord biz.CollectionAddress
	err := cSrv.db.Where("id = ?", id).Limit(1).Find(&statusRecord).Error
	if e = response.CheckErr(err, "Status Find err"); e != nil {
		return
	}
	if statusRecord.Id == 0 {
		return response.AssertArgumentError.Make("记录已不存在!")
	}
	var status uint8

	err = cSrv.db.Model(&statusRecord).Updates(biz.CollectionAddress{Status: status}).Error
	e = response.CheckErr(err, "Status Updates err")
	return
}

// Balance 商户归集钱包数据列表
func (cSrv bizCollectionAddressService) Balance(page request.PageReq, listReq req.BizCollectionAddressBalanceReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	var targets []biz.CollectionAddressAmount

	var queryDb = cSrv.db.Where("m_id = ?", listReq.MId)

	if listReq.Address != "" {
		queryDb = queryDb.Where("address = ? ", listReq.Address)
	}

	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&targets).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询归集账户信息失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)

	var dtResp []resp.BizAddressBalanceAmountResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// List 商户货币数据列表
func (cSrv bizCollectionAddressService) List(page request.PageReq, listReq req.BizCollectionAddressListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	var targets []biz.CollectionAddress

	var queryDb = cSrv.db.Where("m_id = ?", listReq.MId)

	if listReq.Address != "" {
		queryDb = queryDb.Where("address like ?", "%"+listReq.Address+"%")
	}

	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&targets).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询归集账户信息失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)

	var dtResp []resp.BizCollectionAddressResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Detail 归集账户数据详情
func (cSrv bizCollectionAddressService) Detail(id uint) (res resp.BizCollectionAddressResp, e error) {
	var dd biz.CollectionAddress
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

// Add 商户归集账户数据新增
func (cSrv bizCollectionAddressService) Add(addReq req.BizCollectionAddressAddReq) (e error) {

	if r := cSrv.db.Where("m_id = ? AND address = ? AND chain_symbol = ?", addReq.MId, addReq.Address, addReq.ChainSymbol).Limit(1).First(&biz.CollectionAddress{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("数据已存在！")
	}
	var dd biz.CollectionAddress
	dd.MId = addReq.MId
	response.Copy(&dd, addReq)
	err := cSrv.db.Create(&dd).Error

	e = response.CheckErr(err, "Add Create err")
	return

}

// Del 商户归集账户数据删除
func (cSrv bizCollectionAddressService) Del(delReq req.BizCollectionAddressDelReq) (e error) {
	err := cSrv.db.Delete(&biz.CollectionAddress{}, delReq.Ids).Error
	return response.CheckErr(err, "Del Data err")
}
