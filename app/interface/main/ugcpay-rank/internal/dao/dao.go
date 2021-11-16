package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/ugcpay-rank/internal/conf"
	ugcpay_rank "github.com/namelessup/bilibili/app/service/main/ugcpay-rank/api/v1"
)

// Dao dao
type Dao struct {
	ugcPayRankAPI ugcpay_rank.UGCPayRankClient
}

// New init mysql db
func New() (dao *Dao) {
	dao = &Dao{}
	var err error
	if dao.ugcPayRankAPI, err = ugcpay_rank.NewClient(conf.Conf.UGCPayRankGRPC); err != nil {
		panic(err)
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}
