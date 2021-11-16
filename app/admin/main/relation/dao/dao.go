package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/relation/conf"
	accountGRPC "github.com/namelessup/bilibili/app/service/main/account/api"
	relationGRPC "github.com/namelessup/bilibili/app/service/main/relation/api"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
	"github.com/namelessup/bilibili/library/database/orm"

	"github.com/jinzhu/gorm"
)

// Dao dao
type Dao struct {
	c              *conf.Config
	ReadORM        *gorm.DB
	accountClient  accountGRPC.AccountClient
	relationClient relationGRPC.RelationClient
	hbase          *hbase.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:       c,
		ReadORM: orm.NewMySQL(c.ORM.Read),
		hbase:   hbase.NewClient(c.LogHBase),
	}
	relationGRPCClient, err := relationGRPC.NewClient(c.RelationGRPC)
	if err != nil {
		panic(err)
	}

	accountGRPCClient, err := accountGRPC.NewClient(c.AccountGRPC)
	if err != nil {
		panic(err)
	}
	dao.relationClient = relationGRPCClient
	dao.accountClient = accountGRPCClient
	dao.initORM()
	return
}

func (dao *Dao) initORM() {
	dao.ReadORM.LogMode(true)
}

// Close close the resource.
func (dao *Dao) Close() {
	dao.ReadORM.Close()
}

// Ping dao ping
func (dao *Dao) Ping(c context.Context) error {
	return dao.ReadORM.DB().PingContext(c)
}
