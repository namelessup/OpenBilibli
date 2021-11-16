package weeklyhonor

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/archive/model/archive"
	upgrpc "github.com/namelessup/bilibili/app/service/main/up/api/v1"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

const fromWeeklyHonor = 1

// UpCount get archives count.
func (d *Dao) UpCount(c context.Context, mid int64) (count int, err error) {
	var arg = &archive.ArgUpCount2{Mid: mid}
	if count, err = d.arc.UpCount2(c, arg); err != nil {
		log.Error("rpc UpCount2 (%v) error(%v)", mid, err)
		err = ecode.CreativeArcServiceErr
	}
	return
}

// UpActivesList list up-actives
func (d *Dao) UpActivesList(c context.Context, lastID int64, ps int) (upActives []*upgrpc.UpActivity, newid int64, err error) {
	upListReq := upgrpc.UpListByLastIDReq{
		LastID: lastID,
		Ps:     ps,
	}
	reply, err := d.upClient.UpInfoActivitys(c, &upListReq)
	if err != nil {
		log.Error("failed to list up&active info,err(%v)", err)
		return
	}
	newid = reply.GetLastID()
	upActives = reply.GetUpActivitys()
	return
}

// GetUpSwitch get up switch state
func (d *Dao) GetUpSwitch(c context.Context, mid int64) (state uint8, err error) {
	req := upgrpc.UpSwitchReq{
		Mid:  mid,
		From: fromWeeklyHonor,
	}
	reply, err := d.upClient.UpSwitch(c, &req)
	if err != nil {
		log.Error("d.upClient.UpSwitch req(%+v),err(%v)", req, err)
		return
	}
	state = reply.GetState()
	return
}
