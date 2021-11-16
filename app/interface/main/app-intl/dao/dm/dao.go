package dm

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/app-intl/conf"
	dm "github.com/namelessup/bilibili/app/interface/main/dm2/model"
	dmrpc "github.com/namelessup/bilibili/app/interface/main/dm2/rpc/client"
	"github.com/namelessup/bilibili/library/net/metadata"

	"github.com/pkg/errors"
)

// Dao struct
type Dao struct {
	dmRPC *dmrpc.Service
}

// New a dao
func New(c *conf.Config) (d *Dao) {
	return &Dao{
		dmRPC: dmrpc.New(c.DMRPC),
	}
}

// SubjectInfos is.
func (d *Dao) SubjectInfos(c context.Context, typ int32, plat int8, oids ...int64) (res map[int64]*dm.SubjectInfo, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &dm.ArgOids{Type: typ, Plat: plat, Oids: oids, RealIP: ip}
	if res, err = d.dmRPC.SubjectInfos(c, arg); err != nil {
		err = errors.Wrapf(err, "%v", arg)
	}
	return
}
