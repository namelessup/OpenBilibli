package app

import (
	"context"
	"github.com/namelessup/bilibili/app/job/main/tv/model/common"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_passedSn = "SELECT id,ctime FROM tv_ep_season WHERE `check` = ? AND valid = ? AND category = ? AND is_deleted = 0"
)

// PassedSn picks the passed seasons data to sync to search
func (d *Dao) PassedSn(c context.Context, category int) (res []*common.IdxRank, err error) {
	var (
		rows *sql.Rows
	)
	if rows, err = d.DB.Query(c, _passedSn, SeasonPassed, _CMSValid, category); err != nil {
		log.Error("d.PassedSn.Query error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r = &common.IdxRank{}
		if err = rows.Scan(&r.ID, &r.Ctime); err != nil {
			log.Error("PassedSn row.Scan() error(%v)", err)
			return
		}
		res = append(res, r)
	}
	if err = rows.Err(); err != nil {
		log.Error("d.PassedSn.Query error(%v)", err)
	}
	return
}
