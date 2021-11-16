package global

import (
	"runtime"

	accgrpc "github.com/namelessup/bilibili/app/service/main/account/api"
	arcgrpc "github.com/namelessup/bilibili/app/service/main/archive/api"
	"github.com/namelessup/bilibili/app/service/main/up/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"

	"github.com/pkg/errors"
)

var (
	upCrmDB *sql.DB
	worker  *fanout.Fanout
	arcCli  arcgrpc.ArchiveClient
	accCli  accgrpc.AccountClient
)

// GetArcClient .
func GetArcClient() arcgrpc.ArchiveClient {
	return arcCli
}

// GetAccClient .
func GetAccClient() accgrpc.AccountClient {
	return accCli
}

// GetWorker .
func GetWorker() *fanout.Fanout {
	return worker
}

// GetUpCrmDB .
func GetUpCrmDB() *sql.DB {
	return upCrmDB
}

//Init init global
func Init(c *conf.Config) {
	var err error
	if arcCli, err = arcgrpc.NewClient(c.GRPCClient.Archive); err != nil {
		panic(errors.WithMessage(err, "Failed to dial archive service"))
	}
	if accCli, err = accgrpc.NewClient(c.GRPCClient.Account); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	}
	upCrmDB = sql.NewMySQL(c.DB.UpCRM)
	worker = fanout.New("cache", fanout.Worker(runtime.NumCPU()), fanout.Buffer(1024))
}

// Close .
func Close() {
	upCrmDB.Close()
	worker.Close()
}
