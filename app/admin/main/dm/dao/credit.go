package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/namelessup/bilibili/app/admin/main/dm/model"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_sendJudge = "/x/internal/credit/blocked/case/add"
)

// SendJudgement send to judgement
func (d *Dao) SendJudgement(c context.Context, judges []*model.ReportJudge) (err error) {
	params := url.Values{}
	ret := struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	}{}
	data, err := json.Marshal(judges)
	if err != nil {
		log.Error("send judgement params(%s) create error(%v)", data, err)
		return
	}
	params.Set("data", string(data))
	if err = d.httpCli.Post(c, d.sendJudgeURI, "", params, &ret); err != nil {
		log.Error("send judgement request(data: %s) error(%v)", data, err)
		return
	}
	if ret.Code != 0 {
		err = fmt.Errorf("%v", ret)
		log.Error("send judgement request(data: %s) error(%v)", data, err)
	}
	return
}
