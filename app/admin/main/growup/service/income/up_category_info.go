package income

import (
	"context"

	model "github.com/namelessup/bilibili/app/admin/main/growup/model/income"
	"github.com/namelessup/bilibili/library/log"
)

// GetUpInfoByAIDs get up_info by av id
func (s *Service) GetUpInfoByAIDs(c context.Context, avs []*model.ArchiveIncome) (upInfoMap map[int64]string, err error) {
	midMap := make(map[int64]struct{})
	for _, av := range avs {
		midMap[av.MID] = struct{}{}
	}
	midList := []int64{}
	for mid := range midMap {
		midList = append(midList, mid)
	}
	upInfoMap = make(map[int64]string)
	if len(midList) > 0 {
		upInfoMap, err = s.dao.ListUpInfo(c, midList)
		if err != nil {
			log.Error("s.dao.ListUpInfo error(%v)", err)
			return
		}
	}
	return
}
