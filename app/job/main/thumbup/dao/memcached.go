package dao

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/job/main/thumbup/model"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/xstr"
)

func statsKey(businessID, messageID int64) string {
	return fmt.Sprintf("m_%d_b_%d", messageID, businessID)
}

// AddStatsCache .
func (d *Dao) AddStatsCache(c context.Context, businessID int64, ms ...*model.Stats) (err error) {
	if len(ms) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	for _, m := range ms {
		if m == nil {
			continue
		}
		key := statsKey(businessID, m.ID)
		bs := xstr.JoinInts([]int64{m.Likes, m.Dislikes})
		item := memcache.Item{Key: key, Value: []byte(bs), Expiration: d.mcStatsExpire}
		if err = conn.Set(&item); err != nil {
			log.Error("conn.Set(%s) error(%v)", key, err)
			return
		}
	}
	return
}
