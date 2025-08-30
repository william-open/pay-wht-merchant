package biz

import (
	"github.com/gin-gonic/gin"
	"log"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/service/biz"
	"mwhtpay/config"
	"mwhtpay/core"
	"mwhtpay/core/response"
	"mwhtpay/middleware"
	"time"
)

var SafetyGroup = core.Group("/biz/safety", newSafetyHandler, regSafety, middleware.TokenAuth())

func newSafetyHandler(srv biz.IBizSafetyService) *safetyHandler {
	return &safetyHandler{srv: srv}
}

func regSafety(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *safetyHandler) {
		rg.POST("/gen/google_code", handle.genGoogleCode)
		rg.POST("/save/google_code", handle.saveGoogleCode)
		rg.POST("/save/pay_password", handle.savePayPassword)
	})
}

type safetyHandler struct {
	srv biz.IBizSafetyService
}

// genGoogleCode 生成google密钥
func (ch safetyHandler) genGoogleCode(c *gin.Context) {
	username := config.AdminConfig.GetUsername(c)
	if username == "" {
		response.FailWithMsg(c, response.SystemError, "系统错误")
		c.Abort()
		return
	}
	res, err := ch.srv.Gen(username)
	response.CheckAndRespWithData(c, res, err)
}

// saveGoogleCode 保存google密钥
func (ch safetyHandler) saveGoogleCode(c *gin.Context) {
	// 记录请求开始
	startTime := time.Now()
	adminId := uint(0)

	defer func() {
		// 记录请求完成时间和结果
		duration := time.Since(startTime)
		if r := recover(); r != nil {
			log.Printf("saveGoogleCode PANIC: 管理员=%v, 错误=%v, 耗时=%v",
				adminId, r, duration)
			response.FailWithMsg(c, response.SystemError, "系统内部错误")
		}
	}()

	// 1. 解析请求参数
	var saveReq req.BizSaveGoogleCodeReq
	if err := c.ShouldBindJSON(&saveReq); err != nil {
		log.Printf("参数解析失败: %v", err)
		response.FailWithMsg(c, response.SystemError, "参数错误")
		return
	}

	// 2. 获取管理员ID
	adminId = config.AdminConfig.GetAdminId(c)
	if adminId == 0 {
		log.Printf("未获取到管理员ID")
		response.FailWithMsg(c, response.NoPermission, "用户未认证")
		return
	}

	log.Printf("管理员 %v 请求绑定谷歌验证码", adminId)

	// 3. 调用服务层
	err := ch.srv.Save(saveReq, adminId)
	if err != nil {
		log.Printf("服务层处理失败: 管理员=%v, 错误=%v", adminId, err)
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}

	// 4. 成功响应
	log.Printf("处理成功: 管理员=%v, 耗时=%v", adminId, time.Since(startTime))
	response.OkWithMsg(c, "谷歌验证码绑定成功")
}

// savePayPassword 保存支付密码信息
func (ch safetyHandler) savePayPassword(c *gin.Context) {
	// 记录请求开始
	startTime := time.Now()
	adminId := uint(0)

	defer func() {
		// 记录请求完成时间和结果
		duration := time.Since(startTime)
		if r := recover(); r != nil {
			log.Printf("saveGoogleCode PANIC: 管理员=%v, 错误=%v, 耗时=%v",
				adminId, r, duration)
			response.FailWithMsg(c, response.SystemError, "系统内部错误")
		}
	}()

	// 1. 解析请求参数
	var saveReq req.BizSavePayPasswordReq
	if err := c.ShouldBindJSON(&saveReq); err != nil {
		log.Printf("参数解析失败: %v", err)
		response.FailWithMsg(c, response.SystemError, "参数错误")
		return
	}

	// 2. 获取管理员ID
	adminId = config.AdminConfig.GetAdminId(c)
	if adminId == 0 {
		log.Printf("未获取到管理员ID")
		response.FailWithMsg(c, response.NoPermission, "用户未认证")
		return
	}

	log.Printf("管理员 %v 请求绑定谷歌验证码", adminId)

	// 3. 调用服务层
	err := ch.srv.UpdatePayPassword(saveReq, adminId)
	if err != nil {
		log.Printf("服务层处理失败: 管理员=%v, 错误=%v", adminId, err)
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}

	// 4. 成功响应
	log.Printf("处理成功: 管理员=%v, 耗时=%v", adminId, time.Since(startTime))
	response.OkWithMsg(c, "操作成功")
}
