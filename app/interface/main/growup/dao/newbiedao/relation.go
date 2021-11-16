package newbiedao

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/growup/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
	"github.com/namelessup/bilibili/library/xstr"
	"net/url"
	"strconv"
)

// GetRelations get relations
func (d *Dao) GetRelations(c context.Context, mid int64, fids []int64) (res map[int64]*model.Relation, err error) {
	relationsRes := new(model.RelationsRes)

	uv := url.Values{}
	uv.Set("mid", strconv.FormatInt(mid, 10))
	uv.Set("fids", xstr.JoinInts(fids))

	err = d.httpRead.Get(c, d.c.Host.RelationsURI, metadata.String(c, metadata.RemoteIP), uv, relationsRes)
	if err != nil {
		log.Error("s.dao.GetRelations error(%v)", err)
		return
	}
	if relationsRes.Code != ecode.OK.Code() {
		err = ecode.Int(relationsRes.Code)
		log.Error("s.dao.GetRelations get relations failed, ecode: %d", relationsRes.Code)
		return
	}

	res = relationsRes.Data
	return
}
