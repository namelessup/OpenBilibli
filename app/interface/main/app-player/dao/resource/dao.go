package resource

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-player/conf"
	resrpc "github.com/namelessup/bilibili/app/service/main/resource/rpc/client"
	"github.com/namelessup/bilibili/library/log"
)

type Dao struct {
	c *conf.Config
	// rpc
	resRPC *resrpc.Service
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// rpc
		resRPC: resrpc.New(c.ResourceRPC),
	}
	return
}

// PasterCID get all paster cid.
func (d *Dao) PasterCID(c context.Context) (cids map[int64]int64, err error) {
	if cids, err = d.resRPC.PasterCID(c); err != nil {
		log.Error("d.resRPC.PasterCID() error(%v)", err)
	}
	return
}
