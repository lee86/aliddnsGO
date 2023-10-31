//go:build darwin || linux || windows

package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
)
import "github.com/sirupsen/logrus"

var qs = new(QueryStruct)
var isTest = flag.Bool("t", false, "测试")
var configFile = flag.String("f", "/etc/ddns/config.yaml", "the config file")
var renewType = flag.String("rt", "ipv6", "设置certbot更新txt或者设置更新ipv6")
var Key = flag.String("k", os.Getenv("CERTBOT_VALIDATION"), "")
var log = logrus.New()

// var remoteIpv4 = ""
var dnsHeader = "_acme-challenge."
var account Account

func main() {
	flag.Parse()
	if flag.Parsed() {
		//fmt.Println(*Key, *renewType)
		conf.MustLoad(*configFile, &account)

		qs.AccessKeyId = account.AccessKeyId
		qs.AccessSecret = account.AccessSecret
		qs.MainDomain = account.MainDomain
		qs.SubDomain = account.SubDomain
		switch *renewType {
		case "ipv6":
			Ipv6DDNS()
		case "cert":
			CertCheck()
		default:
			return
		}
	}
}
