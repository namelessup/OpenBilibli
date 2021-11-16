package videoup

import (
	"github.com/namelessup/bilibili/app/admin/main/mcn/conf"
	"github.com/namelessup/bilibili/app/service/main/videoup/model/archive"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

const (
	_typeURL = "/videoup/types"
)

// Dao .
type Dao struct {
	c                *conf.Config
	client           *bm.Client
	videTypeURL      string
	videoUpTypeCache map[int]archive.Type
}

// New new a Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// http client
		client:           bm.NewClient(c.HTTPClient),
		videTypeURL:      c.Host.Videoup + _typeURL,
		videoUpTypeCache: make(map[int]archive.Type),
	}
	d.refreshUpType()
	go d.refreshUpTypeAsync()
	return
}
