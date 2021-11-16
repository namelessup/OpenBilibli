// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/FansMedal.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with github.com/namelessup/bilibili/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/FansMedal.proto
	v1/Medal.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/namelessup/bilibili/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ===================
// FansMedal Interface
// ===================

type FansMedal interface {
	// * 获取已佩戴的勋章
	//
	GetWearedMedal(context.Context, *FansMedalGetWearedMedalReq) (*FansMedalGetWearedMedalResp, error)

	// * 用户卡
	// 基于某房间|主播的 用户卡片信息
	TargetsWithMedal(context.Context, *FansMedalTargetsWithMedalReq) (*FansMedalTargetsWithMedalResp, error)
}

// =========================
// FansMedal Live Rpc Client
// =========================

type fansMedalRpcClient struct {
	client *liverpc.Client
}

// NewFansMedalRpcClient creates a Rpc client that implements the FansMedal interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewFansMedalRpcClient(client *liverpc.Client) FansMedal {
	return &fansMedalRpcClient{
		client: client,
	}
}

func (c *fansMedalRpcClient) GetWearedMedal(ctx context.Context, in *FansMedalGetWearedMedalReq) (*FansMedalGetWearedMedalResp, error) {
	out := new(FansMedalGetWearedMedalResp)
	err := doRpcRequest(ctx, c.client, 1, "FansMedal.get_weared_medal", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fansMedalRpcClient) TargetsWithMedal(ctx context.Context, in *FansMedalTargetsWithMedalReq) (*FansMedalTargetsWithMedalResp, error) {
	out := new(FansMedalTargetsWithMedalResp)
	err := doRpcRequest(ctx, c.client, 1, "FansMedal.targetsWithMedal", in, out)
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
