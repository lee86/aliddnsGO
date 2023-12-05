package aliddns

import (
	log "github.com/sirupsen/logrus"
	"net"
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
	conn, err := net.Dial("udp", "223.5.5.5:53")
	if err != nil {
		log.Fatal(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	qs.Value = strings.Split(localAddr.String(), ":")[0]
}
