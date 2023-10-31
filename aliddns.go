package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func GetOutBoundIPV6() (ip string, err error) {
	conn, err := net.Dial("udp6", "[2400:3200::1]:53")
	if err != nil {
		log.Fatal(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), "]")[0]
	ip = strings.Replace(ip, "[", "", -1)
	return
}
func GetOutBoundIPV4() (ip string, err error) {
	conn, err := net.Dial("udp", "223.5.5.5:53")
	if err != nil {
		log.Fatal(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
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
