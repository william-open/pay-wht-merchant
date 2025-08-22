package biz

import (
	"fmt"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"sort"
	"time"
)

type IBizOrderService interface {
	CollectList(page request.PageReq, listReq req.BizCollectOrderListReq) (res response.PageResp, e error)
}

// NewBizOrderService 初始化
func NewBizOrderService() IBizOrderService {
	// 通过DI获取订单数据库连接
	orderDB, exists := core.GetDatabase(core.DBOrder)
	if !exists {
		panic("order database not initialized")
	}
	return &bizOrderService{
		db: orderDB,
	}
}

// bizOrderService 字典数据服务实现类
type bizOrderService struct {
	db *gorm.DB // 默认数据库
}

func generateOrderTableNames(yearMonth string) []string {
	return []string{
		fmt.Sprintf("p_order_%s_p0", yearMonth),
		fmt.Sprintf("p_order_%s_p1", yearMonth),
		fmt.Sprintf("p_order_%s_p2", yearMonth),
		fmt.Sprintf("p_order_%s_p3", yearMonth),
	}
}

// ---------------------- 查询订单列表（聚合分页） ----------------------
func (cSrv bizOrderService) CollectList(page request.PageReq, listReq req.BizCollectOrderListReq) (res response.PageResp, e error) {
	var allOrders []resp.OrderReceiveListResponse
	var count int64

	yearMonth := listReq.YearMonth
	if yearMonth == "" {
		yearMonth = time.Now().Format("200601")
	}

	// 遍历所有分片表
	for _, table := range generateOrderTableNames(yearMonth) {
		var tableOrders []resp.OrderReceiveListResponse
		query := cSrv.db.Table(table)

		if listReq.Keyword != "" {
			query = query.Where("order_id LIKE ? or m_order_id LIKE ?", "%"+listReq.Keyword+"%", "%"+listReq.Keyword+"%")
		}
		if listReq.Status != "" {
			query = query.Where("status = ?", listReq.Status)
		}

		var tableCount int64
		if err := query.Count(&tableCount).Error; err != nil {
			// 如果表不存在或查询出错，跳过继续处理其他表
			continue
		}
		count += tableCount

		if err := query.Find(&tableOrders).Error; err != nil {
			// 如果查询出错，跳过继续处理其他表
			continue
		}
		allOrders = append(allOrders, tableOrders...)
	}

	// 内存排序（按创建时间倒序）
	sort.Slice(allOrders, func(i, j int) bool {
		return allOrders[i].CreateTime.After(allOrders[j].CreateTime)
	})

	// 内存分页
	total := len(allOrders)
	start := (page.PageNo - 1) * page.PageSize
	if start >= total {
		return response.PageResp{
			PageNo:   page.PageNo,
			PageSize: page.PageSize,
			Count:    int64(total),
			Lists:    []resp.OrderReceiveListResponse{},
		}, nil
	}

	end := start + page.PageSize
	if end > total {
		end = total
	}

	pagedOrders := allOrders[start:end]

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    int64(total),
		Lists:    pagedOrders,
	}, nil
}
