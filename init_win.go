//go:build windows

package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var configFile = flag.String("f", "", "阿里云配置文件")

func init() {
	if *configFile == "" {
		path, err := os.Executable()
		if err != nil {
			return
		}
		*configFile = fmt.Sprintf("%v\\%v", filepath.Dir(path), "config.yaml")
		log.Fatal(*configFile)
	}
}
