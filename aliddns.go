package main

import (
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
