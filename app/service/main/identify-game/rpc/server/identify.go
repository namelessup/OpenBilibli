package server

import (
	"github.com/namelessup/bilibili/app/service/main/identify-game/model"
	"github.com/namelessup/bilibili/library/net/rpc/context"
)

// DelCache del token cache.
func (r *RPC) DelCache(c context.Context, arg *model.CleanCacheArgs, res *struct{}) (err error) {
	err = r.s.DelCache(c, arg.Token)
	return
}
