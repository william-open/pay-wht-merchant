package biz

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/schemas/resp"
	"mwhtpay/core"
	"mwhtpay/core/request"
	"mwhtpay/core/response"
	"mwhtpay/model/biz"
	"mwhtpay/model/system"
	"mwhtpay/util"
	"runtime/debug"
	"strings"
)

type IBizDockingService interface {
	List(page request.PageReq, listReq req.BizIpWhiteListReq) (res response.PageResp, e error)
	ConfigInfo(mId uint) (res resp.BizDockingConfigInfoResp, e error)
	ViewKey(req req.BizDockingViewKeyReq, adminId uint) (resp.BizDockingViewKeyResp, error)
}

// NewBizDockingService 初始化
func NewBizDockingService() IBizDockingService {
	// 通过DI获取订单数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizDockingService{db: mainDB}
}

// bizDockingService 字典数据服务实现类
type bizDockingService struct {
	db *gorm.DB
}

// List IP白名单数据所有
func (cSrv bizDockingService) List(page request.PageReq, listReq req.BizIpWhiteListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	dtResp := make([]resp.MerchantWhitelistResponse, 0)

	var queryDb = cSrv.db.Table("w_merchant_whitelist AS a").
		Joins("JOIN w_merchant AS b ON  a.m_id = b.m_id").
		Select("a.id,a.ip_address,a.can_admin,a.can_payout,a.can_receive,a.create_time")
	queryDb = queryDb.Where("a.m_id = ?", listReq.MId)
	var err = queryDb.
		Limit(limit).
		Offset(offset).
		Find(&dtResp).Error
	if err != nil {
		err := response.CheckErrDBNotRecord(err, "查询商户IP白名单失败！")
		if err != nil {
			return response.PageResp{}, err
		}
		return
	}
	// 总数
	var total int64
	queryDb.Count(&total)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    total,
		Lists:    dtResp,
	}, nil
}

// ConfigInfo 对接信息数据详情
func (cSrv bizDockingService) ConfigInfo(mId uint) (res resp.BizDockingConfigInfoResp, e error) {
	res = resp.BizDockingConfigInfoResp{}
	var merchant biz.Merchant
	if mId > 0 {
		e = cSrv.db.Where("m_id = ?", mId).Limit(1).First(&merchant).Error
		if e = response.CheckErr(e, "Console Get Merchant err"); e != nil {
			return res, e
		}
		res.AppId = merchant.AppId
	}
	data, err := util.ConfigUtil.Get(cSrv.db, "website")
	if e = response.CheckErr(err, "Detail Get err"); e != nil {
		return res, e
	}
	res.ApiDoc = data["apiDoc"]
	res.ApiGateway = data["apiGateway"]
	return res, nil
}

// ViewKey 对接信息数据详情
func (cSrv bizDockingService) ViewKey(req req.BizDockingViewKeyReq, adminId uint) (res resp.BizDockingViewKeyResp, e error) {
	res = resp.BizDockingViewKeyResp{} // 显式初始化

	// 捕获 panic，防止 500
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ViewKey] panic recovered: %v\n", r)
			log.Printf("[ViewKey] stack trace:\n%s", debug.Stack())
			e = errors.New("系统异常，请联系管理员")
		}
	}()

	if cSrv.db == nil {
		log.Println("数据库连接未初始化")
		return res, errors.New("系统异常，请联系管理员")
	}

	var admin system.SystemAuthAdmin
	if err := cSrv.db.Where("id = ?", adminId).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.New("管理员不存在")
		}
		log.Printf("查询失败: %v", err)
		return res, errors.New("查询管理员信息失败")
	}

	if admin.GoogleSecret != "" {
		if strings.TrimSpace(req.GoogleCode) == "" {
			return res, errors.New("请输入谷歌验证码")
		}
		valid, err := util.VerifyUtil.ValidateTOTPWithRetry(req.GoogleCode, admin.GoogleSecret)
		if err != nil {
			log.Printf("TOTP验证错误: %v", err)
			return res, errors.New("验证码验证失败")
		}
		if !valid {
			return res, errors.New("谷歌验证码错误")
		}
	}

	if req.PayPassword == "" {
		return res, errors.New("请输入支付密码")
	}
	if admin.PaySalt == "" {
		log.Println("管理员未设置支付盐值")
		return res, errors.New("管理员支付信息异常")
	}

	md5PayPwd := util.ToolsUtil.MakeMd5(req.PayPassword + admin.PaySalt)
	if admin.PayPassword != md5PayPwd {
		return res, errors.New("支付密码不对")
	}

	res.ApiKey = admin.ApiKey
	log.Printf("APIKEY: %s", res.ApiKey)
	return res, nil
}
