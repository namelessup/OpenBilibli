package http

import (
	"fmt"

	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/trace"
)

// WrapHeader 为返回头添加自定义字段
func WrapHeader(ctx *bm.Context) {
	// Traceid
	tracer, _ := trace.FromContext(ctx.Context)
	traceid := fmt.Sprintf("%s", tracer)
	ctx.Writer.Header().Set("traceid", traceid)

	// Sessionid
	sid := ctx.Request.Header.Get("SessionID")
	if sid == "" {
		sid = SessionID(ctx)
	}
	ctx.Set("SessionID", sid)
	ctx.Writer.Header().Set("SessionID", sid)
}
