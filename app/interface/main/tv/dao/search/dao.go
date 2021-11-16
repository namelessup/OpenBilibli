package search

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

const (
	_userSearch = "/main/search"
	_card       = "/pgc/internal/season/search/card"
)

// Dao is search dao.
type Dao struct {
	conf       *conf.Config
	client     *bm.Client
	resultURL  string
	esClient   *elastic.Elastic
	userSearch string
	card       string
	cfgWild    *conf.WildSearch
}

// New account dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		client:    bm.NewClient(c.SearchClient),
		resultURL: c.Search.ResultURL,
		conf:      c,
		esClient: elastic.NewElastic(&elastic.Config{
			Host:       c.Host.ESHost,
			HTTPClient: c.SearchClient,
		}),
		userSearch: c.Search.UserSearch + _userSearch,
		card:       c.Host.APICo + _card,
		cfgWild:    c.Wild.WildSearch,
	}
	return
}
