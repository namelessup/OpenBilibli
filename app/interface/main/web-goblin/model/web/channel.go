package web

import (
	tagmdl "github.com/namelessup/bilibili/app/interface/main/tag/model"
	"github.com/namelessup/bilibili/app/service/main/archive/api"
)

// ChCard channel card .
type ChCard struct {
	ID         int64  `json:"-"`
	Title      string `json:"-"`
	ChannelID  int64  `json:"-"`
	Type       string `json:"-"`
	Value      int64  `json:"-"`
	Reason     string `json:"-"`
	ReasonType int8   `json:"-"`
	Pos        int    `json:"-"`
	FromType   string `json:"-"`
}

// Channel .
type Channel struct {
	*tagmdl.Tag
	Archives []*api.Arc `json:"archives"`
}
