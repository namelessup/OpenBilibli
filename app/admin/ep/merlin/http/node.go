package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/merlin/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func updateNodes(c *bm.Context) {
	var (
		umnr = &model.UpdateMachineNodeRequest{}
		err  error
	)
	if err = c.BindWith(umnr, binding.JSON); err != nil {
		return
	}
	if err = umnr.VerifyNodes(); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, svc.UpdateMachineNode(c, umnr))
}

func queryNodes(c *bm.Context) {
	v := new(struct {
		MachineID int64 `form:"machine_id"`
	})
	if err := c.Bind(v); err != nil {
		return
	}
	c.JSON(svc.QueryMachineNodes(v.MachineID))
}
