package service

import (
	"testing"

	"github.com/namelessup/bilibili/app/service/main/antispam/conf"
)

func TestNewOption(t *testing.T) {
	NewOption(&conf.Config{ServiceOption: &conf.ServiceOption{GcOpt: &conf.GcOpt{}}})
}
