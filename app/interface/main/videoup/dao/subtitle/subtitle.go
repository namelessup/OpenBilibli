package subtitle

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/dm2/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// Update fn
func (d *Dao) Update(c context.Context, aid int64, open bool, lan string) (err error) {
	var arg = &model.ArgSubtitleAllowSubmit{
		Aid:         aid,
		AllowSubmit: open,
		Lan:         lan,
	}
	if err = d.sub.SubtitleSujectSubmit(c, arg); err != nil {
		log.Error("d.sub.SubtitleSujectSubmit (%+v) error(%v)", arg, err)
		err = ecode.CreativeSubtitleAPIErr
	}
	return
}
