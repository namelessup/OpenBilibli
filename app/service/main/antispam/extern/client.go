package extern

import (
	"github.com/namelessup/bilibili/app/service/main/antispam/conf"

	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

type Client struct {
	*ReplyServiceClient
}

func NewClient(c *conf.Config) *Client {
	httpCli := bm.NewClient(c.HTTPClient)

	return &Client{
		ReplyServiceClient: &ReplyServiceClient{
			host: c.ReplyURL,
			commonClient: &commonClient{
				httpCli: httpCli,
				key:     c.App.Key,
				secret:  c.App.Secret,
			},
		},
	}
}

type commonClient struct {
	httpCli     *bm.Client
	key, secret string
}
