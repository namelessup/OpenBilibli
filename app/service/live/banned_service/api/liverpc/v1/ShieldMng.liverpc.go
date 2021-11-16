// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/ShieldMng.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/namelessup/bilibili/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ===================
// ShieldMng Interface
// ===================

type ShieldMng interface {
	// * 查询用户是否被屏蔽
	//
	IsShieldUser(context.Context, *ShieldMngIsShieldUserReq) (*ShieldMngIsShieldUserResp, error)
}

// =========================
// ShieldMng Live Rpc Client
// =========================

type shieldMngRpcClient struct {
	client *liverpc.Client
}

// NewShieldMngRpcClient creates a Rpc client that implements the ShieldMng interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewShieldMngRpcClient(client *liverpc.Client) ShieldMng {
	return &shieldMngRpcClient{
		client: client,
	}
}

func (c *shieldMngRpcClient) IsShieldUser(ctx context.Context, in *ShieldMngIsShieldUserReq) (*ShieldMngIsShieldUserResp, error) {
	out := new(ShieldMngIsShieldUserResp)
	err := doRpcRequest(ctx, c.client, 1, "ShieldMng.is_shield_user", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
