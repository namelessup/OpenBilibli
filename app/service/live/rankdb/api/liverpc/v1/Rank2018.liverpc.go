// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Rank2018.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with github.com/namelessup/bilibili/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/Rank2018.proto
	v1/UserRank.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/namelessup/bilibili/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ==================
// Rank2018 Interface
// ==================

type Rank2018 interface {
	// * 获取上小时榜topN
	//
	GetHourRank(context.Context, *Rank2018GetHourRankReq) (*Rank2018GetHourRankResp, error)
}

// ========================
// Rank2018 Live Rpc Client
// ========================

type rank2018RpcClient struct {
	client *liverpc.Client
}

// NewRank2018RpcClient creates a Rpc client that implements the Rank2018 interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewRank2018RpcClient(client *liverpc.Client) Rank2018 {
	return &rank2018RpcClient{
		client: client,
	}
}

func (c *rank2018RpcClient) GetHourRank(ctx context.Context, in *Rank2018GetHourRankReq) (*Rank2018GetHourRankResp, error) {
	out := new(Rank2018GetHourRankResp)
	err := doRpcRequest(ctx, c.client, 1, "Rank2018.getHourRank", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRpcRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message) (err error) {
	err = client.Call(ctx, version, method, in, out)
	return
}
