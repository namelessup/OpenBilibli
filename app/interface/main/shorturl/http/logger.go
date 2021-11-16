package http

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"github.com/namelessup/bilibili/library/stat"
)

// logger is logger  middleware
func logger() bm.HandlerFunc {
	const noUser = "no_user"
	return func(c *bm.Context) {
		now := time.Now()
		ip := metadata.String(c, metadata.RemoteIP)
		req := c.Request
		path := req.URL.Path
		params := req.Form

		c.Next()

		mid, _ := c.Get("mid")
		userI, _ := c.Get("user")
		err := c.Error
		cerr := ecode.Cause(err)
		dt := time.Since(now)
		// user
		user, ok := userI.(string)
		if !ok || user == "" {
			user = noUser
		}

		realPath := ""
		if strings.HasPrefix(path, "/x/internal/shorturl") {
			realPath = path[1:]
		} else {
			realPath = "shorturl"
		}

		stat.HTTPServer.Incr(user, realPath, strconv.FormatInt(int64(cerr.Code()), 10))
		stat.HTTPServer.Timing(user, int64(dt/time.Millisecond), realPath)

		lf := log.Infov
		errmsg := ""
		if err != nil {
			errmsg = err.Error()
			lf = log.Errorv
		}
		lf(c,
			log.KV("method", req.Method),
			log.KV("mid", mid),
			log.KV("ip", ip),
			log.KV("user", user),
			log.KV("path", path),
			log.KV("params", params.Encode()),
			log.KV("ret", cerr.Code()),
			log.KV("msg", cerr.Message()),
			log.KV("stack", fmt.Sprintf("%+v", err)),
			log.KV("err", errmsg),
		)
	}
}
