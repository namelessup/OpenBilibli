package bangumi

import (
	"context"

	seasongrpc "github.com/namelessup/bilibili/app/service/openplatform/pgc-season/api/grpc/season/v1"
	"github.com/namelessup/bilibili/library/log"
)

// Cards get bangumis.
func (d *Dao) Cards(ctx context.Context, seasonIds []int32) (res map[int32]*seasongrpc.CardInfoProto, err error) {
	arg := &seasongrpc.SeasonInfoReq{
		SeasonIds: seasonIds,
	}
	info, err := d.rpcClient.Cards(ctx, arg)
	if err != nil {
		log.Error("d.rpcClient.Cards error(%v)", err)
		return nil, err
	}
	res = info.Cards
	return
}
