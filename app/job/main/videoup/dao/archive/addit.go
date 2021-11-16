package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/videoup/model/archive"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_additSQL = "SELECT id,aid,source,redirect_url,mission_id,up_from,order_id,dynamic FROM archive_addit WHERE aid=?"
)

// Addit get archive addit.
func (d *Dao) Addit(c context.Context, aid int64) (addit *archive.Addit, err error) {
	row := d.db.QueryRow(c, _additSQL, aid)
	addit = &archive.Addit{}
	if err = row.Scan(&addit.ID, &addit.Aid, &addit.Source, &addit.RedirectURL, &addit.MissionID, &addit.UpFrom, &addit.OrderID, &addit.Dynamic); err != nil {
		if err == sql.ErrNoRows {
			addit = nil
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}
