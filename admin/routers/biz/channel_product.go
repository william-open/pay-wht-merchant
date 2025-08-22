package biz

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/biz"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var ChannelProductGroup = core.Group("/biz/pay_channel", newChannelProductHandler, regChannelProduct, middleware.TokenAuth())

func newChannelProductHandler(srv biz.IBizChannelProductService) *channelProductHandler {
	return &channelProductHandler{srv: srv}
}

func regChannelProduct(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *channelProductHandler) {
		rg.GET("/product", handle.list)
		rg.GET("/detail", handle.detail)
	})
}

type channelProductHandler struct {
	srv biz.IBizChannelProductService
}

// list 通道产品数据列表
func (ch channelProductHandler) list(c *gin.Context) {
	//var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	//var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizChannelProductListReq
	//listReq.MId = mId
	listReq.MId = 18
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ch.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 通道产品数据详情
func (ch channelProductHandler) detail(c *gin.Context) {
	var detailReq req.BizChannelProductDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ch.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}
