package game

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao  define
type Dao struct {
	c           *conf.Config
	client      *httpx.Client
	gameListURL string
	gameInfoURL string
}

// New init dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:           c,
		client:      httpx.NewClient(c.HTTPClient.Slow),
		gameListURL: c.Game.OpenHost + _gameListURI,
		gameInfoURL: c.Game.OpenHost + _gameInfoURI,
	}
	return
}
