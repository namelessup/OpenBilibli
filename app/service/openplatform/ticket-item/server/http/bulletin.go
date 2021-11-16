package http

/**import (
	"github.com/pkg/errors"

	"github.com/namelessup/bilibili/app/service/openplatform/ticket-item/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// @params ParamID
// @router get /openplatform/internal/ticket/item/getbulletins
// @response BulletinInfo
func getBulletins(c *bm.Context) {
	arg := new(model.ParamID)
	if err := c.Bind(arg); err != nil {
		errors.Wrap(err, "参数验证失败")
		return
	}
	c.JSON(itemSvc.GetBulletins(c, &arg.ID))
}**/
