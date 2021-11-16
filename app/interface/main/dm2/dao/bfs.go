package dao

import (
	"context"

	"github.com/namelessup/bilibili/library/database/bfs"
	"github.com/namelessup/bilibili/library/log"
)

// UploadBfs .
func (d *Dao) UploadBfs(c context.Context, fileName string, bs []byte) (location string, err error) {
	if location, err = d.bfsCli.Upload(c, &bfs.Request{
		Bucket:      d.conf.Bfs.BucketSubtitle,
		Filename:    fileName,
		ContentType: "application/json",
		File:        bs,
	}); err != nil {
		log.Error("bfs.BfsDmUpload error(%v)", err)
		return
	}
	return
}
