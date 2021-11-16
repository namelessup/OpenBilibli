package http

import (
	"encoding/json"
	"io/ioutil"

	"github.com/namelessup/bilibili/app/admin/main/macross/model/publish"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// dashboard get user.
func dashboard(c *bm.Context) {
	req := c.Request
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("ioutil.ReadAll() error(%v)", err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	req.Body.Close()
	// params
	var d = &publish.Dashboard{}
	if err = json.Unmarshal(bs, d); err != nil {
		log.Error("http dashboard() json.Unmarshal(%s) error(%v)", string(bs), err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if d.Name == "" || d.Label == "" || d.Commit == "" || d.TextSizeArm64 == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, svr.Dashborad(c, d))
}
