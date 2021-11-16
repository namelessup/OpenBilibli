package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/block/conf"
	rpcaccount "github.com/namelessup/bilibili/app/service/main/account/rpc/client"
	rpcfigure "github.com/namelessup/bilibili/app/service/main/figure/rpc/client"
	rpcspy "github.com/namelessup/bilibili/app/service/main/spy/rpc/client"
	"github.com/namelessup/bilibili/library/cache/memcache"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/pkg/errors"
)

// Dao .
type Dao struct {
	mc         *memcache.Pool
	db         *xsql.DB
	httpClient *bm.Client
	spyRPC     *rpcspy.Service
	figureRPC  *rpcfigure.Service
	accountRPC *rpcaccount.Service3
}

// New init mysql db
func New() (dao *Dao) {
	dao = &Dao{
		mc:         memcache.NewPool(conf.Conf.Memcache),
		db:         xsql.NewMySQL(conf.Conf.MySQL),
		httpClient: bm.NewClient(conf.Conf.HTTPClient),
		spyRPC:     rpcspy.New(conf.Conf.RPCClients.Spy),
		figureRPC:  rpcfigure.New(conf.Conf.RPCClients.Figure),
		accountRPC: rpcaccount.New3(conf.Conf.RPCClients.Account),
	}
	return
}

// BeginTX .
func (d *Dao) BeginTX(c context.Context) (tx *xsql.Tx, err error) {
	if tx, err = d.db.Begin(c); err != nil {
		err = errors.WithStack(err)
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		return
	}
	return
}
