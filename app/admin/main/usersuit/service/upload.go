package service

import (
	"context"
	"io"

	"github.com/namelessup/bilibili/library/log"
)

// Upload http upload file.
func (s *Service) Upload(c context.Context, fileName, fileType string, expire int64, body io.Reader) (location string, err error) {
	if location, err = s.d.Upload(c, fileName, fileType, expire, body); err != nil {
		log.Error("s.upload.Upload() error(%v)", err)
	}
	return
}
