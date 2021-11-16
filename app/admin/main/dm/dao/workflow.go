package dao

import (
	"context"
	"net/url"
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

const (
	_workFlowAppealDelete = "/x/internal/workflow/appeal/v3/delete"
)

// WorkFlowAppealDelete .
func (d *Dao) WorkFlowAppealDelete(c context.Context, bid, oid, subtitleID int64) (err error) {
	var (
		res struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		params = url.Values{}
		ip     = metadata.String(c, metadata.RemoteIP)
	)
	params.Set("bid", strconv.FormatInt(bid, 10))
	params.Set("oid", strconv.FormatInt(oid, 10))
	params.Set("eid", strconv.FormatInt(subtitleID, 10))
	if err = d.httpCli.Post(c, d.workFlowURI, ip, params, &res); err != nil {
		log.Error("WorkFlowTagList(params:%v),error(%v)", params, err)
		return
	}
	if err = ecode.Int(res.Code); err != ecode.OK {
		log.Error("WorkFlowTagList(params:%v),error(%v)", params, err)
		return
	}
	err = nil
	return
}
