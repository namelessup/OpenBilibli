package service

import (
	"github.com/namelessup/bilibili/app/job/main/block/conf"
)

// MSGRemoveInfo get msg info
func (s *Service) MSGRemoveInfo() (code string, title, content string) {
	code = conf.Conf.Property.MSG.BlockRemove.Code
	title = conf.Conf.Property.MSG.BlockRemove.Title
	content = conf.Conf.Property.MSG.BlockRemove.Content
	return
}
