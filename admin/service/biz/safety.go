package biz

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
	"log"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/schemas/resp"
	"mwhtpay/core"
	"mwhtpay/model/system"
	"mwhtpay/util"
	"strings"
	"time"
)

type IBizSafetyService interface {
	Gen(username string) (resp.BizGenGoogleCodeResp, error)
	Save(param req.BizSaveGoogleCodeReq, adminId uint) error
	UpdatePayPassword(param req.BizSavePayPasswordReq, adminId uint) error
}

// NewBizSafetyService 初始化
func NewBizSafetyService() IBizSafetyService {
	// 通过DI获取订单数据库连接
	mainDB, exists := core.GetDatabase(core.DBMain)
	if !exists {
		panic("main database not initialized")
	}
	return &bizSafetyService{db: mainDB}
}

// bizSafetyService 字典数据服务实现类
type bizSafetyService struct {
	db *gorm.DB
}

// Gen 生成谷歌验证码
func (cSrv *bizSafetyService) Gen(username string) (resp.BizGenGoogleCodeResp, error) {
	var res resp.BizGenGoogleCodeResp
	//log.Printf("登录用户: %v", username)
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "WhtMchSys",
		AccountName: username,
	})
	if err != nil {
		return res, errors.New("生成密钥失败")
	}

	// 生成二维码
	png, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		return res, errors.New("生成二维码失败")
	}
	// 返回二维码图片（Base64）
	encoded := base64.StdEncoding.EncodeToString(png)
	qrcodeUrl := "data:image/png;base64," + encoded
	res.Qrcode = qrcodeUrl
	res.Secret = key.Secret()
	return res, nil
}

// Save 保存谷歌验证码 - 安全版本
func (cSrv *bizSafetyService) Save(param req.BizSaveGoogleCodeReq, adminId uint) error {
	// 输入验证
	if cSrv == nil || cSrv.db == nil {
		log.Printf("严重错误: service 或 db 为空指针")
		return errors.New("系统内部错误")
	}

	if adminId == 0 {
		return errors.New("管理员ID无效")
	}

	if strings.TrimSpace(param.GoogleSecret) == "" {
		return errors.New("谷歌密钥不能为空")
	}

	log.Printf("开始处理管理员 %v 的谷歌验证码绑定", adminId)

	// 使用事务确保数据一致性
	return cSrv.db.Transaction(func(tx *gorm.DB) error {
		// 查询管理员
		var admin system.SystemAuthAdmin
		if err := tx.Where("id = ?", adminId).First(&admin).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("管理员不存在")
			}
			log.Printf("查询失败: %v", err)
			return errors.New("查询管理员信息失败")
		}

		log.Printf("管理员信息: ID=%v, Name=%v", admin.ID, admin.Nickname)

		// 验证现有谷歌验证码（如果已绑定）
		if admin.GoogleSecret != "" {
			if param.GoogleCode == "" {
				return errors.New("请输入旧的谷歌验证码")
			}

			// 添加TOTP验证的容错机制
			valid, err := util.VerifyUtil.ValidateTOTPWithRetry(param.GoogleCode, admin.GoogleSecret)
			if err != nil {
				log.Printf("TOTP验证错误: %v", err)
				return errors.New("验证码验证失败")
			}

			if !valid {
				return errors.New("谷歌验证码错误")
			}
		}

		// 更新数据
		updateData := map[string]interface{}{
			"google_secret":     param.GoogleSecret,
			"is_google_enabled": true,
			"update_time":       time.Now().Unix(),
		}

		if err := tx.Model(&system.SystemAuthAdmin{}).
			Where("id = ?", adminId).
			Updates(updateData).Error; err != nil {
			log.Printf("更新失败: %v", err)
			return errors.New("保存谷歌验证码失败")
		}

		log.Printf("谷歌验证码绑定成功: 管理员 %v", adminId)
		return nil
	})
}

// UpdatePayPassword 保存支付密码 - 安全版本
func (cSrv *bizSafetyService) UpdatePayPassword(param req.BizSavePayPasswordReq, adminId uint) error {
	// 输入验证
	if cSrv == nil || cSrv.db == nil {
		log.Printf("严重错误: service 或 db 为空指针")
		return errors.New("系统内部错误")
	}

	if adminId == 0 {
		return errors.New("管理员ID无效")
	}

	// 使用事务确保数据一致性
	return cSrv.db.Transaction(func(tx *gorm.DB) error {
		// 查询管理员
		var admin system.SystemAuthAdmin
		if err := tx.Where("id = ?", adminId).First(&admin).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("管理员不存在")
			}
			log.Printf("查询失败: %v", err)
			return errors.New("查询管理员信息失败")
		}

		log.Printf("开始处理管理员 %v 的谷歌验证码绑定", adminId)

		log.Printf("管理员信息: ID=%v, Name=%v", admin.ID, admin.Nickname)

		log.Printf("旧密码: %+v", strings.Trim(param.OldPayPassword, " "))
		log.Printf("支付盐值: %+v", admin.PaySalt)
		md5PayPwd := util.ToolsUtil.MakeMd5(strings.Trim(param.OldPayPassword, " ") + admin.PaySalt)
		log.Printf("加密: %+v", md5PayPwd)
		if admin.PayPassword != md5PayPwd {
			return errors.New("旧支付密码不对")
		}

		// 验证现有谷歌验证码（如果已绑定）
		if admin.IsGoogleEnabled > 0 {
			if strings.TrimSpace(param.GoogleCode) == "" {
				return errors.New("请输入谷歌验证码")
			}

			// 添加TOTP验证的容错机制
			valid, err := util.VerifyUtil.ValidateTOTPWithRetry(param.GoogleCode, admin.GoogleSecret)
			if err != nil {
				log.Printf("TOTP验证错误: %v", err)
				return errors.New("验证码验证失败")
			}

			if !valid {
				return errors.New("谷歌验证码错误")
			}
		}

		// 重新生成支付盐值
		salt := util.ToolsUtil.RandomString(5)
		newMd5PayPwd := util.ToolsUtil.MakeMd5(param.NewPayPassword + salt)
		// 更新数据
		updateData := map[string]interface{}{
			"pay_salt":     salt,
			"pay_password": newMd5PayPwd,
			"update_time":  time.Now().Unix(),
		}

		if err := tx.Model(&system.SystemAuthAdmin{}).
			Where("id = ?", adminId).
			Updates(updateData).Error; err != nil {
			log.Printf("更新失败: %v", err)
			return errors.New("保存支付密码失败")
		}

		log.Printf("支付密码更新成功: 管理员 %v", adminId)
		return nil
	})
}
