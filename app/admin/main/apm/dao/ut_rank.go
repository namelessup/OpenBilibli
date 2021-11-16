package dao

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/namelessup/bilibili/app/admin/main/apm/conf"
	"github.com/namelessup/bilibili/app/admin/main/apm/model/ut"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_gitUsersAPI = "http://git.bilibili.co/api/v4/users"
)

// GitLabFace  return face of gitlab.
func (d *Dao) GitLabFace(c context.Context, username string) (avatarURL string, err error) {
	params := url.Values{}
	params.Set("username", username)
	params.Set("private_token", conf.Conf.Gitlab.Token)
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, _gitUsersAPI, strings.NewReader(params.Encode())); err != nil {
		log.Error("http.NewRequest(%s) error(%v)", username, err)
		return
	}
	res := make([]*ut.Image, 0)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err = d.client.Do(c, req, &res); err != nil {
		log.Error("d.client.Do(%s) error(%v)", username, err)
		return
	}
	for _, v := range res {
		avatarURL = v.AvatarURL
	}
	return
}
