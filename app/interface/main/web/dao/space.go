package dao

import (
	"context"
	"net/url"
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/web/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

// TopPhoto getTopPhoto from space
func (d *Dao) TopPhoto(c context.Context, mid int64) (space *model.Space, err error) {
	var (
		params   = url.Values{}
		remoteIP = metadata.String(c, metadata.RemoteIP)
	)
	params.Set("mid", strconv.FormatInt(mid, 10))
	var res struct {
		Code int `json:"code"`
		model.Space
	}
	if err = d.httpR.Get(c, d.spaceTopPhotoURL, remoteIP, params, &res); err != nil {
		log.Error("TopPhoto space url(%s) error(%v)", d.spaceTopPhotoURL+"?"+params.Encode(), err)
		return
	}
	if res.Code != 0 {
		log.Error("TopPhoto space url(%s) error(%v)", d.spaceTopPhotoURL+"?"+params.Encode(), res.Code)
		err = ecode.Int(res.Code)
		return
	}
	space = &res.Space
	return
}
