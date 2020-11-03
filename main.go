package main

import (
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"time"
)

func main() {
	path := gcmd.GetArg(1)
	if len(path) == 0 {
		panic("please input path")
	}
	realPath := gfile.RealPath(path)
	dirs, err := gfile.ScanDir(realPath, "*")
	if err != nil {
		panic(err)
	}
	for _, v := range dirs {
		if gfile.IsDir(v) {
			info, err := gfile.Info(v)
			if err != nil {
				return
			}
			if time.Now().Sub(info.ModTime()).Hours() > 72 {
				gfile.Remove(v)
			}
		}
	}
}
