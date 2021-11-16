package model

import xtime "github.com/namelessup/bilibili/library/time"

// UserContract represents user contract record.
type UserContract struct {
	ID         int32
	Mid        int64
	ContractId string
	OrderNo    string
	IsDeleted  int8
	Ctime      xtime.Time
	Mtime      xtime.Time
}
