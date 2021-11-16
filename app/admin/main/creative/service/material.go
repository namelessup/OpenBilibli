package service

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/creative/model/material"
)

// CategoryByID .
func (s *Service) CategoryByID(c context.Context, id int64) (cate *material.Category, err error) {
	return s.dao.CategoryByID(c, id)
}

// BindWithCategory .
func (s *Service) BindWithCategory(c context.Context, MaterialID, CategoryID, index int64) (id int64, err error) {
	return s.dao.BindWithCategory(c, MaterialID, CategoryID, index)
}
