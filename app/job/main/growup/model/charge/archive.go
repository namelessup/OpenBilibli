package charge

import (
	"github.com/namelessup/bilibili/library/time"
)

// Archive archive detail
type Archive struct {
	ID        int64
	IncCharge int64
	TagID     int64
	Date      time.Time
}
