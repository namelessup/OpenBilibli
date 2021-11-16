package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/namelessup/bilibili/app/service/main/filter/api/grpc/v1"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	xtime "github.com/namelessup/bilibili/library/time"
	"log"
	"time"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "127.0.0.1:9000", "server addr")
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
	client := v1.NewFilterClient(cc)
	resp, err := client.Filter(context.Background(), &v1.FilterReq{
		Area:    "reply",
		Message: "习大大",
	})
	if err != nil {
		log.Fatalf("filter failed!err:=%+v", err)
		return
	}
	fmt.Printf("got FilterReply:%+v", resp)
}
