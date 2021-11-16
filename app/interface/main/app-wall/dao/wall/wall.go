package wall

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-wall/conf"
	"github.com/namelessup/bilibili/app/interface/main/app-wall/model/wall"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

const (
	//wall
	_wallSQL = "SELECT logo,download,name,package,size,remark FROM wall WHERE state=1 ORDER BY rank DESC"
)

type Dao struct {
	db      *sql.DB
	wallSQL *sql.Stmt
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: sql.NewMySQL(c.MySQL.Show),
	}
	d.wallSQL = d.db.Prepared(_wallSQL)
	return
}

func (d *Dao) WallAll(ctx context.Context) (res []*wall.Wall, err error) {
	rows, err := d.wallSQL.Query(ctx)
	if err != nil {
		log.Error("query error (%v)", err)
		return
	}
	defer rows.Close()
	res = []*wall.Wall{}
	for rows.Next() {
		a := &wall.Wall{}
		if err = rows.Scan(&a.Logo, &a.Download, &a.Name, &a.Package, &a.Size, &a.Remark); err != nil {
			log.Error("rows.Scan err (%v)", err)
			return
		}
		res = append(res, a)
	}
	return
}

func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.Ping(c)
}
