package http

import (
	"github.com/namelessup/bilibili/app/interface/main/esports/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func game(c *bm.Context) {
	p := new(model.ParamGame)
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(eSvc.Game(c, p))
}

func items(c *bm.Context) {
	p := new(model.ParamLeidas)
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(eSvc.Items(c, p))
}

func heroes(c *bm.Context) {
	p := new(model.ParamLeidas)
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(eSvc.Heroes(c, p))
}

func abilities(c *bm.Context) {
	p := new(model.ParamLeidas)
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(eSvc.Abilities(c, p))
}
func players(c *bm.Context) {
	p := new(model.ParamLeidas)
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(eSvc.Players(c, p))
}
