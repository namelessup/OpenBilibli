package dao

import (
	"context"
	"net/http"
	"net/url"

	"github.com/namelessup/bilibili/app/service/main/identify/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

// AccessCookie .
func (d *Dao) AccessCookie(c context.Context, cookie string) (res *model.IdentifyInfo, err error) {
	params := url.Values{}
	// new request
	req, err := d.client.NewRequest(http.MethodGet, d.cookieURI, metadata.String(c, metadata.RemoteIP), params)
	if err != nil {
		log.Error("client.NewRequest(GET, %s) error(%v)", req.URL.String(), err)
		return
	}
	req.Header.Set("Cookie", cookie)
	var response struct {
		Code int                `json:"code"`
		Data model.IdentifyInfo `json:"data"`
	}
	if err = d.client.Do(c, req, &response); err != nil {
		log.Error("client.Do(%s) error(%v)", req.URL.String(), err)
		return
	}
	if response.Code != ecode.OK.Code() {
		log.Warn("identify auth url(%s) code(%d)", req.URL.String(), response.Code)
		err = ecode.Int(response.Code)
		return
	}
	res = &response.Data
	return
}

// AccessToken .
func (d *Dao) AccessToken(c context.Context, accesskey string) (res *model.IdentifyInfo, err error) {
	params := url.Values{}
	params.Set("access_key", accesskey)
	// new request
	req, err := d.client.NewRequest(http.MethodGet, d.tokenURI, metadata.String(c, metadata.RemoteIP), params)
	if err != nil {
		log.Error("client.NewRequest(GET, %s) error(%v)", req.URL.String(), err)
		return
	}
	var response struct {
		Code int                `json:"code"`
		Data model.IdentifyInfo `json:"data"`
	}
	if err = d.client.Do(c, req, &response); err != nil {
		log.Error("client.Do(%s) error(%v)", req.URL.String(), err)
		return
	}
	if response.Code != 0 {
		log.Warn("identify auth url(%s) code(%d)", req.URL.String(), response.Code)
		err = ecode.Int(response.Code)
		return
	}
	res = &response.Data
	return
}
