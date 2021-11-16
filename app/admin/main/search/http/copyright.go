package http

import (
	"encoding/json"
	"github.com/namelessup/bilibili/app/admin/main/search/dao"
	"github.com/namelessup/bilibili/app/admin/main/search/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func copyRight(c *bm.Context) {
	var (
		err      error
		bulkItem []dao.BulkItem
		d        []*model.CopyRight
		form     = c.Request.Form
	)
	data := form.Get("data")
	if data == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if err = json.Unmarshal([]byte(data), &d); err != nil {
		log.Error("json.Unmarshal error(%v)", err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	for _, v := range d {
		bulkItem = append(bulkItem, v)
	}
	if err = svr.Index(c, "internalPublic", bulkItem); err != nil {
		log.Error("srv.Index error(%v)", err)
	}
	c.JSON(nil, err)
}
