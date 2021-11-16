package dao

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/namelessup/bilibili/app/service/main/passport-auth/model"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_getCookieSessionSQL = "SELECT mid,session,csrf,type,expires FROM user_cookie_%s where session = ? limit 1"
)

// Cookie get cookie by session
func (d *Dao) Cookie(c context.Context, sd []byte, ct time.Time) (res *model.Cookie, session []byte, err error) {
	row := d.db.QueryRow(c, fmt.Sprintf(_getCookieSessionSQL, formatSuffix(ct)), sd)
	res = new(model.Cookie)
	var csrf []byte
	if err = row.Scan(&res.Mid, &session, &csrf, &res.Type, &res.Expires); err != nil {
		if err == xsql.ErrNoRows {
			res = nil
			err = nil
		} else {
			log.Error("row.Scan() error(%v)", err)
		}
		return
	}
	res.CSRF = hex.EncodeToString(csrf)
	return
}
