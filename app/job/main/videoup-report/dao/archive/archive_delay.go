package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/videoup-report/model/archive"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_delaysSQL = "SELECT id,aid,dtime,type,state FROM archive_delay WHERE aid=? AND deleted_at = 0 ORDER BY dtime DESC LIMIT 1"
)

// Delay get delay by aid
func (d *Dao) Delay(c context.Context, aid int64) (delay *archive.Delay, err error) {
	rows := d.db.QueryRow(c, _delaysSQL, aid)
	delay = &archive.Delay{}
	if err = rows.Scan(&delay.ID, &delay.Aid, &delay.DTime, &delay.Type, &delay.State); err != nil {
		if err == sql.ErrNoRows {
			delay = nil
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}
