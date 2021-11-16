package model

import (
	arcMDL "github.com/namelessup/bilibili/app/service/main/archive/model/archive"
)

// ArcVisible .
func ArcVisible(state int32) bool {
	return state == arcMDL.StateOpen || state == arcMDL.StateOrange || state == arcMDL.StateForbidFixed
}
