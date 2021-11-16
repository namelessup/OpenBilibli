package service

import (
	"context"
	"encoding/json"

	"github.com/namelessup/bilibili/app/job/live-userexp/model"
	"github.com/namelessup/bilibili/library/log"
)

func (s *Service) levelCacheUpdate(nwMsg []byte, oldMsg []byte) {
	exp := &model.Exp{}
	if err := json.Unmarshal(nwMsg, exp); err != nil {
		log.Error("json.Unmarshal(%s) error(%v)", string(nwMsg), err)
		return
	}
	level := model.FormatLevel(exp)
	s.dao.SetLevelCache(context.TODO(), level)
}
