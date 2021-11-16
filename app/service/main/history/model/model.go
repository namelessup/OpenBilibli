package model

//go:generate $GOPATH/src/github.com/namelessup/bilibili/app/tool/warden/protoc.sh

const (
	// HideStateON 不记录播放历史
	HideStateON = 1
	// HideStateOFF 记录播放历史
	HideStateOFF = 0
	// HideStateNotFound not found
	HideStateNotFound = -1
)
