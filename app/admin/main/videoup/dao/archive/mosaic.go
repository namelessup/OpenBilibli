package archive

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/videoup/model/archive"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const _mosaic = "SELECT id, aid, cid, coordinate,ctime FROM archive_video_mosaic WHERE cid=?"

//Mosaic get mosaic
func (d *Dao) Mosaic(c context.Context, cid int64) (m []*archive.Mosaic, err error) {
	var rows *sql.Rows
	m = []*archive.Mosaic{}
	if rows, err = d.rddb.Query(c, _mosaic, cid); err != nil {
		log.Error("Mosaic d.rddb.Query error(%v) cid(%d)", err, cid)
		return
	}
	defer rows.Close()

	for rows.Next() {
		ms := new(archive.Mosaic)
		if err = rows.Scan(&ms.ID, &ms.AID, &ms.CID, &ms.Coordinate, &ms.CTime); err != nil {
			log.Error("Mosaic rows.Scan error(%v) cid(%d)", err, cid)
			return
		}

		m = append(m, ms)
	}
	return
}
