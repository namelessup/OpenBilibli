package tag

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/tag/model"
	"github.com/namelessup/bilibili/library/log"
)

//CheckChannelReview check whether archive in channel review list
func (d *Dao) CheckChannelReview(c context.Context, aids []int64) (response map[int64]*model.ResChannelCheckBack, err error) {
	arg := &model.ArgResChannel{
		Oids: aids,
		Type: 3,
	}

	if response, err = d.tagRPC.ResChannelCheckBack(c, arg); err != nil {
		log.Error("CheckChannelReview d.tagDisRPC.ResChannelCheckBack error(%v) aids(%+v)", err, aids)
	}
	return
}
