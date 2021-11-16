package v1

import (
	"github.com/namelessup/bilibili/app/interface/bbq/app-bbq/model"
)

// LocationRequest .
type LocationRequest struct {
	PID int32 `form:"pid"`
}

// LocationResponse .
type LocationResponse struct {
	List []*model.Location
}
