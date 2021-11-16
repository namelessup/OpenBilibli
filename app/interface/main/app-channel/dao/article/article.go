package article

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-channel/conf"
	article "github.com/namelessup/bilibili/app/interface/openplatform/article/model"
	artrpc "github.com/namelessup/bilibili/app/interface/openplatform/article/rpc/client"
	"github.com/namelessup/bilibili/library/net/metadata"

	"github.com/pkg/errors"
)

// Dao is archive dao.
type Dao struct {
	// rpc
	artRPC *artrpc.Service
}

// New new a archive dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// rpc
		artRPC: artrpc.New(c.ArticleRPC),
	}
	return
}

func (d *Dao) Articles(c context.Context, aids []int64) (ms map[int64]*article.Meta, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	arg := &article.ArgAids{Aids: aids, RealIP: ip}
	if ms, err = d.artRPC.ArticleMetas(c, arg); err != nil {
		err = errors.Wrapf(err, "%v", aids)
	}
	return
}
