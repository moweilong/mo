package ip

import (
	"fmt"
	"net"
)

// GetOutboundIP 使用默认的阿里云公共DNS (223.5.5.5:80) 获取本地出口IP地址
func GetOutboundIP() net.IP {
	return GetOutboundIPWithDNS("223.5.5.5:80")
}

// GetOutboundIPWithDNS 使用自定义DNS地址获取本地出口IP地址
// dnsAddr: DNS服务器地址，格式为"host:port"，例如"8.8.8.8:53"
func GetOutboundIPWithDNS(dnsAddr string) net.IP {
	conn, err := net.Dial("udp", dnsAddr)
	if err != nil {
		fmt.Println("get outbound ip fail:", err)
		return []byte{}
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
