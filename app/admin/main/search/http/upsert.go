package http

import (
	"encoding/json"
	"strings"

	"github.com/namelessup/bilibili/app/admin/main/search/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func upsert(c *bm.Context) {
	up := &model.UpsertParams{}
	if err := c.Bind(up); err != nil {
		return
	}
	dataBody := map[string][]model.MapData{}
	decoder := json.NewDecoder(strings.NewReader(up.DataStr))
	decoder.UseNumber()
	if err := decoder.Decode(&dataBody); err != nil {
		log.Error("s.http.upsert(%v) json error(%v)", err, dataBody)
	}
	if len(dataBody) == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	for _, n := range dataBody {
		for _, m := range n {
			if err := m.NumberToInt64(); err != nil {
				log.Error("s.http.upsert(%v) to int64 error(%v)", err, m)
			}
		}
	}
	c.JSON(nil, svr.Upsert(c, up, dataBody))
}
