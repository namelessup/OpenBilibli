package server

import (
	"github.com/namelessup/bilibili/app/service/main/usersuit/model"
	"github.com/namelessup/bilibili/library/net/rpc/context"
)

// PointFlag obtain new pendant noify.
func (r *RPC) PointFlag(c context.Context, arg *model.ArgMID, res *model.PointFlag) (err error) {
	var pf *model.PointFlag
	if pf, err = r.s.PointFlag(c, arg); err == nil && pf != nil {
		*res = *pf
	}
	return
}
