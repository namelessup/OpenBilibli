package conf

import (
	xtime "github.com/namelessup/bilibili/library/time"
)

// JPushConfig 极光推送配置
type JPushConfig struct {
	AppKey         string
	SecretKey      string
	Timeout        xtime.Duration
	ApnsProduction bool
}
