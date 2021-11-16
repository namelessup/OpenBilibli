package http

import (
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func segment(c *bm.Context) {
	var params = c.Request.Form
	idStr := params.Get("id")
	content := params.Get("content")
	withTagStr := params.Get("with_tag")

	id, _ := strconv.Atoi(idStr)
	withTag, _ := strconv.Atoi(withTagStr)
	if id == 0 || content == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(artSrv.Segment(c, int32(id), content, withTag, "draft"))
}
