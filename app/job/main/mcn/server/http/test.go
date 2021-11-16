package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func run(c *bm.Context) {
	srv.UpMcnSignStateCron()
	srv.UpMcnUpStateCron()
	srv.UpExpirePayCron()
	//srv.UpMcnDataSummaryCron()
	srv.McnRecommendCron()
	srv.DealFailRecommendCron()
	srv.CheckDateDueCron()
	c.JSON("job is run", nil)
}
