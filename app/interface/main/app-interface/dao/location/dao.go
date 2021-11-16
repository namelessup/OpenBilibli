package location

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	locmdl "github.com/namelessup/bilibili/app/service/main/location/model"
	locrpc "github.com/namelessup/bilibili/app/service/main/location/rpc/client"
	"github.com/namelessup/bilibili/library/log"
)

// Dao is location dao.
type Dao struct {
	// rpc
	locRPC *locrpc.Service
}

// New new a location dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// rpc
		locRPC: locrpc.New(c.LocationRPC),
	}
	return
}

func (d *Dao) Info(c context.Context, ipaddr string) (info *locmdl.Info, err error) {
	if info, err = d.locRPC.Info(c, &locmdl.ArgIP{IP: ipaddr}); err != nil {
		log.Error("%v", err)
	}
	return
}
