package collect

import (
	"github.com/namelessup/bilibili/app/service/main/dapper/pkg/process"
)

// Collecter collect span from different source
type Collecter interface {
	Start() error
	RegisterProcess(p process.Processer)
	Close() error
}
