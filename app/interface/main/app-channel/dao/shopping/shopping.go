package shopping

import (
	"context"
	"net/url"

	"github.com/namelessup/bilibili/app/interface/main/app-card/model/card/show"
	"github.com/namelessup/bilibili/app/interface/main/app-channel/conf"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"github.com/namelessup/bilibili/library/xstr"

	"github.com/pkg/errors"
)

const (
	_getCard = "/api/ticket/project/getcard"
)

// Dao is shopping dao.
type Dao struct {
	// http client
	client *bm.Client
	// live
	getCard string
}

// New new a shopping dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// http client
		client:  bm.NewClient(c.HTTPShopping),
		getCard: c.Host.Shopping + _getCard,
	}
	return d
}

func (d *Dao) Card(c context.Context, ids []int64) (rs map[int64]*show.Shopping, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	params := url.Values{}
	params.Set("id", xstr.JoinInts(ids))
	params.Set("for", "1")
	params.Set("price", "1")
	var res struct {
		Code int              `json:"errno"`
		Data []*show.Shopping `json:"data"`
	}
	if err = d.client.Get(c, d.getCard, ip, params, &res); err != nil {
		return
	}
	if res.Code != ecode.OK.Code() {
		err = errors.Wrap(err, d.getCard+"?"+params.Encode())
		return
	}
	rs = make(map[int64]*show.Shopping, len(res.Data))
	for _, r := range res.Data {
		rs[r.ID] = r
	}
	return
}
