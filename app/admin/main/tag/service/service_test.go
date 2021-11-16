package service

import (
	"flag"
	"path/filepath"
	"testing"

	"github.com/namelessup/bilibili/app/admin/main/tag/conf"
	"github.com/namelessup/bilibili/library/log"
)

var testSvc *Service

func TestMain(m *testing.M) {
	dir, _ := filepath.Abs("../cmd/tag-admin.toml")
	flag.Set("conf", dir)
	conf.Init()
	log.Init(conf.Conf.Log)
	testSvc = New(conf.Conf)
}
