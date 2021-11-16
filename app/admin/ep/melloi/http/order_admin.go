package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/melloi/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

// get administrator for order by current username
func queryOrderAdmin(c *bm.Context) {
	userName, _ := c.Request.Cookie("username")
	c.JSON(srv.QueryOrderAdmin(userName.Value))
}

// add administrator for order
func addOrderAdmin(c *bm.Context) {
	admin := model.OrderAdmin{}
	if err := c.BindWith(&admin, binding.Form); err != nil {
		return
	}
	c.JSON(nil, srv.AddOrderAdmin(&admin))
}
