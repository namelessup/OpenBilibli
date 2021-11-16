package music

import (
	"context"
	"fmt"
	"net/url"

	"github.com/namelessup/bilibili/app/interface/main/favorite/conf"
	"github.com/namelessup/bilibili/app/interface/main/favorite/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"github.com/namelessup/bilibili/library/xstr"
)

const _music = "http://api.bilibili.co/x/internal/v1/audio/songs/batch"

// Dao defeine fav Dao
type Dao struct {
	httpClient *httpx.Client
}

// New return fav dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		httpClient: httpx.NewClient(c.HTTPClient),
	}
	return
}

// MusicMap return the music map data(all state).
func (d *Dao) MusicMap(c context.Context, musicIds []int64) (data map[int64]*model.Music, err error) {
	params := url.Values{}
	params.Set("level", "1")
	params.Set("ids", xstr.JoinInts(musicIds))
	res := new(model.MusicResult)
	ip := metadata.String(c, metadata.RemoteIP)
	if err = d.httpClient.Get(c, _music, ip, params, res); err != nil {
		log.Error("d.HTTPClient.Get(%s?%s) error(%v)", _music, params.Encode())
		return
	}
	if res.Code != ecode.OK.Code() {
		log.Error("d.HTTPClient.Get(%s?%s) code:%d msg:%s", _music, params.Encode(), res.Code)
		err = fmt.Errorf("Get Music failed!code:=%v", res.Code)
		return
	}
	return res.Data, nil
}
