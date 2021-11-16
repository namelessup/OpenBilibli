package bfs

import (
	"bytes"
	"context"

	"github.com/namelessup/bilibili/app/admin/main/app/conf"
	bfsdao "github.com/namelessup/bilibili/app/admin/main/app/dao/bfs"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// Service bfs service.
type Service struct {
	dao        *bfsdao.Dao
	bfsMaxSize int
}

// New new a bfs service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao:        bfsdao.New(c),
		bfsMaxSize: c.Bfs.MaxFileSize,
	}
	return
}

// ClientUpCover client upload cover.
func (s *Service) ClientUpCover(c context.Context, fileType string, body []byte) (url string, err error) {
	if len(body) == 0 {
		err = ecode.FileNotExists
		return
	}
	if len(body) > s.bfsMaxSize {
		err = ecode.FileTooLarge
		return
	}
	url, err = s.dao.Upload(c, fileType, bytes.NewReader(body))
	if err != nil {
		log.Error("s.bfs.Upload error(%v)", err)
	}
	return
}
