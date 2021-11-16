package v1

import (
	"context"

	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"google.golang.org/grpc"
)

// DiscoveryID season
const DiscoveryID = "season.service"

// NewClient new identify grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (SeasonClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+DiscoveryID)
	if err != nil {
		return nil, err
	}
	return NewSeasonClient(conn), nil
}
