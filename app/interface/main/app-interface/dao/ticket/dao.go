package ticket

import (
	"context"
	"net/url"
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"

	"github.com/pkg/errors"
)

const _favCount = "/api/ticket/user/favcountinner"

type Dao struct {
	client   *bm.Client
	favCount string
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		client:   bm.NewClient(c.HTTPClient),
		favCount: c.Host.Show + _favCount,
	}
	return
}

func (d *Dao) FavCount(c context.Context, mid int64) (count int32, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	params := url.Values{}
	params.Set("mid", strconv.FormatInt(mid, 10))
	var res struct {
		Code int   `json:"errno"`
		Data int32 `json:"data"`
	}
	if err = d.client.Get(c, d.favCount, ip, params, &res); err != nil {
		return
	}
	if res.Code != ecode.OK.Code() {
		err = errors.Wrap(ecode.Int(res.Code), d.favCount+"?"+params.Encode())
		return
	}
	count = res.Data
	return
}
