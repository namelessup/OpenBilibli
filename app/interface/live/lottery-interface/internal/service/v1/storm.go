package v1

import (
	"github.com/namelessup/bilibili/app/interface/live/lottery-interface/internal/service"
	xlottery "github.com/namelessup/bilibili/app/service/live/xlottery/api/grpc/v1"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"net/http"
	"time"

	"github.com/namelessup/bilibili/library/net/http/blademaster/render"
)

//StormJoin  StormJoin
func StormJoin(c *bm.Context) {
	p := new(xlottery.JoinStormReq)
	if err := c.Bind(p); err != nil {
		c.JSON(nil, err)
		return
	}
	// if p.GetRoomid() == 0 {
	// 	c.Render(http.StatusOK, render.MapJSON{
	// 		"code": 400,
	// 		"msg":  "参数错误",
	// 	})
	// 	return
	// }
	mid, _ := c.Get("mid")
	platform := c.Request.URL.Query().Get("platform")
	if platform == "" {
		platform = "web"
	}
	if p.Platform == "" {
		p.Platform = platform
	}
	log.Info("mid:%s platform:%s", mid, platform)
	p.Mid = mid.(int64)

	//反作弊打点
	lancer(c, p)
	resp, err := service.ServiceInstance.StormClient.Join(c, p)
	if err != nil {
		c.Render(http.StatusOK, render.MapJSON{
			"code": 400,
			"msg":  "没抢到",
			"data": []int{},
		})
		return
	}

	switch resp.GetCode() {
	case 1005002:
		c.Render(http.StatusOK, render.MapJSON{
			"code": 429,
			"msg":  resp.GetMsg(),
		})
		return
	case 1005001:
		c.Render(http.StatusOK, render.MapJSON{
			"code": 401,
			"msg":  resp.GetMsg(),
		})
		return
	case 1005003, 1005004, 1005005, 1005016:
		c.Render(http.StatusOK, render.MapJSON{
			"code": 400,
			"msg":  resp.GetMsg(),
		})
		return
	case 0:
		c.JSON(resp.GetJoin(), nil)
		return
	default:
		c.Render(http.StatusOK, render.MapJSON{
			"code": int(resp.GetCode()),
			"msg":  resp.GetMsg(),
		})

	}

}

// StormCheck   StormCheck 检查
func StormCheck(c *bm.Context) {
	p := new(xlottery.CheckStormReq)
	if err := c.Bind(p); err != nil {
		c.JSON(nil, err)
		return
	}
	mid, isexit := c.Get("mid")
	if isexit {
		p.Uid = mid.(int64)
		log.Info("StormCheck uid = %s", mid)
	}
	resp, err := service.ServiceInstance.StormClient.Check(c, p)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	if resp.GetCheck() != nil {
		c.JSON(resp.GetCheck(), nil)
		return
	}
	c.Render(http.StatusOK, render.MapJSON{
		"code": 0,
		"msg":  "",
		"data": []int{},
	})

}

var lacnertype = "rhythmic"

func lancer(c *bm.Context, req *xlottery.JoinStormReq) {
	uid := req.GetMid()
	roomid := req.GetRoomid()
	lotteryid := req.GetId()
	action := "join"
	ip := metadata.String(c, metadata.RemoteIP)
	ts := time.Now().Unix()
	platform := req.GetPlatform()
	clientver := c.Request.URL.Query().Get("version")
	buvid := c.Request.Header.Get("Buvid")
	ua := c.Request.Header.Get("User-Agent")
	referer := c.Request.Header.Get("Referer")
	cookie := getCookie(c.Request)
	abnoormal := 0
	requesturi := c.Request.RequestURI

	err := service.ServiceInstance.Infoc.Info(
		uid,
		roomid,
		lacnertype,
		lotteryid,
		action,
		ip,
		ts,
		platform,
		clientver,
		buvid,
		ua,
		referer,
		cookie,
		abnoormal,
		requesturi)
	if err != nil {
		if err != infoc.ErrFull {
			log.Error("Infoc.Info err: %s", err.Error())
		}
	}
}
func getCookie(req *http.Request) string {
	var cookie string
	for _, v := range req.Cookies() {
		if v.Name != "SESSDATA" {
			cookie = cookie + v.Name + "=" + v.Value + "; "
		}
	}
	return cookie
}
