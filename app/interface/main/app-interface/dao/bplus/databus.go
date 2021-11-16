package bplus

import (
	"context"
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/model/space"
	"github.com/namelessup/bilibili/library/net/metadata"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/pkg/errors"
)

// NotifyContribute .
func (d *Dao) NotifyContribute(c context.Context, vmid int64, attrs *space.Attrs, ctime xtime.Time) (err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	value := struct {
		Vmid  int64        `json:"vmid"`
		Attrs *space.Attrs `json:"attrs"`
		CTime xtime.Time   `json:"ctime"`
		IP    string       `json:"ip"`
	}{vmid, attrs, ctime, ip}
	if err = d.pub.Send(c, strconv.FormatInt(vmid, 10), value); err != nil {
		err = errors.Wrapf(err, "%v", value)
	}
	return
}
