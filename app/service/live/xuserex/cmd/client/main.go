package main

import (
	"context"
	"flag"
	"github.com/namelessup/bilibili/app/service/live/xuserex/api/grpc/v1"
	"log"
	"time"

	"fmt"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	xtime "github.com/namelessup/bilibili/library/time"
)

var name, addr string

func init() {
	flag.StringVar(&name, "name", "lily", "name")
	flag.StringVar(&addr, "addr", "127.0.0.1:9004", "server addr")
}

func main() {
	flag.Parse()
	cfg := &warden.ClientConfig{
		Dial:    xtime.Duration(time.Second * 3),
		Timeout: xtime.Duration(time.Second * 3),
	}
	cc, err := warden.NewClient(cfg).Dial(context.Background(), addr)
	if err != nil {
		log.Fatalf("new client failed!err:=%v", err)
		return
	}

	client := v1.NewRoomNoticeClient(cc)

	resp, err := client.BuyGuard(context.Background(), &v1.RoomNoticeBuyGuardReq{
		Uid:      10000,
		TargetId: 11,
	})
	fmt.Printf("****** buyguard :******* \n %v \n %v \n", resp, err)
}
