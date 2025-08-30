package routers

import (
	"mwhtpay/admin/routers/biz"
	"mwhtpay/admin/routers/common"
	"mwhtpay/admin/routers/monitor"
	"mwhtpay/admin/routers/setting"
	"mwhtpay/admin/routers/system"
	"mwhtpay/core"
)

var InitRouters = []*core.GroupBase{
	// common
	common.AlbumGroup,
	common.IndexGroup,
	common.UploadGroup,
	// monitor
	monitor.MonitorGroup,
	// setting
	setting.CopyrightGroup,
	setting.DictDataGroup,
	setting.DictTypeGroup,
	setting.ProtocolGroup,
	setting.StorageGroup,
	setting.WebsiteGroup,
	// system
	system.AdminGroup,
	system.DeptGroup,
	system.LogGroup,
	system.LoginGroup,
	system.MenuGroup,
	system.PostGroup,
	system.RoleGroup,
	// biz
	biz.CurrencyGroup,
	biz.TransactionGroup,
	biz.AddressGroup,
	biz.CollectionGroup,
	biz.CollectionAddressGroup,
	biz.ChannelProductGroup,
	biz.OrderGroup,
	biz.DockingGroup,
	biz.SafetyGroup,
}
