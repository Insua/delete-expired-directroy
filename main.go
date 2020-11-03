package main

import (
	"github.com/gogf/gf/os/gfile"
	"time"
)

func main() {
	path := gfile.RealPath(".")
	dirs, err := gfile.ScanDir(path, "*")
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
