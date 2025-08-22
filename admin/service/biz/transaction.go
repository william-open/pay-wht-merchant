package biz

import (
	"gorm.io/gorm"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/schemas/resp"
	"mwhtpay/core"
	"mwhtpay/core/request"
	"mwhtpay/core/response"
	"mwhtpay/model/biz"
)

type IBizTransactionService interface {
	List(page request.PageReq, listReq req.BizTransactionListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.BizTransactionResp, e error)
}

// NewBizTransactionService 初始化
func NewBizTransactionService() IBizTransactionService {
	// 通过DI获取主数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizTransactionService{db: mainDB}
}

// bizTransactionService 字典数据服务实现类
type bizTransactionService struct {
	db *gorm.DB
}

// List 交易数据所有
func (cSrv bizTransactionService) List(page request.PageReq, listReq req.BizTransactionListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	var targets []biz.Transaction

	var queryDb = cSrv.db.Where("m_id = ?", listReq.MId)
	if listReq.Keyword != "" {
		queryDb = queryDb.Where("biz_no like ?", "%"+listReq.Keyword+"%").Or("tx_id like ?", "%"+listReq.Keyword+"%").Or("to_address like ?", "%"+listReq.Keyword+"%")
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

	var dtResp []resp.BizTransactionResp
	response.Copy(&dtResp, targets)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// Detail 交易数据详情
func (cSrv bizTransactionService) Detail(id uint) (res resp.BizTransactionResp, e error) {
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
