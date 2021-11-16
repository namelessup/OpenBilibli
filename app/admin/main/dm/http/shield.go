package http

import (
	"bytes"
	"io"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func shieldUpload(c *bm.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Error("shieldUpload.file.illegal,err::%v", err)
		c.JSON(nil, ecode.FileNotExists)
		return
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, file); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, dmSvc.DmShield(c, buf.Bytes()))
}
