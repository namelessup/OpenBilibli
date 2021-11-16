package drawyoo

import (
	"context"
	"net/url"
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/reply/conf"
	"github.com/namelessup/bilibili/app/interface/main/reply/model/drawyoo"
	"github.com/namelessup/bilibili/library/log"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/xstr"
)

// Dao Dao
type Dao struct {
	url  string
	http *httpx.Client
}

// New New
func New(c *conf.Config) *Dao {
	d := &Dao{
		url:  "http://h.bilibili.com/api/pushS",
		http: httpx.NewClient(c.DrawyooHTTPClient),
	}
	return d
}

// Info Info
func (dao *Dao) Info(c context.Context, hid int64) (info *drawyoo.Drawyoo, err error) {
	params := url.Values{}
	params.Set("act", "getHidInfo")
	params.Set("hid", strconv.FormatInt(hid, 10))
	var res struct {
		State int                `json:"state"`
		Data  []*drawyoo.Drawyoo `json:"data"`
	}
	if err = dao.http.Post(c, dao.url, "", params, &res); err != nil {
		log.Error("drawyoo url(%v),err (%v)", dao.url+"?"+params.Encode(), err)
		return
	}
	if res.State != 200 || len(res.Data) == 0 {
		log.Error("drawyoo url (%v),err (%v)", dao.url+"?"+params.Encode(), err)
		return
	}
	info = res.Data[0]
	return
}

// Infos Infos
func (dao *Dao) Infos(c context.Context, hids []int64) (info map[int64]interface{}, err error) {
	params := url.Values{}
	params.Set("act", "getHidInfo")
	params.Set("hid", xstr.JoinInts(hids))
	var res struct {
		State int                `json:"state"`
		Data  []*drawyoo.Drawyoo `json:"data"`
	}
	if err = dao.http.Post(c, dao.url, "", params, &res); err != nil {
		log.Error("drawyoo url(%v),err (%v)", dao.url+"?"+params.Encode(), err)
		return
	}
	if res.State != 200 || len(res.Data) == 0 {
		log.Error("drawyoo url (%v),err (%v)", dao.url+"?"+params.Encode(), err)
		return
	}
	info = make(map[int64]interface{}, len(res.Data))
	for _, r := range res.Data {
		info[r.Hid] = r
	}
	return
}
