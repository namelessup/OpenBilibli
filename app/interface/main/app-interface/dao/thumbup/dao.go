package thumbup

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	thumbup "github.com/namelessup/bilibili/app/service/main/thumbup/model"
	thumbuprpc "github.com/namelessup/bilibili/app/service/main/thumbup/rpc/client"
	"github.com/namelessup/bilibili/library/net/metadata"

	"github.com/pkg/errors"
)

// Dao is tag dao
type Dao struct {
	thumbupRPC *thumbuprpc.Service
}

// New initial tag dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		thumbupRPC: thumbuprpc.New(c.ThumbupRPC),
	}
	return
}

// UserLikes user likes list
func (d *Dao) UserTotalLike(c context.Context, mid int64, business string, pn, ps int) (res []*thumbup.ItemLikeRecord, count int, err error) {
	var (
		likes *thumbup.UserTotalLike
		ip    = metadata.String(c, metadata.RemoteIP)
	)
	arg := &thumbup.ArgUserLikes{Mid: mid, Business: business, Pn: pn, Ps: ps, RealIP: ip}
	if likes, err = d.thumbupRPC.UserTotalLike(c, arg); err != nil {
		err = errors.Wrapf(err, "%v", arg)
		return
	}
	if likes != nil {
		res = likes.List
		count = likes.Total
	}
	return
}
