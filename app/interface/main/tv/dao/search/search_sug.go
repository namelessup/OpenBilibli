package search

import (
	"context"
	"fmt"
	"net/url"

	searchMdl "github.com/namelessup/bilibili/app/interface/main/tv/model/search"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"

	"github.com/pkg/errors"
)

// SearchSug gets the search sug detail data from Search API
func (d *Dao) SearchSug(ctx context.Context, req *searchMdl.ReqSug) (result searchMdl.SugResponse, err error) {
	var (
		config = d.conf.Search
		params = url.Values{}
	)
	// common params
	params.Set("main_ver", config.MainVer)
	params.Set("sug_num", fmt.Sprintf("%d", config.SugNum))
	params.Set("suggest_type", config.SugType)
	params.Set("highlight", config.Highlight)
	params.Set("build", req.Build)
	params.Set("mobi_app", req.MobiApp)
	params.Set("platform", req.Platform)
	params.Set("term", req.Term) // search term
	if err = d.client.Get(ctx, config.URL, "", params, &result); err != nil {
		log.Error("ClientGet URL %s error[%v]", config.URL, err)
		return
	}
	if result.Code != ecode.OK.Code() {
		log.Error("ClientGet Code Result Not OK [%v]", result)
		err = errors.Wrap(ecode.ServerErr, "Search API Error")
		return
	}
	return
}
