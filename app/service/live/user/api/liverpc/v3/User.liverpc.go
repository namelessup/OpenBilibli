// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v3/User.proto

/*
Package v3 is a generated liverpc stub package.
This code was generated with github.com/namelessup/bilibili/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v3/User.proto
*/
package v3

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/namelessup/bilibili/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ==============
// User Interface
// ==============

type User interface {
	// * uid获取房间信息
	//
	GetMultiple(context.Context, *UserGetMultipleReq) (*UserGetMultipleResp, error)

	// * uid获取房间信息
	//
	GetUserLevelInfo(context.Context, *UserGetUserLevelInfoReq) (*UserGetUserLevelInfoResp, error)
}

// ====================
// User Live Rpc Client
// ====================

type userRpcClient struct {
	client *liverpc.Client
}

// NewUserRpcClient creates a Rpc client that implements the User interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewUserRpcClient(client *liverpc.Client) User {
	return &userRpcClient{
		client: client,
	}
}

func (c *userRpcClient) GetMultiple(ctx context.Context, in *UserGetMultipleReq) (*UserGetMultipleResp, error) {
	out := new(UserGetMultipleResp)
	err := doRpcRequest(ctx, c.client, 3, "User.getMultiple", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) GetUserLevelInfo(ctx context.Context, in *UserGetUserLevelInfoReq) (*UserGetUserLevelInfoResp, error) {
	out := new(UserGetUserLevelInfoResp)
	err := doRpcRequest(ctx, c.client, 3, "User.getUserLevelInfo", in, out)
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
