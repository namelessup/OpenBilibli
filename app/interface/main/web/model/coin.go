package model

import (
	v1 "github.com/namelessup/bilibili/app/service/main/archive/api"
)

// CoinArc coin archive.
type CoinArc struct {
	*v1.Arc
	Coins int64  `json:"coins"`
	Time  int64  `json:"time"`
	IP    string `json:"ip"`
}
