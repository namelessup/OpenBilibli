package http

import (
	"net/http"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/app/model/audit"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/render"
)

func audits(c *bm.Context) {
	c.JSON(auditSvc.Audits(c))
}

func auditByID(c *bm.Context) {
	v := &audit.Param{}
	if err := c.Bind(v); err != nil {
		return
	}
	c.JSON(auditSvc.AuditByID(c, v.ID))
}

func auditSave(c *bm.Context) {
	var (
		err error
		v   = &audit.Param{}
	)
	if err = c.Bind(v); err != nil {
		return
	}
	if v.Build == 0 || v.MobiApp == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if v.ID > 0 {
		err = auditSvc.UpdateAudit(c, v, time.Now())
	} else {
		err = auditSvc.AddAudit(c, v, time.Now())
	}
	switch err {
	case ecode.NotModified:
		data := map[string]interface{}{
			"code":    err,
			"message": "该build号和mobi_app已存在",
		}
		c.Render(http.StatusOK, render.MapJSON(data))
	default:
		c.JSON(nil, err)
	}
}

func auditDelByIDs(c *bm.Context) {
	v := &audit.Param{}
	if err := c.Bind(v); err != nil {
		return
	}
	if v.ID == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, auditSvc.DelAudit(c, v.ID))
}
