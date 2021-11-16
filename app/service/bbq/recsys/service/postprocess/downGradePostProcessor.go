package postprocess

import (
	"context"
	"github.com/namelessup/bilibili/app/service/bbq/recsys/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/bbq/recsys/model"
	"math/rand"
)

//DownGradeProcessor ..
type DownGradeProcessor struct {
	Processor
}

func (p *DownGradeProcessor) name() (name string) {
	name = "DownGradeProcessor"
	return
}

func (p *DownGradeProcessor) process(ctx context.Context, request *v1.RecsysRequest, response *v1.RecsysResponse, u *model.UserProfile) (err error) {

	if _, ok := response.Message[model.ResponseDownGrade]; ok {
		rand.Shuffle(len(response.List), func(i, j int) {
			response.List[i], response.List[j] = response.List[j], response.List[i]
		})
	}
	return
}
