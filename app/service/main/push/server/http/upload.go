package http

import (
	"path/filepath"

	"github.com/namelessup/bilibili/app/service/main/push/conf"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var imgExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func upimg(ctx *bm.Context) {
	f, h, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Error("upimg error(%v)", err)
		ctx.JSON(nil, err)
		return
	}
	defer f.Close()
	if h.Size > conf.Conf.Push.UpimgMaxSize {
		log.Error("filesize error name(%s) size(%d)", h.Filename, h.Size)
		ctx.JSON(nil, ecode.PushServiceFileSizeErr)
		return
	}
	if ok := imgExts[filepath.Ext(h.Filename)]; !ok {
		log.Error("file ext error name(%s)", h.Filename)
		ctx.JSON(nil, ecode.PushServiceFileExtErr)
		return
	}
	url, err := pushSrv.Upimg(ctx, f)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(map[string]string{"url": url}, nil)
}
