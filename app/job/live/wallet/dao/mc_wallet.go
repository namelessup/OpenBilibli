package dao

import (
	"context"
	"fmt"
	mc "github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_walletMcKey = "wu:%d"
)

func mcKey(uid int64) string {
	return fmt.Sprintf(_walletMcKey, uid)
}

// DelWalletCache 删除等级缓存
func (d *Dao) DelWalletCache(c context.Context, uid int64) (err error) {
	key := mcKey(uid)
	conn := d.mc.Get(c)
	defer conn.Close()

	if err = conn.Delete(key); err == mc.ErrNotFound {
		err = nil
	} else if err != nil {
		log.Error("[dao.mc_wallet|DelWalletCache] conn.Delete(%s) error(%v)", key, err)
	}
	return
}
