package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/videoup-report/model/archive"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_tpsSQL = "SELECT id,pid,name FROM archive_type"
)

// TypeMapping is second types opposite first types.
func (d *Dao) TypeMapping(c context.Context) (rmap map[int16]*archive.Type, err error) {
	rows, err := d.db.Query(c, _tpsSQL)
	if err != nil {
		log.Error("d.tpsStmt.Query error(%v)", err)
		return
	}
	defer rows.Close()
	rmap = map[int16]*archive.Type{}
	for rows.Next() {
		tp := &archive.Type{}
		if err = rows.Scan(&tp.ID, &tp.PID, &tp.Name); err != nil {
			log.Error("rows.Scan error(%v)", err)
			return
		}
		rmap[tp.ID] = tp
	}
	return
}
