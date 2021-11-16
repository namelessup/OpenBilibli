package service

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	conf "github.com/namelessup/bilibili/app/interface/main/kvo/conf"
	"github.com/namelessup/bilibili/app/interface/main/kvo/model/module"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/container/pool"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/net/netutil/breaker"
	xtime "github.com/namelessup/bilibili/library/time"

	. "github.com/smartystreets/goconvey/convey"
)

func getService() *Service {
	s := New(&conf.Config{
		Rule: &conf.Rule{DocLimit: 1024 * 1024 * 1024},
		Memcache: &conf.KvoMemcache{
			Kvo: &memcache.Config{
				Config: &pool.Config{
					Active:      10,
					Idle:        4,
					IdleTimeout: xtime.Duration(time.Second),
				},
				Name:         "kvo",
				Proto:        "tcp",
				Addr:         "127.0.0.1:11211",
				DialTimeout:  xtime.Duration(time.Second),
				ReadTimeout:  xtime.Duration(time.Second),
				WriteTimeout: xtime.Duration(time.Second),
			},
			Expire: xtime.Duration(time.Hour),
		},
		Mysql: &sql.Config{
			Addr:         "localhost:3306",
			DSN:          "root:123@tcp(localhost:3306)/bilibili?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4",
			Active:       10,
			Idle:         4,
			IdleTimeout:  xtime.Duration(time.Second),
			QueryTimeout: xtime.Duration(time.Second),
			ExecTimeout:  xtime.Duration(time.Second),
			TranTimeout:  xtime.Duration(time.Second),
			Breaker: &breaker.Config{
				Window:  xtime.Duration(time.Second),
				Sleep:   xtime.Duration(time.Second),
				Bucket:  10,
				Ratio:   0.5,
				Request: 100,
			},
		},
	})
	return s
}

func TestAddDocument(t *testing.T) {
	Convey("", t, func() {
		s := getService()
		p := &module.Player{
			PlayerWebDanmakuAutoscaling: false,
		}
		bs, _ := json.Marshal(p)
		_, err := s.AddDocument(context.Background(), 1, "player", string(bs), 0, 0, time.Now())
		So(nil, ShouldEqual, err)
	})
}

func TestDocument(t *testing.T) {
	Convey("", t, func() {
		s := getService()
		_, err := s.Document(context.Background(), 1, "player", 1234, 12345)
		So(nil, ShouldEqual, err)
	})
}
