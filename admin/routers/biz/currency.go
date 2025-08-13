package biz

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/biz"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var CurrencyGroup = core.Group("/biz/currency", newCurrencyHandler, regCurrency, middleware.TokenAuth())

func newCurrencyHandler(srv biz.IBizCurrencyService) *currencyHandler {
	return &currencyHandler{srv: srv}
}

func regCurrency(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *currencyHandler) {
		rg.GET("/all", handle.all)
		rg.GET("/m/list", handle.list)
		rg.GET("/detail", handle.detail)
		rg.POST("/m//add", handle.add)
		rg.POST("/m/del", handle.del)
	})
}

type currencyHandler struct {
	srv biz.IBizCurrencyService
}

// all 平台所有货币信息
func (ch currencyHandler) all(c *gin.Context) {
	var allReq req.BizCurrencyListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &allReq)) {
		return
	}
	res, err := ch.srv.All(allReq)
	response.CheckAndRespWithData(c, res, err)
}

// list 商户货币数据列表
func (ch currencyHandler) list(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizMCurrencyListReq
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

// detail 货币数据详情
func (ch currencyHandler) detail(c *gin.Context) {
	var detailReq req.BizCurrentDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ch.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 商户货币数据新增
func (ch currencyHandler) add(c *gin.Context) {

	var addReq req.BizMCurrencyAddReq
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	addReq.Mid = mId
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, ch.srv.Add(addReq))
}

// del 字典数据删除
func (ch currencyHandler) del(c *gin.Context) {
	var delReq req.BizMCurrencyDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, ch.srv.Del(delReq))
}
