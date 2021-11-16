package service

import "github.com/namelessup/bilibili/app/admin/ep/melloi/model"

//AddClientMoni add ClientMoni
func (s *Service) AddClientMoni(clm *model.ClientMoni) (int, error) {
	return s.dao.AddClientMoni(clm)
}
