package archive

import (
	"context"
	"time"

	model "github.com/namelessup/bilibili/app/interface/main/creative/model/archive"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_bizsByTimeSQL = "SELECT aid,type,ctime FROM archive_biz WHERE mtime >= ? AND mtime < ? AND type = ? ORDER BY mtime"
)

// BIZsByTime list businesses by time and type
func (d *Dao) BIZsByTime(c context.Context, start, end *time.Time, tp int8) (bizs []*model.BIZ, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(c, _bizsByTimeSQL, start, end, tp); err != nil {
		log.Error("d.db.Query error(%v)", err)
		return
	}
	defer rows.Close()
	bizs = []*model.BIZ{}
	for rows.Next() {
		var b = new(model.BIZ)
		if err = rows.Scan(&b.Aid, &b.Type, &b.Ctime); err != nil {
			log.Error("row.Scan error(%v)", err)
			return
		}
		bizs = append(bizs, b)
	}
	return
}
