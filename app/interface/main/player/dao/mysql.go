package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/player/model"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_param = "SELECT `name`,`value` FROM `param` WHERE `plat` = 9 AND `state` = 0"
)

// Param return player setting params.
func (d *Dao) Param(c context.Context) (param []*model.Param, err error) {
	var (
		rows *sql.Rows
		pa   *model.Param
	)
	if rows, err = d.paramStmt.Query(c); err != nil {
		log.Error("d.paramStmt.Query() error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		pa = &model.Param{}
		if err = rows.Scan(&pa.Name, &pa.Value); err != nil {
			log.Error("rows.Scan() error(%v)", err)
			return
		}
		param = append(param, pa)
	}
	return
}
