package audio

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/app-intl/conf"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"net/url"

	"github.com/namelessup/bilibili/app/interface/main/app-intl/model/view"
	"github.com/namelessup/bilibili/library/xstr"

	"github.com/pkg/errors"
)

const (
	_audioByCids = "/audio/music-service-c/internal/songs-by-cids"
)

// Dao is archive dao.
type Dao struct {
	// http client
	client         *bm.Client
	audioByCidsURL string
}

// New new a archive dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// http client
		client:         bm.NewClient(c.HTTPAudio),
		audioByCidsURL: c.Host.APICo + _audioByCids,
	}
	return
}

// AudioByCids is.
func (d *Dao) AudioByCids(c context.Context, cids []int64) (vam map[int64]*view.Audio, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	params := url.Values{}
	params.Set("cids", xstr.JoinInts(cids))
	var res struct {
		Code int                   `json:"code"`
		Data map[int64]*view.Audio `json:"data"`
	}
	if err = d.client.Get(c, d.audioByCidsURL, ip, params, &res); err != nil {
		return
	}
	if res.Code != ecode.OK.Code() {
		err = errors.Wrap(ecode.Int(res.Code), d.audioByCidsURL+"?"+params.Encode())
		return
	}
	vam = res.Data
	return
}
