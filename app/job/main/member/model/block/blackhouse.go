package block

import (
	xtime "github.com/namelessup/bilibili/library/time"
)

// CreditAnswerMSG param struct
type CreditAnswerMSG struct {
	MID   int64      `json:"mid"`
	MTime xtime.Time `json:"mtime"`
}
