package dao

import (
	"context"
	"github.com/namelessup/bilibili/app/service/bbq/video/model"
	"github.com/namelessup/bilibili/library/log"
)

const (
	keyOnBoard = "OnboardVideo"
)

// CmsPub pub cms data into databus.
func (d *Dao) CmsPub(c context.Context, data *model.DataTopicCmsData) (err error) {
	if err = d.cmsPub.Send(c, keyOnBoard, data); err != nil {
		log.Error("d.databus.Send error(%v)", err)
	}
	return
}
