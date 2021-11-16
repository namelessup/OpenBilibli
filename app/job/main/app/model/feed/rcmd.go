package feed

import "github.com/namelessup/bilibili/app/service/main/archive/api"

type RcmdItem struct {
	ID      int64    `json:"id,omitempty"`
	Tid     int64    `json:"tid,omitempty"`
	Archive *api.Arc `json:"archive,omitempty"`
	Tag     *Tag     `json:"tag,omitempty"`
}
