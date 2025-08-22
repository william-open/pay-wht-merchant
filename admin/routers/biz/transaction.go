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

var TransactionGroup = core.Group("/biz/transaction", newTransactionHandler, regTransaction, middleware.TokenAuth())

func newTransactionHandler(srv biz.IBizTransactionService) *transactionHandler {
	return &transactionHandler{srv: srv}
}

func regTransaction(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *transactionHandler) {
		rg.GET("/list", handle.list)
		rg.GET("/detail", handle.detail)
	})
}

type transactionHandler struct {
	srv biz.IBizTransactionService
}

// list 商户货币数据列表
func (ch transactionHandler) list(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizTransactionListReq
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
func (ch transactionHandler) detail(c *gin.Context) {
	var detailReq req.BizCurrentDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ch.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}
