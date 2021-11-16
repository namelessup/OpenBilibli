package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/melloi/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func treesQuery(c *bm.Context) {
	c.JSON(srv.TreesQuery())
}

func treeNumQuery(c *bm.Context) {
	c.JSON(srv.TreeNumQuery())
}

func topHttpQuery(c *bm.Context) {
	c.JSON(srv.TopHttpQuery())
}

func topGrpcQuery(c *bm.Context) {
	c.JSON(srv.TopGrpcQuery())
}

func topSceneQuery(c *bm.Context) {
	c.JSON(srv.TopSceneQuery())
}

func topDeptQuery(c *bm.Context) {
	c.JSON(srv.TopDeptQuery())
}

func buildLineQuery(c *bm.Context) {
	rank := model.Rank{}
	summary := model.ReportSummary{}
	if err := c.BindWith(&rank, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}
	if err := c.BindWith(&summary, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(srv.BuildLineQuery(&rank, &summary))
}

func stateLineQuery(c *bm.Context) {
	c.JSON(srv.StateLineQuery())
}
