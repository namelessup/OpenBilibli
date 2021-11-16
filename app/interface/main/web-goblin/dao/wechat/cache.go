package wechat

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web-goblin/model/wechat"
)

//go:generate $GOPATH/src/github.com/namelessup/bilibili/app/tool/cache/gen
type _cache interface {
	// cache
	AccessToken(c context.Context) (*wechat.AccessToken, error)
}
