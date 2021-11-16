package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/archive/api"
	"github.com/namelessup/bilibili/app/service/main/archive/model/archive"
	"github.com/namelessup/bilibili/library/log"
)

// Stat get a archive stat.
func (d *Dao) Stat(c context.Context, aid int64) (st *api.Stat, err error) {
	if st, err = d.statCache(c, aid); err != nil {
		log.Error("%+v", err)
	} else if st != nil {
		return
	}
	arg := &archive.ArgAid2{Aid: aid}
	if st, err = d.arcRPC.Stat3(c, arg); err != nil {
		log.Error("d.arcRPC.Stat3(%v) error(%v)", arg, err)
	}
	return
}
