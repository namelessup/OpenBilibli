package global

import (
	"github.com/namelessup/bilibili/app/admin/main/up-rating/conf"
	accrpc "github.com/namelessup/bilibili/app/service/main/account/rpc/client"
)

var (
	accRPC *accrpc.Service3
)

// Init resources
func Init(c *conf.Config) {
	accRPC = accrpc.New3(c.RPCClient.Account)
}

// GetAccRPC .
func GetAccRPC() *accrpc.Service3 {
	return accRPC
}
