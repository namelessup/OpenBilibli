package service

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/namelessup/bilibili/app/admin/main/apm/conf"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// PlatformProxy proxy.
func (s *Service) PlatformProxy(c context.Context, method, loc string, params url.Values) (data interface{}, err error) {
	uri := conf.Conf.Host.MANAGERCo + "/" + loc
	data, err = s.pcurl(c, method, uri, loc, params)
	return
}

func (s *Service) pcurl(c context.Context, method, uri, loc string, params url.Values) (data json.RawMessage, err error) {
	var res struct {
		Code    int             `json:"code"`
		Data    json.RawMessage `json:"data"`
		Message string          `json:"message"`
	}
	if method == "GET" {
		if err = s.client.Get(c, uri, "", params, &res); err != nil {
			log.Error("apmSvc.PlatformProxy get url:"+uri+" params:(%v) error(%v)", params.Encode(), err)
			return
		}
	} else {
		if err = s.client.Post(c, uri, "", params, &res); err != nil {
			log.Error("apmSvc.PlatformProxy post url:"+uri+" params:(%v) error(%v)", params.Encode(), err)
			return
		}
	}
	if res.Code != 0 {
		err = ecode.Int(res.Code)
		log.Error("apmSvc.PlatformProxy url:"+uri+" params:(%v) ecode(%v)", params.Encode(), err)
		return
	}
	data = res.Data
	return
}
