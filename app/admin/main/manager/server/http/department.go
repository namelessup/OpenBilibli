package http

import (
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// departments .
func departments(c *bm.Context) {
	c.JSON(mngSvc.Departments(c))
}

// roles .
func roles(c *bm.Context) {
	c.JSON(mngSvc.Roles(c))
}

// userByDepartment .
func usersByDepartment(c *bm.Context) {
	ID, _ := strconv.ParseInt(c.Request.Form.Get("id"), 10, 64)
	if ID <= 0 {
		c.JSON(nil, ecode.RequestErr)
		log.Error("ID unnarmal (%d)", ID)
		return
	}
	c.JSON(mngSvc.UsersByDepartment(c, ID))
}

// userByRole .
func usersByRole(c *bm.Context) {
	ID, _ := strconv.ParseInt(c.Request.Form.Get("id"), 10, 64)
	if ID <= 0 {
		c.JSON(nil, ecode.RequestErr)
		log.Error("ID unnarmal (%d)", ID)
		return
	}
	c.JSON(mngSvc.UsersByRole(c, ID))
}
