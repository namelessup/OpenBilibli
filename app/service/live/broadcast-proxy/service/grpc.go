package service

import (
	"errors"
	v1pb "github.com/namelessup/bilibili/app/service/live/broadcast-proxy/api/v1"
	"github.com/namelessup/bilibili/app/service/live/broadcast-proxy/server"
	v1srv "github.com/namelessup/bilibili/app/service/live/broadcast-proxy/service/v1"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	xtime "github.com/namelessup/bilibili/library/time"
	"google.golang.org/grpc"
	"time"
)

func NewGrpcService(p *server.BroadcastProxy, d *server.CometDispatcher) (*warden.Server, error) {
	if p == nil || d == nil {
		return nil, errors.New("empty proxy")
	}
	ws := warden.NewServer(&warden.ServerConfig{
		Timeout: xtime.Duration(30 * time.Second),
	}, grpc.MaxRecvMsgSize(1024 * 1024 * 1024), grpc.MaxSendMsgSize(1024 * 1024 * 1024))
	v1pb.RegisterDanmakuServer(ws.Server(), v1srv.NewDanmakuService(p, d))
	ws, err := ws.Start()
	if err != nil {
		return nil, err
	}
	return ws, nil
}
