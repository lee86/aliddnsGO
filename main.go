//go:build darwin || linux || windows

package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"go-aliddns/aliddns"
	"os"
)

var qs = new(aliddns.QueryStruct)
var isTest = flag.Bool("t", false, "测试用例，开启测试")
var renewType = flag.String("rt", "ipv6", "设置检查类型，支持ipv4,cert,ipv6三种")
var Key = flag.String("k", os.Getenv("CERTBOT_VALIDATION"), "自动获取系统certbot参数，也可以进行传参")

var dnsHeader = "_acme-challenge."
var account Account

func main() {
	flag.Parse()
	if flag.Parsed() {
		qs.IsTest = *isTest

		conf.MustLoad(*configFile, &account)

		qs.AccessKeyId = account.AccessKeyId
		qs.AccessSecret = account.AccessSecret
		qs.MainDomain = account.MainDomain
		qs.SubDomain = account.SubDomain
		switch *renewType {
		case "ipv4":
			qs.ValueType = "A"
			qs.GetOutBoundIP()
		case "ipv6":
			qs.ValueType = "AAAA"
			qs.GetOutBoundIP()
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
