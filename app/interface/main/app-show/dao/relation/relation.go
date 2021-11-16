package relation

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-show/conf"
	relation "github.com/namelessup/bilibili/app/service/main/relation/model"
	relrpc "github.com/namelessup/bilibili/app/service/main/relation/rpc/client"
	"github.com/namelessup/bilibili/library/log"
)

// Dao is rpc dao.
type Dao struct {
	// relation rpc
	relRPC *relrpc.Service
}

// New new a relation dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// relation rpc
		relRPC: relrpc.New(c.RelationRPC),
	}
	return
}

// Relations fids relations
func (d *Dao) Relations(ctx context.Context, mid int64, fids []int64) (res map[int64]*relation.Following, err error) {
	arg := &relation.ArgRelations{
		Mid:  mid,
		Fids: fids,
	}
	if res, err = d.relRPC.Relations(ctx, arg); err != nil {
		log.Error("d.relRPC.Relations(%v) error(%v)", arg, err)
		res = nil
		return
	}
	return
}

// Stats fids stats
func (d *Dao) Stats(ctx context.Context, mids []int64) (res map[int64]*relation.Stat, err error) {
	arg := &relation.ArgMids{
		Mids: mids,
	}
	if res, err = d.relRPC.Stats(ctx, arg); err != nil {
		log.Error("d.relRPC.Stats(%v) error(%v)", arg, err)
		res = nil
		return
	}
	return
}
