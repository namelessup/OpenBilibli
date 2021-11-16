// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Banner.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/namelessup/bilibili/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ================
// Banner Interface
// ================

type Banner interface {
	// * 获取新后台配置的banner
	//
	GetNewBanner(context.Context, *BannerGetNewBannerReq) (*BannerGetNewBannerResp, error)
}

// ======================
// Banner Live Rpc Client
// ======================

type bannerRpcClient struct {
	client *liverpc.Client
}

// NewBannerRpcClient creates a Rpc client that implements the Banner interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewBannerRpcClient(client *liverpc.Client) Banner {
	return &bannerRpcClient{
		client: client,
	}
}

func (c *bannerRpcClient) GetNewBanner(ctx context.Context, in *BannerGetNewBannerReq) (*BannerGetNewBannerResp, error) {
	out := new(BannerGetNewBannerResp)
	err := doRpcRequest(ctx, c.client, 1, "Banner.getNewBanner", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
