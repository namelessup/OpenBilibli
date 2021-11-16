package broadcast

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/interface/main/app-resource/conf"
	pb "github.com/namelessup/bilibili/app/service/main/broadcast/api/grpc/v1"
	warden "github.com/namelessup/bilibili/app/service/main/broadcast/api/grpc/v1"
	wardenclient "github.com/namelessup/bilibili/app/service/main/broadcast/api/grpc/v1"
	"github.com/namelessup/bilibili/library/log"
)

type Dao struct {
	c *conf.Config
	// grpc
	rpcClient pb.ZergClient
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
	}
	var err error
	if d.rpcClient, err = wardenclient.NewClient(c.BroadcastRPC); err != nil {
		panic(fmt.Sprintf("BroadcastRPC warden.NewClient error (%+v)", err))
	}
	return
}

// ServerList warden server list
func (d *Dao) ServerList(ctx context.Context, platform string) (res *warden.ServerListReply, err error) {
	arg := &warden.ServerListReq{
		Platform: platform,
	}
	if res, err = d.rpcClient.ServerList(ctx, arg); err != nil {
		log.Error("d.rpcClient.ServerList error(%v)", err)
		return
	}
	return
}
