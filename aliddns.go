package main

import (
	"bufio"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
)

const (
	ak_id   = "aliddns.base.app_key"
	ak_sec  = "aliddns.base.app_secret"
	rrid    = "aliddns.base.record_id"
	main_dm = "aliddns.base.main_domain"
	sub_dm  = "aliddns.base.sub_domain"
	iface   = "aliddns.base.interface"
	test    = "ip addr"
)

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "114.114.115.115:53")
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
func getKeys(key string) string {
	cmd := exec.Command("/bin/sh", "-c", "uci -q get "+key)
	stdout, err := cmd.StdoutPipe()
	checkError(err)
	err = cmd.Start()
	checkError(err)
	var value string
	reader := bufio.NewReader(stdout)
	for {
		value, err = reader.ReadString('_')
		if err != nil || io.EOF == err {
			break
		}
	}
	value = strings.Replace(value, "\r\n", "", -1)
	value = strings.Replace(value, "\r", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.Replace(value, " ", "", -1)
	return value
}
