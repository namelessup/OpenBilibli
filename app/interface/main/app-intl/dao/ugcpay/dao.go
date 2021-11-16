package ugcpay

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/interface/main/app-intl/conf"
	ugcpay "github.com/namelessup/bilibili/app/service/main/ugcpay/api/grpc/v1"
	"github.com/namelessup/bilibili/library/log"
)

// Dao is ugcpay dao.
type Dao struct {
	// rpc
	ugcpayRPC ugcpay.UGCPayClient
}

// New new a archive dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{}
	var err error
	d.ugcpayRPC, err = ugcpay.NewClient(c.UGCpayClient)
	if err != nil {
		panic(fmt.Sprintf("ugcpay NewClient error(%v)", err))
	}
	return
}

// AssetRelation is
func (d *Dao) AssetRelation(c context.Context, aid, mid int64) (relation *ugcpay.AssetRelationResp, err error) {
	if relation, err = d.ugcpayRPC.AssetRelation(c, &ugcpay.AssetRelationReq{Oid: aid, Mid: mid, Otype: "archive"}); err != nil {
		log.Error("d.ugcpayRPC.AssetRelationDetail(%d) error(%+v)", aid, err)
		return
	}
	return
}
