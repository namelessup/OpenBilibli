package sp

import (
	"context"
	"net/url"
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	"github.com/namelessup/bilibili/app/interface/main/app-interface/model/sp"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

const (
	_specil = "/sp/list"
)

// Dao is favorite dao
type Dao struct {
	client *httpx.Client
	specil string
}

// New initial favorite dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		client: httpx.NewClient(c.HTTPClient),
		specil: c.Host.APICo + _specil,
	}
	return
}

// Specil get specil from old BiliWEB api.
func (d *Dao) Specil(c context.Context, accessKey, actionKey, device, mobiApp, platform string, build, pn, ps int) (res *sp.Specil, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	params := url.Values{}
	params.Set("access_key", accessKey)
	params.Set("actionKey", actionKey)
	params.Set("build", strconv.Itoa(build))
	params.Set("device", device)
	params.Set("mobi_app", mobiApp)
	params.Set("page", strconv.Itoa(pn))
	params.Set("pagesize", strconv.Itoa(ps))
	params.Set("platform", platform)
	err = d.client.Get(c, d.specil, ip, params, &res)
	return
}
