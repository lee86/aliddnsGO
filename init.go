//go:build linux || darwin

package main

import "flag"

var configFile = flag.String("f", "/etc/ddns/config.yaml", "阿里云配置文件")
