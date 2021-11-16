package geetest

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/videoup/conf"
	"github.com/namelessup/bilibili/app/interface/main/videoup/model/geetest"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

const (
	_validate = "/validate.php"
)

// Dao is account dao.
type Dao struct {
	c *conf.Config
	// url
	validateURI string
	// http client
	clientX *httpx.Client
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:           c,
		validateURI: c.Host.Geetest + _validate,
		clientX:     httpx.NewClient(c.HTTPClient.Read),
	}
	return
}

// Validate recheck the challenge code and get to seccode
func (d *Dao) Validate(c context.Context, challenge, seccode, clientType, captchaID string, mid int64) (res *geetest.ValidateRes, err error) {
	params := url.Values{}
	params.Set("seccode", seccode)
	params.Set("challenge", challenge)
	params.Set("captchaid", captchaID)
	params.Set("client_type", clientType)
	params.Set("ip_address", metadata.String(c, metadata.RemoteIP))
	params.Set("json_format", "1")
	params.Set("sdk", "golang_3.0.0")
	params.Set("user_id", strconv.FormatInt(mid, 10))
	params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	req, err := http.NewRequest("POST", d.validateURI, strings.NewReader(params.Encode()))
	if err != nil {
		log.Error("http.NewRequest error(%v) | uri(%s) params(%s)", err, d.validateURI, params.Encode())
		err = ecode.CreativeGeetestAPIErr
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err = d.clientX.Do(c, req, &res); err != nil {
		log.Error("d.client.Do error(%v)", err)
		err = ecode.CreativeGeetestAPIErr
		return
	}
	return
}
