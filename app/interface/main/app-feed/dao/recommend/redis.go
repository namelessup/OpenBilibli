package recommend

import (
	"context"
	"strconv"

	"github.com/namelessup/bilibili/library/cache/redis"

	"github.com/pkg/errors"
)

const (
	_prefixPos = "p_"
)

func keyPos(mid int64) string {
	return _prefixPos + strconv.FormatInt(mid%100000, 10)
}

func (d *Dao) PositionCache(c context.Context, mid int64) (pos int, err error) {
	var (
		key  = keyPos(mid)
		conn = d.redis.Get(c)
	)
	defer conn.Close()
	if pos, err = redis.Int(conn.Do("HGET", key, mid)); err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		err = errors.Wrapf(err, "conn.Do(HGET,%s,%d)", key, mid)
	}
	return
}

func (d *Dao) AddPositionCache(c context.Context, mid int64, pos int) (err error) {
	var (
		key  = keyPos(mid)
		conn = d.redis.Get(c)
	)
	defer conn.Close()
	if _, err = conn.Do("HSET", key, mid, pos); err != nil {
		err = errors.Wrapf(err, "conn.Do(HSET,%s,%d,%d)", key, mid, pos)
	}
	return
}
