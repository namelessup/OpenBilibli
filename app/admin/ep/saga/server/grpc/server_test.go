package grpc

import (
	"context"
	"testing"

	"github.com/namelessup/bilibili/app/admin/ep/saga/api/grpc/v1"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/rpc/warden/resolver"
	"github.com/namelessup/bilibili/library/net/rpc/warden/resolver/direct"
)

func TestGRPC(t *testing.T) {
	resolver.Register(direct.New())
	conn, err := warden.NewClient(nil).Dial(context.Background(), "direct://default/127.0.0.1:9000")
	if err != nil {
		t.Fatal(err)
	}
	client := v1.NewSagaAdminClient(conn)
	if _, err = client.PushMsg(context.Background(), &v1.PushMsgReq{Username: []string{"wuwei"}, Content: "test"}); err != nil {
		t.Error(err)
	}
}
