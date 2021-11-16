package dao

import (
	"context"
	"github.com/namelessup/bilibili/app/service/bbq/recsys/model"
	"github.com/namelessup/bilibili/app/service/bbq/user/api"
	"github.com/namelessup/bilibili/library/log"
)

//GetUserFollow ...
func (d *Dao) GetUserFollow(c context.Context, mid int64, u *model.UserProfile) (err error) {

	if mid == 0 {
		return
	}

	relationReq := &api.ListRelationReq{Mid: mid}
	listRelationReply, err := d.UserClient.ListFollow(c, relationReq)
	if err != nil {
		log.Errorv(c)
		return
	}
	for _, MID := range listRelationReply.List {
		u.BBQFollow[MID] = 1
	}
	return
}

//GetUserBlack ...
func (d *Dao) GetUserBlack(c context.Context, mid int64, u *model.UserProfile) (err error) {

	if mid == 0 {
		return
	}

	relationReq := &api.ListRelationReq{Mid: mid}
	listRelationReply, err := d.UserClient.ListBlack(c, relationReq)
	if err != nil {
		log.Errorv(c)
		return
	}
	for _, MID := range listRelationReply.List {
		u.BBQBlack[MID] = 1
	}
	return
}
