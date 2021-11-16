package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/namelessup/bilibili/app/service/live/recommend/api/grpc/v1"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	xtime "github.com/namelessup/bilibili/library/time"
)

var name, addr string

func init() {
	flag.StringVar(&name, "name", "lily", "name")
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
	client := v1.NewRecommendClient(cc)
	resp, err := client.RandomRecsByUser(context.Background(), &v1.GetRandomRecReq{
		Uid: 4158272, Count: 5,
	})
	if err != nil {
		log.Fatalf("say hello failed!err:=%v", err)
		return
	}
	fmt.Printf("got HelloReply:%+v", resp)
}
