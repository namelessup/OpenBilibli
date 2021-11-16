package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/upload/model"
	"github.com/namelessup/bilibili/library/log"
)

// AddDir .
func (s *Service) AddDir(c context.Context, adp *model.AddDirParam) (err error) {
	d := &model.DirLimit{}
	if err = s.orm.Model(d).
		Where(&model.DirLimit{BucketName: adp.BucketName, Dir: adp.DirName}).
		Assign(&model.DirLimit{
			BucketName: adp.BucketName,
			Dir:        adp.DirName,
			ConfigPic:  adp.Pic,
			ConfigRate: adp.Rate,
		}).
		FirstOrCreate(d).Error; err != nil {
		log.Error("Failed to add dir (%+v): %v", d, err)
		return
	}
	return
}
