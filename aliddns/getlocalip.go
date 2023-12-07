package aliddns

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"strings"
)

// GetOutBoundIP 获取本地公网ip
func (qs *QueryStruct) GetOutBoundIP() {
	switch qs.ValueType {
	case "A":
		qs.getOutBoundIPV4()
	case "AAAA":
		qs.getOutBoundIPV6()
	default:
		log.Error("不支持此类型")
	}
}

// GetOutBoundIPV6 获取本地可访问的IPV6
func (qs *QueryStruct) getOutBoundIPV6() {
	conn, err := net.Dial("udp6", "[2400:3200::1]:53")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), "]")[0]
	qs.Value = strings.Replace(ip, "[", "", -1)
}

// GetOutBoundIPV4 获取本地可访问的IPV4
func (qs *QueryStruct) getOutBoundIPV4() {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	ip, _ := io.ReadAll(resp.Body)
	qs.Value = fmt.Sprintf("%v", string(ip))
}
