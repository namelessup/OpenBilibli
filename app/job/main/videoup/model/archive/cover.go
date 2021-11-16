package archive

import xtime "github.com/namelessup/bilibili/library/time"

// Cover str
type Cover struct {
	Filename string
	TotalNum int
	IndexNum int
	NFSPath  string
	BFSPath  string
	Used     int64
	CTime    xtime.Time
	MTime    xtime.Time
}
