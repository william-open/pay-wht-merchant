package biz

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/service/biz"
	"mwhtpay/config"
	"mwhtpay/core"
	"mwhtpay/core/request"
	"mwhtpay/core/response"
	"mwhtpay/middleware"
	"mwhtpay/util"
)

var CollectionAddressGroup = core.Group("/biz/collection_address", newCollectionAddressHandler, regCollectionAddress, middleware.TokenAuth())

func newCollectionAddressHandler(srv biz.IBizCollectionAddressService) *collectionAddressHandler {
	return &collectionAddressHandler{srv: srv}
}

func regCollectionAddress(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *collectionAddressHandler) {
		rg.GET("/list", handle.list)
		rg.GET("/detail", handle.detail)
		rg.POST("/add", handle.add)
		rg.POST("/del", handle.del)
		rg.POST("/status", handle.status)
		rg.GET("/balance", handle.balance)
	})
}

type collectionAddressHandler struct {
	srv biz.IBizCollectionAddressService
}

// list 商户归集钱包地址数据列表
func (ch collectionAddressHandler) list(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizCollectionAddressListReq
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

// detail 商户归集钱包地址数据详情
func (ch collectionAddressHandler) detail(c *gin.Context) {
	var detailReq req.BizCurrentDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ch.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 商户归集钱包地址新增
func (ch collectionAddressHandler) add(c *gin.Context) {

	var addReq req.BizCollectionAddressAddReq
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	addReq.MId = mId
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, ch.srv.Add(addReq))
}

// del 商户归集钱包地址删除
func (ch collectionAddressHandler) del(c *gin.Context) {
	var delReq req.BizCollectionAddressDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, ch.srv.Del(delReq))
}

// status 商户归集钱包状态切换
func (ch collectionAddressHandler) status(c *gin.Context) {
	var disableReq req.BizCollectionAddressStatusReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	response.CheckAndResp(c, ch.srv.Status(c, disableReq.ID))
}

// balance 商户钱包余额数据列表
func (ch collectionAddressHandler) balance(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizCollectionAddressBalanceReq
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
