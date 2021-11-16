package medal

import (
	"context"
	"github.com/namelessup/bilibili/app/service/main/usersuit/model"
)

// Grant sent a medal to user.
func (d *Dao) Grant(c context.Context, mid, nid int64) (err error) {
	err = d.suitRPC.MedalGrant(c, &model.ArgMIDNID{MID: mid, NID: nid})
	return
}
