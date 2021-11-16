package income

import (
	"github.com/namelessup/bilibili/library/time"
)

// Signed signed up
type Signed struct {
	MID          int64
	AccountState int
	SignedAt     time.Time
	QuitAt       time.Time
	IsDeleted    int
}
