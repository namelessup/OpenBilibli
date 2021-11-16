package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/melloi/model"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func queryLabels(c *bm.Context) {
	c.JSON(srv.QueryLabel(c))
}

func addLabel(c *bm.Context) {
	label := model.Label{}
	if err := c.BindWith(&label, binding.JSON); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, srv.AddLabel(&label))
}

func delLabel(c *bm.Context) {
	v := new(struct {
		ID int64 `form:"id"`
	})
	if err := c.Bind(v); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.DeleteLabel(v.ID))
}

func addLabelRelation(c *bm.Context) {
	lr := model.LabelRelation{}
	if err := c.BindWith(&lr, binding.JSON); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, srv.AddLabelRelation(&lr))
}

func delLabelRelation(c *bm.Context) {
	v := new(struct {
		ID int64 `form:"id"`
	})
	if err := c.Bind(v); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DeleteLabelRelation(v.ID))
}
