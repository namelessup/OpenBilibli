package manager

import (
	"context"

	"github.com/namelessup/bilibili/library/log"
)

const (
	_upsSQL = "SELECT mid,type FROM ups"
)

// Uppers get uppers by type.
func (d *Dao) Uppers(c context.Context) (um map[int8]map[int64]struct{}, err error) {
	rows, err := d.upsStmt.Query(c)
	if err != nil {
		log.Error("d.upsStmt.Query error(%v)", err)
		return
	}
	defer rows.Close()
	um = map[int8]map[int64]struct{}{}
	for rows.Next() {
		var (
			mid int64
			tp  int8
		)
		if err = rows.Scan(&mid, &tp); err != nil {
			log.Error("rows.Scan error(%v)", err)
			return
		}
		if mm, ok := um[tp]; ok {
			mm[mid] = struct{}{}
		} else {
			um[tp] = map[int64]struct{}{mid: struct{}{}}
		}
	}
	return
}
