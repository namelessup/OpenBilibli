package archive

import (
	"context"
	"database/sql"

	"github.com/namelessup/bilibili/app/job/main/archive/model/archive"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_arcPassedOperSQL = "SELECT id FROM archive_oper WHERE aid=? AND state>=? LIMIT 1"
)

// PassedOper check archive passed
func (d *Dao) PassedOper(c context.Context, aid int64) (id int64, err error) {
	row := d.db.QueryRow(c, _arcPassedOperSQL, aid, archive.StateOpen)
	if err = row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}
