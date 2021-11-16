package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-tag/conf"
	"github.com/namelessup/bilibili/app/service/main/archive/api"
	arcrpc "github.com/namelessup/bilibili/app/service/main/archive/api/gorpc"
	"github.com/namelessup/bilibili/app/service/main/archive/model/archive"
	"github.com/namelessup/bilibili/library/log"
)

var (
	_emptyArchives = map[int64]*api.Arc{}
)

// Dao is archive dao.
type Dao struct {
	// rpc
	arcRpc *arcrpc.Service2
}

// New new a archive dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// rpc
		arcRpc: arcrpc.New2(c.ArchiveRPC),
	}
	return
}

// Archives multi get archives.
func (d *Dao) Archives(ctx context.Context, aids []int64) (as map[int64]*api.Arc, err error) {
	arg := &archive.ArgAids2{Aids: aids}
	if as, err = d.arcRpc.Archives3(ctx, arg); err != nil {
		log.Error("d.arcRpc.Archives2(%v) error(%v)", arg, err)
		return
	}
	return
}
