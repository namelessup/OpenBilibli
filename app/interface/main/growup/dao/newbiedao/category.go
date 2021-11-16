package newbiedao

import (
	"context"
	"github.com/pkg/errors"

	"github.com/namelessup/bilibili/app/interface/main/growup/model"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"

	"net/url"
)

// GetCategories get categories
func (d *Dao) GetCategories(c context.Context) (err error) {
	categoriesRes := new(model.CategoriesRes)
	err = d.httpRead.Get(c, d.c.Host.CategoriesURI, metadata.String(c, metadata.RemoteIP), url.Values{}, categoriesRes)
	if err != nil {
		log.Error("s.dao.GetCategories error(%v)", err)
		return
	}
	if categoriesRes.Code != ecode.OK.Code() || len(categoriesRes.Data) <= 0 {
		err = errors.Wrap(ecode.Int(categoriesRes.Code), "get categories failed")
		log.Error("s.dao.GetCategories failed, ecode: %d", categoriesRes.Code)
		return
	}

	Categories = categoriesRes.Data
	return
}
