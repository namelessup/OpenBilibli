package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/dm2/model"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_dataRankURI = "/data/rank/recent_region-%d-%d.json"
)

// RankList get data rank by tid
func (d *Dao) RankList(c context.Context, tid int64, day int32) (resp *model.RankRecentResp, err error) {
	if err = d.httpCli.RESTfulGet(c, d.conf.Host.DataRank+_dataRankURI, "", nil, &resp, tid, day); err != nil {
		log.Error("RankList(tid:%v,day:%v),error(%v)", tid, day, err)
		return
	}
	return
}
