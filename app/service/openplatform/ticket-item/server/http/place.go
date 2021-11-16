package http

import (
	item "github.com/namelessup/bilibili/app/service/openplatform/ticket-item/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-item/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/pkg/errors"
)

// @params PlaceInfoParam
// @router post /openplatform/internal/ticket/item/placeInfo
// @response PlaceInfoReply
func placeInfo(c *bm.Context) {
	arg := new(model.PlaceInfoParam)
	if err := c.Bind(arg); err != nil {
		errors.Wrap(err, "参数验证失败")
		return
	}
	c.JSON(itemSvc.PlaceInfo(c, &item.PlaceInfoRequest{
		ID:      arg.ID,
		Status:  arg.Status,
		Name:    arg.Name,
		BasePic: arg.BasePic,
		Venue:   arg.Venue,
		DWidth:  arg.DWidth,
		DHeight: arg.DHeight,
	}))
}
