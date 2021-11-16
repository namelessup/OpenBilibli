package service

import (
	"image"
	"io"

	"github.com/namelessup/bilibili/library/log"
)

// Pixel get width height
func Pixel(file io.Reader) (width, height int, err error) {
	var c image.Config
	if c, _, err = image.DecodeConfig(file); err != nil {
		log.Error("decode config error", err)
		return
	}
	return c.Width, c.Height, err
}
