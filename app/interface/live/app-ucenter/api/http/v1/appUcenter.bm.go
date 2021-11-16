// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: appUcenter.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with github.com/namelessup/bilibili/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	appUcenter.proto
*/
package v1

import (
	"context"

	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathRoomGetInfo = "/live.appucenter.v1.Room/GetInfo"
var PathRoomCreate = "/live.appucenter.v1.Room/Create"

var PathTopicGetTopicList = "/live.appucenter.v1.Topic/GetTopicList"
var PathTopicCheckTopic = "/live.appucenter.v1.Topic/CheckTopic"

// ==============
// Room Interface
// ==============

type RoomBMServer interface {
	// 获取房间基本信息
	// `method:"GET" midware:"auth"`
	GetInfo(ctx context.Context, req *GetRoomInfoReq) (resp *GetRoomInfoResp, err error)

	// 创建房间
	// `method:"POST" midware:"auth"`
	Create(ctx context.Context, req *CreateReq) (resp *CreateResp, err error)
}

var v1RoomSvc RoomBMServer

func roomGetInfo(c *bm.Context) {
	p := new(GetRoomInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomSvc.GetInfo(c, p)
	c.JSON(resp, err)
}

func roomCreate(c *bm.Context) {
	p := new(CreateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomSvc.Create(c, p)
	c.JSON(resp, err)
}

// RegisterV1RoomService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1RoomService(e *bm.Engine, svc RoomBMServer, midMap map[string]bm.HandlerFunc) {
	auth := midMap["auth"]
	v1RoomSvc = svc
	e.GET("/xlive/app-ucenter/v1/room/GetInfo", auth, roomGetInfo)
	e.POST("/xlive/app-ucenter/v1/room/Create", auth, roomCreate)
}

// RegisterRoomBMServer Register the blademaster route
func RegisterRoomBMServer(e *bm.Engine, server RoomBMServer) {
	v1RoomSvc = server
	e.GET("/live.appucenter.v1.Room/GetInfo", roomGetInfo)
	e.POST("/live.appucenter.v1.Room/Create", roomCreate)
}

// ===============
// Topic Interface
// ===============

type TopicBMServer interface {
	// 获取话题列表
	// `method:"GET" midware:"auth"`
	GetTopicList(ctx context.Context, req *GetTopicListReq) (resp *GetTopicListResp, err error)

	// 检验话题是否有效
	// `method:"GET" midware:"auth"`
	CheckTopic(ctx context.Context, req *CheckTopicReq) (resp *CheckTopicResp, err error)
}

var v1TopicSvc TopicBMServer

func topicGetTopicList(c *bm.Context) {
	p := new(GetTopicListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TopicSvc.GetTopicList(c, p)
	c.JSON(resp, err)
}

func topicCheckTopic(c *bm.Context) {
	p := new(CheckTopicReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1TopicSvc.CheckTopic(c, p)
	c.JSON(resp, err)
}

// RegisterV1TopicService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1TopicService(e *bm.Engine, svc TopicBMServer, midMap map[string]bm.HandlerFunc) {
	auth := midMap["auth"]
	v1TopicSvc = svc
	e.GET("/xlive/app-ucenter/v1/topic/GetTopicList", auth, topicGetTopicList)
	e.GET("/xlive/app-ucenter/v1/topic/CheckTopic", auth, topicCheckTopic)
}

// RegisterTopicBMServer Register the blademaster route
func RegisterTopicBMServer(e *bm.Engine, server TopicBMServer) {
	v1TopicSvc = server
	e.GET("/live.appucenter.v1.Topic/GetTopicList", topicGetTopicList)
	e.GET("/live.appucenter.v1.Topic/CheckTopic", topicCheckTopic)
}
