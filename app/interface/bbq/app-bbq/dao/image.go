package dao

import (
	"context"
	image "github.com/namelessup/bilibili/app/service/bbq/video-image/api/grpc/v1"
	"github.com/namelessup/bilibili/library/log"
)

//Upload .
func (d *Dao) Upload(c context.Context, fileName string, filePath string, file []byte) (location string, err error) {
	imageReq := &image.ImgUploadRequest{
		Filename: fileName,
		Dir:      filePath,
		File:     file,
	}
	imageRes, err := d.imageClient.ImgUpload(c, imageReq)
	if err != nil {
		log.Errorv(c, log.KV("event", "grpc/imageupload"), log.KV("err", err))
		return
	}
	if imageRes != nil {
		location = imageRes.Location
	}
	return
}
