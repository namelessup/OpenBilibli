package api

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "community.service.coin"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (CoinClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewCoinClient(cc), nil
}

//go:generate $GOPATH/src/github.com/namelessup/bilibili/app/tool/warden/protoc.sh
