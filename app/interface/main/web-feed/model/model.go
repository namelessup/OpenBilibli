package model

import (
	accv1 "github.com/namelessup/bilibili/app/service/main/account/api"
	feedmdl "github.com/namelessup/bilibili/app/service/main/feed/model"
)

// Feed feed
type Feed struct {
	*feedmdl.Feed
	OfficialVerify *accv1.OfficialInfo `json:"official_verify,omitempty"`
}
