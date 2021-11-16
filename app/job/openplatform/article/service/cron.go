package service

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/job/openplatform/article/dao"
	"github.com/namelessup/bilibili/library/log"
)

func (s *Service) updateSortproc() {
	for {
		if err := s.UpdateSort(context.TODO()); err != nil {
			log.Error("s.UpdateSort error(%+v)", err)
			dao.PromError("service:刷新分区投稿")
		} else {
			dao.PromInfo("service:刷新分区投稿")
		}
		time.Sleep(s.updateSortInterval)
	}
}
