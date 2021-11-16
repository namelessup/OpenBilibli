package aids

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/app/conf"
	"github.com/namelessup/bilibili/app/admin/main/app/model/aids"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/log"

	"github.com/jinzhu/gorm"
)

const (
	_insertSQL = `INSERT INTO audit_aids(aid) VALUES (?)`
)

// Dao aids dao
type Dao struct {
	db *gorm.DB
}

// New new a aids db
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: orm.NewMySQL(c.ORM.Show),
	}
	d.initORM()
	return
}

// initORM init
func (d *Dao) initORM() {
	d.db.LogMode(true)
}

// Insert insert
func (d *Dao) Insert(ctx context.Context, a *aids.Param) (err error) {
	if err = d.db.Exec(_insertSQL, a.Aid).Error; err != nil {
		log.Error("d.db.Exec err(%v)", err)
		return
	}
	return
}

// Close close connect
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
