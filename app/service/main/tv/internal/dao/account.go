package dao

import (
	"context"
	"github.com/namelessup/bilibili/app/service/main/account/api"
	"github.com/namelessup/bilibili/library/log"
)

// AccountInfo queries account info by user id.
func (d *Dao) AccountInfo(c context.Context, mid int64) (ai *api.Info, err error) {
	req := &api.MidReq{Mid: int64(mid)}
	res, err := d.accCli.Info3(c, req)
	if err != nil {
		log.Error("d.AccountInfo(%d) err(%v)", mid, err)
		return
	}
	log.Info("d.AccountInfo(%d) res(%+v)", mid, res.Info)
	return res.Info, nil
}
