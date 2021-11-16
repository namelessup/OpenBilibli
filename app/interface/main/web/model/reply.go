package model

import (
	"encoding/json"

	rplmdl "github.com/namelessup/bilibili/app/interface/main/reply/model/reply"
)

// ReplyHot reply hot
type ReplyHot struct {
	Page    json.RawMessage `json:"page"`
	Replies []*rplmdl.Reply `json:"replies"`
}
