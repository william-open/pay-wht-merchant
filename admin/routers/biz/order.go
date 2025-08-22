package biz

import (
	"github.com/gin-gonic/gin"
	"mwhtpay/admin/schemas/req"
	"mwhtpay/admin/service/biz"
	"mwhtpay/core"
	"mwhtpay/core/request"
	"mwhtpay/core/response"
	"mwhtpay/middleware"
	"mwhtpay/util"
)

var OrderGroup = core.Group("/biz/order", newOrderHandler, regOrder, middleware.TokenAuth())

func newOrderHandler(srv biz.IBizOrderService) *OrderHandler {
	return &OrderHandler{srv: srv}
}

func regOrder(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *OrderHandler) {
		rg.GET("/collect_list", handle.collectList)
	})
}

type OrderHandler struct {
	srv biz.IBizOrderService
}

// list 收款订单数据列表
func (ch OrderHandler) collectList(c *gin.Context) {
	//var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	//var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizCollectOrderListReq
	//listReq.MId = mId
	listReq.MId = 18
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ch.srv.CollectList(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}
