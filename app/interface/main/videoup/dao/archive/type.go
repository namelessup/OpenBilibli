package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/videoup/model/archive"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_typesURL = "/videoup/types"
)

// TypeMapping is second types opposite first types.
func (d *Dao) TypeMapping(c context.Context) (rmap map[int16]*archive.Type, err error) {
	var res struct {
		Code    int                     `json:"code"`
		Message string                  `json:"message"`
		Data    map[int16]*archive.Type `json:"data"`
	}
	if err = d.httpR.Get(c, d.typesURI, "", nil, &res); err != nil {
		log.Error("videoup view archive error(%v) | typesURI(%s)", err, d.typesURI)
		err = ecode.CreativeArchiveAPIErr
		return
	}
	if res.Code != 0 {
		err = ecode.CreativeArchiveAPIErr
		log.Error("get archive type failed res.Code(%d) | typesURI(%s) res(%v)", res.Code, d.typesURI, res)
		return
	}
	rmap = make(map[int16]*archive.Type, len(res.Data))
	for _, v := range res.Data {
		if v.PID != 0 {
			rmap[v.ID] = v
		}
	}
	return
}
