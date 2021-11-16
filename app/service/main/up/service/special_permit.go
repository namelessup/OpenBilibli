package service

import (
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

//SpecialGroupPermit check permit for special groups
func (s *Service) SpecialGroupPermit(c *blademaster.Context, groupID int64) (err error) {
	if handlerList, ok := s.permCheckMap[groupID]; ok && handlerList != nil {
		for _, f := range handlerList {
			f(c)
			if c.IsAborted() {
				err = ecode.AccessDenied
				break
			}
		}
	}
	return
}
