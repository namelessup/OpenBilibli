package dao

import (
	"context"

	arc "github.com/namelessup/bilibili/app/service/main/archive/api"
	"github.com/namelessup/bilibili/library/ecode"
)

// ArchiveUGCPay get archive ugcpay flag.
func (d *Dao) ArchiveUGCPay(ctx context.Context, aid int64) (pay bool, err error) {
	var (
		req = &arc.ArcRequest{
			Aid: aid,
		}
		reply *arc.ArcReply
	)
	if reply, err = d.archiveAPI.Arc(ctx, req); err != nil {
		if err == ecode.NothingFound {
			err = nil
			pay = false
			return
		}
		return
	}
	if reply != nil && reply.Arc != nil && reply.Arc.Rights.UGCPay == 1 {
		pay = true
	} else {
		pay = false
	}
	return
}
