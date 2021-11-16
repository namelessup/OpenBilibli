package cache

import (
	"github.com/namelessup/bilibili/app/service/ops/log-agent/event"
)

type Cahce interface {
	WriteToCache(e *event.ProcessorEvent)
	ReadFromCache() (e *event.ProcessorEvent)
}
