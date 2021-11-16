package newbiedao

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/growup/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	//"github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"

	"net/url"
)

// const text
const (
	// ActivityTypeVideo type "videoall"
	ActivityTypeVideo = "videoall"
)

// GetActivities get activities
func (d *Dao) GetActivities(c context.Context) (res []*model.Activity, err error) {
	activitiesRes := new(model.ActivitiesRes)
	err = d.httpRead.Get(c, d.c.Host.ActivitiesURI+ActivityTypeVideo, metadata.String(c, metadata.RemoteIP), url.Values{}, &activitiesRes)
	if err != nil {
		log.Error("s.dao.GetActivities error(%v)", err)
		return
	}
	if activitiesRes.Code != ecode.OK.Code() {
		err = ecode.Int(activitiesRes.Code)
		log.Error("s.dao.GetActivities get activities failed, ecode: %d", activitiesRes.Code)
		return
	}

	res = activitiesRes.Data
	return
}
