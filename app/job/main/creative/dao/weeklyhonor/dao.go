package weeklyhonor

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/job/main/creative/conf"
	arcrpc "github.com/namelessup/bilibili/app/service/main/archive/api/gorpc"
	upgrpc "github.com/namelessup/bilibili/app/service/main/up/api/v1"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
	"github.com/namelessup/bilibili/library/database/sql"
	binfoc "github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is creative dao.
type Dao struct {
	// config
	c *conf.Config
	// db
	db *sql.DB
	// httpClient
	httpClient *bm.Client
	// hbase
	hbase        *hbase.Client
	hbaseTimeOut time.Duration
	// rpc
	arc   *arcrpc.Service2
	infoc *binfoc.Infoc
	// grpc
	upClient upgrpc.UpClient
}

// New init api url
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:          c,
		db:         sql.NewMySQL(c.DB.Creative),
		httpClient: bm.NewClient(c.HTTPClient.Normal),
		// hbase
		hbase:        hbase.NewClient(c.HBaseOld.Config),
		hbaseTimeOut: time.Duration(time.Millisecond * 200),
		arc:          arcrpc.New2(c.ArchiveRPC),
		infoc:        binfoc.New(c.WeeklyHonorInfoc),
	}
	var err error
	d.upClient, err = upgrpc.NewClient(c.UpGRPCClient)
	if err != nil {
		panic(err)
	}
	return
}

// Ping creativeDb
func (d *Dao) Ping(c context.Context) (err error) {
	return d.pingMySQL(c)
}

// Close creativeDb
func (d *Dao) Close() (err error) {
	_ = d.infoc.Close()
	return d.db.Close()
}
