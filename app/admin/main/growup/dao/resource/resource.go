package resource

import (
	"github.com/namelessup/bilibili/app/admin/main/growup/conf"
	accgrpc "github.com/namelessup/bilibili/app/service/main/account/api"
	vip "github.com/namelessup/bilibili/app/service/main/vip/rpc/client"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/pkg/errors"
)

var (
	vipRPC             *vip.Service
	client             *httpx.Client
	accCli             accgrpc.AccountClient
	videoCategoryURL   string
	articleCategoryURL string
)

// Init .
func Init(c *conf.Config) {
	var err error
	vipRPC = vip.New(c.VipRPC)
	client = httpx.NewClient(c.HTTPClient)
	videoCategoryURL = c.Host.VideoType + "/videoup/types"
	articleCategoryURL = c.Host.ColumnType + "/x/article/categories"
	if accCli, err = accgrpc.NewClient(c.Account); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	}
}
