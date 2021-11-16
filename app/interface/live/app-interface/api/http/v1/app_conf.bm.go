// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: app_conf.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with github.com/namelessup/bilibili/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	app_conf.proto
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

var PathConfigGetConf = "/live.appinterface.v1.config/getConf"

// ================
// Config Interface
// ================

type ConfigBMServer interface {
	GetConf(ctx context.Context, req *GetConfReq) (resp *GetConfResp, err error)
}

var v1ConfigSvc ConfigBMServer

func configGetConf(c *bm.Context) {
	p := new(GetConfReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1ConfigSvc.GetConf(c, p)
	c.JSON(resp, err)
}

// RegisterV1ConfigService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1ConfigService(e *bm.Engine, svc ConfigBMServer, midMap map[string]bm.HandlerFunc) {
	v1ConfigSvc = svc
	e.GET("/xlive/app-interface/v1/config/getConf", configGetConf)
}

// RegisterConfigBMServer Register the blademaster route
func RegisterConfigBMServer(e *bm.Engine, server ConfigBMServer) {
	v1ConfigSvc = server
	e.GET("/live.appinterface.v1.config/getConf", configGetConf)
}
