package card

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/interface/main/app-show/model/card"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_hotTenprefix = "%d_hclist_v2"
)

func getHotKey(i int) string {
	return fmt.Sprintf(_hotTenprefix, i)
}

// AddPopularCardTenCache add mc list
func (d *Dao) AddPopularCardTenCache(c context.Context, i int, cards []*card.PopularCard) (err error) {
	var (
		key  = getHotKey(i)
		conn = d.mc.Get(c)
	)
	if err = conn.Set(&memcache.Item{Key: key, Object: cards, Flags: memcache.FlagJSON, Expiration: 0}); err != nil {
		log.Error("addCards d.mc.Set(%s,%v) error(%v)", key, cards, err)
	}
	conn.Close()
	return
}

// PopularCardTenCache get cards list
func (d *Dao) PopularCardTenCache(c context.Context, i int) (cards []*card.PopularCard, err error) {
	var (
		key  = getHotKey(i)
		conn = d.mc.Get(c)
		r    *memcache.Item
	)
	defer conn.Close()
	if r, err = conn.Get(key); err != nil {
		log.Error("cardsCache MemchDB.Get(%s) error(%v)", key, err)
		return
	}
	if err = conn.Scan(r, &cards); err != nil {
		log.Error("r.Scan(%s) error(%v)", r.Value, err)
	}
	return
}
