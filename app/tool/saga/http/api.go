package http

import (
	"github.com/namelessup/bilibili/app/tool/saga/model"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func buildContributors(c *bm.Context) {
	var (
		err  error
		repo = &model.RepoInfo{}
	)

	if err = c.BindWith(repo, binding.JSON); err != nil {
		log.Error("BindWith error(%v)", err)
		return
	}
	c.JSON(nil, svc.HandleBuildContributors(c, repo))
}
