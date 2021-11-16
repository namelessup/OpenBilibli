package dao

import (
	"context"

	"github.com/namelessup/bilibili/library/database/bfs"
	"github.com/namelessup/bilibili/library/log"
)

// Upload .
func (d *Dao) Upload(c context.Context, bucket, fileName, contentType string, file []byte) (err error) {
	if _, err = d.bfsCli.Upload(c, &bfs.Request{
		Bucket:      bucket,
		ContentType: contentType,
		Filename:    fileName,
		File:        file,
	}); err != nil {
		log.Error("Upload(err:%v)", err)
	}
	return
}
