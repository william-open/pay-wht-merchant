package biz

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/service/biz"
	"mwhtpay/config"
	"mwhtpay/core"
	"mwhtpay/core/request"
	"mwhtpay/core/response"
	"mwhtpay/middleware"
	"mwhtpay/util"
	"time"
)

var DockingGroup = core.Group("/biz/docking", newDockingHandler, regDocking, middleware.TokenAuth())

func newDockingHandler(srv biz.IBizDockingService) *dockingHandler {
	return &dockingHandler{srv: srv}
}

func regDocking(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *dockingHandler) {
		rg.GET("/ip_whitelist", handle.list)
		rg.GET("/config_info", handle.configInfo)
		rg.POST("/view_key", handle.viewKey)
	})
}

type dockingHandler struct {
	srv biz.IBizDockingService
}

// list IP白名单数据列表
func (ch dockingHandler) list(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizIpWhiteListReq
	listReq.MId = mId
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ch.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// configInfo 对接信息数据详情
func (ch dockingHandler) configInfo(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)

	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	if mId <= 0 {
		response.FailWithMsg(c, response.SystemError, "管理员身份异常")
		return
	}
	res, err := ch.srv.ConfigInfo(mId)
	response.CheckAndRespWithData(c, res, err)
}

// viewKey 查看密钥
func (ch dockingHandler) viewKey(c *gin.Context) {
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
	var detailReq req.BizDockingViewKeyReq

	// 获取管理员 ID
	adminIdRaw, exists := c.Get(config.AdminConfig.ReqAdminIdKey)
	if !exists {
		log.Println("[viewKey] 未获取到管理员ID")
		response.FailWithMsg(c, response.SystemError, "管理员身份异常")
		return
	}

	adminId, err := util.ToolsUtil.StringToUint(fmt.Sprintf("%v", adminIdRaw))
	if err != nil || adminId == 0 {
		log.Printf("[viewKey] 管理员ID转换失败: %v", adminIdRaw)
		response.FailWithMsg(c, response.SystemError, "管理员身份异常")
		return
	}

	// 参数验证
	if err := util.VerifyUtil.VerifyBody(c, &detailReq); err != nil {
		log.Printf("[viewKey] 参数验证失败: %v", err)
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}

	// 调用服务层
	res, err := ch.srv.ViewKey(detailReq, adminId)
	if err != nil {
		log.Printf("[viewKey] ViewKey服务调用失败: %v", err)
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}
	response.CheckAndRespWithData(c, res, err)
}
