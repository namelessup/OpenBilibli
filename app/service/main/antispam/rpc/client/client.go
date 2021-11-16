package rpc

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/antispam/model"
	"github.com/namelessup/bilibili/library/net/rpc"
)

const (
	_appid = "antispam.service"
)

// Client .
type Client struct {
	*rpc.Client2
}

// NewClient .
func NewClient(c *rpc.ClientConfig) *Client {
	s := &Client{}
	s.Client2 = rpc.NewDiscoveryCli(_appid, c)
	return s
}

// Filter .
func (cli *Client) Filter(ctx context.Context, arg *model.Suspicious) (res *model.SuspiciousResp, err error) {
	err = cli.Call(ctx, "Filter.Check", arg, &res)
	return
}
