package http

import (
	"testing"

	"github.com/namelessup/bilibili/app/admin/main/mcn/model"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func TestValidater(t *testing.T) {
	var req = model.MCNSignEntryReq{
		MCNMID:    1,
		BeginDate: "2018-01-01",
		EndDate:   "2018-01-01",
	}
	var err = binding.Validator.ValidateStruct(&req)
	if err == nil {
		t.FailNow()
	} else {
		t.Logf("err=%s", err)
	}
}
