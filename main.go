//go:build darwin || linux || windows

package main

import (
	"aliDDNS/aliddns"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
)
import "github.com/sirupsen/logrus"

var qs = new(aliddns.QueryStruct)
var isTest = flag.Bool("t", false, "测试用例，开启测试")
var configFile = flag.String("f", "/etc/ddns/config.yaml", "阿里云配置文件")
var renewType = flag.String("rt", "ipv6", "设置检查类型，支持ipv4,cert,ipv6三种")
var Key = flag.String("k", os.Getenv("CERTBOT_VALIDATION"), "自动获取系统certbot参数，也可以进行传参")
var log = logrus.New()

// var remoteIpv4 = ""
var dnsHeader = "_acme-challenge."
var account Account

func main() {
	flag.Parse()
	if flag.Parsed() {
		qs.IsTest = *isTest
		//fmt.Println(*Key, *renewType)
		conf.MustLoad(*configFile, &account)

		qs.AccessKeyId = account.AccessKeyId
		qs.AccessSecret = account.AccessSecret
		qs.MainDomain = account.MainDomain
		qs.SubDomain = account.SubDomain
		switch *renewType {
		case "ipv4":
			qs.ValueType = "A"
			qs.GetOutBoundIPV4()
		case "ipv6":
			qs.ValueType = "AAAA"
			qs.GetOutBoundIPV6()
		case "cert":
			qs.ValueType = "TXT"
			qs.SubDomain = dnsHeader + qs.SubDomain
			qs.Value = *Key
		default:
			return
		}
		qs.DnsCheck()
	}
}
