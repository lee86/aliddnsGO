//go:build darwin || linux || windows

package main

import (
	"flag"
	"fmt"
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

func Ipv6DDNS() {
	localIpv6, err := GetOutBoundIPV6()
	if err != nil {
		return
	}
	qs.Ipv6 = localIpv6

	qs.CreateClient()
	checkError(err)
	qs.queryDomain()
	if len(qs.Body.DomainRecords.Record) == 0 {
		log.Infof("本地ip:%v\t远端不存在此解析\t现在开始添加!...", qs.Ipv6)
		qs.addDomain("AAAA", qs.Ipv6)
		return
	}
	for _, record := range qs.Body.DomainRecords.Record {
		if record.Type == "AAAA" {
			if qs.Ipv6 == record.Value {
				log.Infof("本地ip:%v\t远端ip:%v\t无需更新", qs.Ipv6, record.Value)
				return
			} else {
				log.Infof("本地ip:%v\t远端ip:%v\t现在更新", qs.Ipv6, record.Value)
				qs.updateDomain("AAAA", qs.Ipv6)
				return
			}
		}
	}
}

func CertCheck() {
	qs.SubDomain = dnsHeader + qs.SubDomain
	qs.Ipv6 = *Key
	fmt.Println(qs.Ipv6)
	qs.CreateClient()
	qs.queryDomain()
	if len(qs.Body.DomainRecords.Record) == 0 {
		log.Infof("本地ip:%v\t远端不存在此解析\t现在开始添加!...", qs.Ipv6)
		qs.addDomain("TXT", qs.Ipv6)
		return
	}
	for _, record := range qs.Body.DomainRecords.Record {
		if record.Type == "TXT" {
			if qs.Ipv6 == record.Value {
				log.Infof("本地ip:%v\t远端ip:%v\t无需更新", qs.Ipv6, record.Value)
				return
			} else {
				log.Infof("本地ip:%v\t远端ip:%v\t现在更新", qs.Ipv6, record.Value)
				qs.updateDomain("TXT", qs.Ipv6)
				return
			}
		}
	}
}
