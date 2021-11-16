package dao

import (
	"context"
	recsys "github.com/namelessup/bilibili/app/service/bbq/recsys/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/bbq/recsys/model"
	"github.com/namelessup/bilibili/app/service/bbq/recsys/service/retrieve"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/log"
	"strconv"
	"strings"
)

//DownGradeRecall ...
func (d *Dao) DownGradeRecall(c context.Context) (response *recsys.RecsysResponse, err error) {
	conn := d.redis.Get(c)
	defer conn.Close()

	key := retrieve.RecallHotDefault

	var str string
	if str, err = redis.String(conn.Do("GET", key)); err != nil {
		if err == redis.ErrNil {
			err = nil
		} else {
			log.Errorw(c, "recall", "get hot recall error", "err", err)
		}
	}
	response = new(recsys.RecsysResponse)
	response.Message = make(map[string]string)
	records := make([]*recsys.RecsysRecord, 0)

	response.Message[model.ResponseDownGrade] = "2"
	for _, svidStr := range strings.Split(str, ",") {
		svid, _ := strconv.ParseInt(svidStr, 10, 64)
		record := &recsys.RecsysRecord{
			Svid:  svid,
			Score: 0,
			Map:   make(map[string]string),
		}
		record.Map[model.RecallClasses] = retrieve.HotRecall
		records = append(records, record)
	}

	key = retrieve.RecallOpVideoKey

	if str, err = redis.String(conn.Do("GET", key)); err != nil {
		if err == redis.ErrNil {
			err = nil
		} else {
			log.Errorw(c, "recall", "get selection recall error", "err", err)
		}
	}
	response = new(recsys.RecsysResponse)
	response.Message = make(map[string]string)
	for _, svidStr := range strings.Split(str, ",") {
		svid, _ := strconv.ParseInt(svidStr, 10, 64)
		record := &recsys.RecsysRecord{
			Svid:  svid,
			Score: 0,
			Map:   make(map[string]string),
		}
		record.Map[model.RecallClasses] = retrieve.SelectionRecall
		records = append(records, record)
	}
	response.List = records
	return
}
