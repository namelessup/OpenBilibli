package service

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/apm/model/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// GetCodes ...
func (s *Service) GetCodes(c context.Context, Interval1, Interval2 string) (data []*codes.Codes, err error) {
	data, err = s.dao.GetCodes(c, Interval1, Interval2)
	if err != nil {
		log.Error("service GetCodes error(%v)", err)
	}
	return
}
