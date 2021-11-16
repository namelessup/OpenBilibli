package model

import "github.com/namelessup/bilibili/library/time"

// SignUp table sign_up
type SignUp struct {
	ID        int64
	Mid       int64
	State     int
	BeginDate time.Time
	EndDate   time.Time
}
