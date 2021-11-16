package http

import (
	"bytes"
	"io"

	"github.com/namelessup/bilibili/app/admin/main/upload/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

// InternalUploadAdminImage .
func InternalUploadAdminImage(c *bm.Context) {
	var err error
	up := new(model.UploadParam)
	if err = c.BindWith(up, binding.FormMultipart); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	up.WMInit()
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Error("upload.UploadImage.file.illegal,err::%v", err.Error())
		c.JSON(nil, ecode.RequestErr)
		return
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, file); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(uaSvc.UploadAdminRecord(c, "internal_admin_upload", up, buf.Bytes()))
}
