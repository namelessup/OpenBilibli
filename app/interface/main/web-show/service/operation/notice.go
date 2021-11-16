package operation

import (
	"context"

	opmdl "github.com/namelessup/bilibili/app/interface/main/web-show/model/operation"
)

// Notice Service
func (s *Service) Notice(c context.Context, arg *opmdl.ArgOp) (res map[string][]*opmdl.Operation) {
	res = s.operation(arg.Tp, arg.Rank, arg.Count)
	return
}
