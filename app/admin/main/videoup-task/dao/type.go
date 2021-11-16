package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/videoup-task/model"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_tpsSQL = "SELECT id,pid,name,description FROM archive_type"
)

// TypeMapping is second types opposite first types.
func (d *Dao) TypeMapping(c context.Context) (tmap map[int16]*model.Type, err error) {
	rows, err := d.arcDB.Query(c, _tpsSQL)
	if err != nil {
		log.Error("d.arcDB.Query error(%v)", err)
		return
	}
	defer rows.Close()
	tmap = make(map[int16]*model.Type)
	for rows.Next() {
		t := &model.Type{}
		if err = rows.Scan(&t.ID, &t.PID, &t.Name, &t.Desc); err != nil {
			log.Error("rows.Scan error(%v)", err)
			return
		}
		tmap[t.ID] = t
	}
	return
}
