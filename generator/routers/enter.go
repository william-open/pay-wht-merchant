package routers

import (
	"mwhtpay/core"
	"mwhtpay/generator/routers/gen"
)

var InitRouters = []*core.GroupBase{
	// gen
	gen.GenGroup,
}
