package service

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/figure/model"
	spym "github.com/namelessup/bilibili/app/service/main/spy/model"
)

// PutSpyScore handle user spy score chenage message
func (s *Service) PutSpyScore(c context.Context, sc *spym.ScoreChange) (err error) {
	s.figureDao.PutSpyScore(c, sc.Mid, sc.Score)
	if sc.Reason == spym.CoinReason {
		if sc.RiskLevel == spym.CoinHighRisk {
			s.figureDao.PutCoinUnusual(c, sc.Mid, model.ACColumnHighRisk)
		} else {
			s.figureDao.PutCoinUnusual(c, sc.Mid, model.ACColumnLowRisk)
		}
	}
	return
}
