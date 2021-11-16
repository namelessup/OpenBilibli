package external

import (
	"github.com/namelessup/bilibili/app/interface/main/app-feed/conf"
	"github.com/namelessup/bilibili/app/interface/main/app-feed/dao/dynamic"
)

// Service .
type Service struct {
	dynamic *dynamic.Dao
}

// New .
func New(c *conf.Config) *Service {
	return &Service{
		dynamic: dynamic.New(c),
	}
}
