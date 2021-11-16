package service

import (
	"github.com/namelessup/bilibili/library/log"
)

//TransToCheckBack ..
func (s *Service) TransToCheckBack() {
	log.Info("deliveryNewVdieoToCms begin")
	s.dao.TransToCheckBack()
}

//TransToReview ...
func (s *Service) TransToReview() {
	log.Info("TransToReview begin")
	s.dao.TransToReview()
}
