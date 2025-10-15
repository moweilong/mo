package ip

import (
	"net"
	"testing"
)

// TestGetOutboundIP 测试获取出站IP的功能
func TestGetOutboundIP(t *testing.T) {
	// 获取本地出站IP
	ip := GetOutboundIP()

	// 验证IP地址是否有效
	if len(ip) == 0 {
		t.Errorf("期望获取到有效IP地址，但得到了空IP")
		return
	}

	// 检查IP地址格式是否正确
	ipStr := ip.String()
	parsedIP := net.ParseIP(ipStr)
	if parsedIP == nil {
		t.Errorf("获取的IP地址格式无效: %s", ipStr)
		return
	}

	// 打印获取到的IP地址，便于调试
	t.Logf("成功获取出站IP: %s", ipStr)
}

// TestGetOutboundIPWithDNS 测试使用自定义DNS地址获取出站IP的功能
func TestGetOutboundIPWithDNS(t *testing.T) {
	// 测试使用Google DNS
	ip := GetOutboundIPWithDNS("8.8.8.8:53")

	// 验证IP地址是否有效
	if len(ip) == 0 {
		t.Errorf("使用Google DNS期望获取到有效IP地址，但得到了空IP")
		return
	}

	// 检查IP地址格式是否正确
	ipStr := ip.String()
	parsedIP := net.ParseIP(ipStr)
	if parsedIP == nil {
		t.Errorf("使用Google DNS获取的IP地址格式无效: %s", ipStr)
		return
	}

	// 打印获取到的IP地址，便于调试
	t.Logf("使用Google DNS成功获取出站IP: %s", ipStr)

	// 测试使用阿里云DNS（与默认相同）
	ip2 := GetOutboundIPWithDNS("223.5.5.5:80")
	if len(ip2) == 0 {
		t.Errorf("使用阿里云DNS期望获取到有效IP地址，但得到了空IP")
		return
	}
	t.Logf("使用阿里云DNS成功获取出站IP: %s", ip2.String())
}

// TestGetOutboundIPWithInvalidDNS 测试使用无效DNS地址时的行为
func TestGetOutboundIPWithInvalidDNS(t *testing.T) {
	// 使用RFC 5737中定义的专用于文档和测试的IP地址
	invalidDNS := "192.0.2.0:12345"
	ip := GetOutboundIPWithDNS(invalidDNS)

	// 注意：由于UDP的无连接特性，即使远程地址无效，
	// 操作系统也会立即分配本地地址，所以通常仍能获取到IP
	// 此测试主要验证函数在面对各种DNS地址时的稳定性
	if len(ip) == 0 {
		t.Logf("使用无效DNS地址 %s 时返回了空IP", invalidDNS)
	} else {
		t.Logf("使用无效DNS地址 %s 仍获取到了IP: %s\n注意：这是UDP协议的正常行为，因为UDP连接建立不验证远程地址的可达性", invalidDNS, ip.String())
	}
}
