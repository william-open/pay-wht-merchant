package routers

import (
	"likeadmin/admin/routers/biz"
	"likeadmin/admin/routers/common"
	"likeadmin/admin/routers/monitor"
	"likeadmin/admin/routers/setting"
	"likeadmin/admin/routers/system"
	"likeadmin/core"
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
}
