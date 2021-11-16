package resource

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	resmdl "github.com/namelessup/bilibili/app/service/main/resource/model"
	resrpc "github.com/namelessup/bilibili/app/service/main/resource/rpc/client"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

type Dao struct {
	c *conf.Config
	// rpc
	resRPC *resrpc.Service
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// rpc
		resRPC: resrpc.New(c.ResourceRPC),
	}
	return
}

// Banner get search banner
func (d *Dao) Banner(c context.Context, mobiApp, device, network, channel, buvid, adExtra, resIDStr string, build int, plat int8, mid int64) (res map[int][]*resmdl.Banner, err error) {
	var (
		bs *resmdl.Banners
		ip = metadata.String(c, metadata.RemoteIP)
	)
	arg := &resmdl.ArgBanner{
		MobiApp: mobiApp,
		Device:  device,
		Network: network,
		Channel: channel,
		IP:      ip,
		Buvid:   buvid,
		AdExtra: adExtra,
		ResIDs:  resIDStr,
		Build:   build,
		Plat:    plat,
		MID:     mid,
		IsAd:    true,
	}
	if bs, err = d.resRPC.Banners(c, arg); err != nil || bs == nil {
		log.Error("d.resRPC.Banners(%v) error(%v) or bs is nil", arg, err)
		return
	}
	if len(bs.Banner) > 0 {
		res = bs.Banner
	}
	return
}
