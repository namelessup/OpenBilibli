package service

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/workflow/model"
	"github.com/namelessup/bilibili/library/log"
)

// AddEvent add event
func (s *Service) AddEvent(c context.Context, cid int32, content, attachments string, event int8) (row int32, err error) {
	et := &model.Event{Cid: cid, Event: event, Content: content, Attachments: attachments}
	if err = s.dao.DB.Create(et).Error; err != nil {
		log.Error("s.workflow.AddEvent(%+v) error(%v)", et, err)
		return
	}
	row = et.ID
	return
}
