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

var CollectionGroup = core.Group("/biz/collection", newCollectionHandler, regCollection, middleware.TokenAuth())

func newCollectionHandler(srv biz.IBizCollectionService) *collectionHandler {
	return &collectionHandler{srv: srv}
}

func regCollection(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *collectionHandler) {
		rg.GET("/list", handle.list)
		rg.GET("/detail", handle.detail)
	})
}

type collectionHandler struct {
	srv biz.IBizCollectionService
}

// list 商户归集数据列表
func (ch collectionHandler) list(c *gin.Context) {
	var mIdStr, _ = c.Get(config.AdminConfig.ReqAdminMIdKey)
	var mId, _ = util.ToolsUtil.StringToUint(fmt.Sprintf("%v", mIdStr))
	var page request.PageReq
	var listReq req.BizCollectionListReq
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

// detail 商户归集地址数据详情
func (ch collectionHandler) detail(c *gin.Context) {
	var detailReq req.BizCurrentDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ch.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}
