package service

import (
	"context"
	"testing"

	"github.com/namelessup/bilibili/app/infra/config/model"
	"github.com/namelessup/bilibili/library/log"
)

func TestConfig(t *testing.T) {
	arg := &model.ArgConf{
		App:      "zjx_test",
		BuildVer: "1_0_0_0",
		Ver:      62,
		Env:      "3",
	}
	if err := svr.confSvr.Push(context.TODO(), arg); err != nil {
		log.Error("push(%v) error(%v)", arg, err)
		t.FailNow()
	}
	argT := &model.ArgToken{
		App:   "zjx_test",
		Env:   "3",
		Token: "123",
	}
	if err := svr.confSvr.SetToken(context.TODO(), argT); err != nil {
		log.Error("push(%v) error(%v)", argT, err)
		t.FailNow()
	}

}
