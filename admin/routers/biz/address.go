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

var AddressGroup = core.Group("/biz/address", newAddressHandler, regAddress, middleware.TokenAuth())

func newAddressHandler(srv biz.IBizAddressService) *addressHandler {
	return &addressHandler{srv: srv}
}

func regAddress(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *addressHandler) {
		rg.GET("/list", handle.list)
		rg.GET("/detail", handle.detail)
		rg.GET("/balance", handle.balance)
	})
}

type addressHandler struct {
	srv biz.IBizAddressService
}

// list 商户钱包地址数据列表
func (ch addressHandler) list(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizAddressListReq
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

// balance 商户钱包余额数据列表
func (ch addressHandler) balance(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizAddressBalanceReq
	listReq.MId = mId
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ch.srv.Balance(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 商户钱包地址数据详情
func (ch addressHandler) detail(c *gin.Context) {
	var detailReq req.BizCurrentDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ch.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}
