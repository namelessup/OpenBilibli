package dm

import (
	"github.com/namelessup/bilibili/app/job/main/activity/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao struct.
type Dao struct {
	// http
	broadcastURL string
	httpCli      *bm.Client
}

// New return dm dao instance.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		broadcastURL: "http://api.bilibili.co/x/internal/chat/push/room",
		httpCli:      bm.NewClient(c.HTTPClient),
	}
	return
}
