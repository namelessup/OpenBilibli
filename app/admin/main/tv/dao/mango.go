package dao

import (
	"context"
	"encoding/json"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/tv/model"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	xtime "github.com/namelessup/bilibili/library/time"
)

const (
	_mangoKey = "mango_cms_recom"
	_delRecom = "UPDATE mango_recom SET deleted = 1 WHERE id = ?"
	_maxOrder = "SELECT MAX(rorder) AS ord FROM mango_recom WHERE deleted = 0"
)

//MangoRecom is used to set mango recom cache in MC
func (d *Dao) MangoRecom(c context.Context, ids []int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	mcItem := &model.MRecomMC{
		RIDs:    ids,
		Pubtime: xtime.Time(time.Now().Unix()),
	}
	itemJSON := &memcache.Item{
		Key:        _mangoKey,
		Object:     mcItem,
		Flags:      memcache.FlagJSON,
		Expiration: 0,
	}
	if err = conn.Set(itemJSON); err != nil {
		log.Error("MangoRecom Ids %v, Err %v", ids, err)
	}
	return
}

// GetMRecom get mango recom mc data
func (d *Dao) GetMRecom(c context.Context) (res *model.MRecomMC, err error) {
	var item *memcache.Item
	conn := d.mc.Get(c)
	defer conn.Close()
	if item, err = conn.Get(_mangoKey); err != nil {
		log.Error("GetMRecom MangoKey, Err %v", _mangoKey, err)
		return
	}
	err = json.Unmarshal(item.Value, &res)
	return
}

// DelMRecom deletes an mango recom position
func (d *Dao) DelMRecom(c *bm.Context, id int64) (err error) {
	if err = d.DB.Exec(_delRecom, id).Error; err != nil {
		log.Error("DelMRecom Error %v", err)
	}
	return
}

// MaxOrder picks the max rorder from the table
func (d *Dao) MaxOrder(ctx context.Context) int {
	var maxR = new(struct {
		Ord int
	})
	if err := d.DB.Raw(_maxOrder).Scan(&maxR).Error; err != nil {
		log.Error("MaxOrder Error %v", err)
		return 0
	}
	return maxR.Ord
}
