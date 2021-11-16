package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/workflow/model"
	"github.com/namelessup/bilibili/library/log"
)

// BusinessAttr .
func (d *Dao) BusinessAttr(c context.Context) (res []*model.BusinessAttr, err error) {
	if err = d.ReadORM.Table("workflow_business_attr").Select("id, bid, name, deal_type, expire_time, assign_type, assign_max, group_type").Find(&res).Error; err != nil {
		log.Error("d.BusinessAttr error(%v)", err)
	}
	return
}
