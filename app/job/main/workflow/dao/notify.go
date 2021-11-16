package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/workflow/model/param"
	"github.com/namelessup/bilibili/app/job/main/workflow/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

// SendMessage .
func (d *Dao) SendMessage(c context.Context, chs []*model.ChallRes, msg *param.MessageParam) (err error) {
	params := msg.Query()
	var res struct {
		Code int `json:"code"`
	}
	if err = d.httpSearch.Post(c, d.messageURL, "", params, &res); err != nil {
		return
	}
	if res.Code != 0 {
		err = ecode.Int(res.Code)
		return
	}
	for _, ch := range chs {
		report.Manager(&report.ManagerInfo{
			UID:      2233,
			Uname:    "",
			Business: 11,
			Type:     2,
			Oid:      ch.OID,
			Ctime:    time.Now(),
			Action:   "notify_users_received",
			Content: map[string]interface{}{
				"mid":     ch.MID,
				"message": "zhoushuguang",
			},
			Index: []interface{}{ch.ID, ch.GID, ch.MID},
		})
	}
	return
}
