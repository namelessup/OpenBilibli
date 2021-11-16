package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/passport-user-compare/model"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

var (
	secretSQL = "SELECT us.key_type, us.key FROM user_secret us"
)

// LoadSecret load secret
func (d *Dao) LoadSecret(c context.Context) (res []*model.Secret, err error) {
	var rows *xsql.Rows
	if rows, err = d.secretDB.Query(c, secretSQL); err != nil {
		log.Error("fail to get secretSQL, dao.secretDB.Query(%s) error(%v)", secretSQL, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.Secret)
		if err = rows.Scan(&r.KeyType, &r.Key); err != nil {
			log.Error("row.Scan() error(%v)", err)
			res = nil
			return
		}
		res = append(res, r)
	}
	return
}
