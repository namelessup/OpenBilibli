package http

import (
	item "github.com/namelessup/bilibili/app/service/openplatform/ticket-item/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-item/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/pkg/errors"
)

// @params AreaInfoParam
// @router post /openplatform/internal/ticket/item/areaInfo
// @response AreaInfoReply
func areaInfo(c *bm.Context) {
	arg := new(model.AreaInfoParam)
	if err := c.Bind(arg); err != nil {
		errors.Wrap(err, "参数验证失败")
		return
	}
	c.JSON(itemSvc.AreaInfo(c, &item.AreaInfoRequest{
		ID:         arg.ID,
		AID:        arg.AID,
		Name:       arg.Name,
		Place:      arg.Place,
		Coordinate: arg.Coordinate,
	}))
}
