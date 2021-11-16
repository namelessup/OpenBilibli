package service

import (
	"github.com/namelessup/bilibili/app/admin/main/videoup/model/archive"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// TxUpPorder .
func (s *Service) TxUpPorder(tx *sql.Tx, ap *archive.ArcParam) (err error) {
	//区分自首还是审核回查添加
	if _, err = s.arc.TxUpPorder(tx, ap.Aid, ap); err != nil {
		log.Error("s.arc.TxUpPorder(%d,%+v) error(%v)", ap.Aid, ap, err)
		return
	}
	log.Info("TxUpPorder aid(%d) update archive_porder", ap.Aid)
	return
}
