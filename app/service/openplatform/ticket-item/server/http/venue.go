package http

import (
	item "github.com/namelessup/bilibili/app/service/openplatform/ticket-item/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-item/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/pkg/errors"
)

// @params VenueSearchParam
// @router get /openplatform/internal/ticket/item/venue/search
// @response VenueSearchList
func venueSearch(c *bm.Context) {
	req := &model.VenueSearchParam{}
	if err := c.Bind(req); err != nil {
		return
	}
	c.JSON(itemSvc.VenueSearch(c, req))
}

// @params VenueInfoParam
// @router post /openplatform/internal/ticket/item/venueInfo
// @response VenueInfoReply
func venueInfo(c *bm.Context) {
	arg := new(model.VenueInfoParam)
	if err := c.Bind(arg); err != nil {
		errors.Wrap(err, "参数验证失败")
		return
	}
	c.JSON(itemSvc.VenueInfo(c, &item.VenueInfoRequest{
		ID:            arg.ID,
		Name:          arg.Name,
		Status:        arg.Status,
		Province:      arg.Province,
		City:          arg.City,
		District:      arg.District,
		AddressDetail: arg.AddressDetail,
		Traffic:       arg.Traffic,
	}))
}
