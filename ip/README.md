# IP 工具包

## 功能概述

IP 工具包是一个提供 IP 地址相关功能的工具库，目前主要提供获取本地机器出站 IP 地址的功能。

## 主要函数

### GetOutboundIP
```go
func GetOutboundIP() net.IP
```

获取本机的出站 IP 地址。该函数是GetOutboundIPWithDNS的包装函数，默认使用阿里云公共DNS服务器（223.5.5.5:80）。

返回值：
- `net.IP`：本机的出站 IP 地址，如果获取失败则返回空字节切片。

### GetOutboundIPWithDNS
```go
func GetOutboundIPWithDNS(dnsAddr string) net.IP
```

使用自定义DNS地址获取本机的出站 IP 地址。该函数通过创建一个到指定DNS服务器的 UDP 连接（不实际发送数据）来确定本机的出站 IP 地址。

参数：
- `dnsAddr`：DNS服务器地址，格式为"host:port"，例如"8.8.8.8:53"

返回值：
- `net.IP`：本机的出站 IP 地址，如果获取失败则返回空字节切片。

## 安装

```bash
go get github.com/moweilong/mo/ip
```

## 使用示例

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/moweilong/mo/ip"
)

func main() {
	// 使用默认DNS获取本机出站IP
	outboundIP := ip.GetOutboundIP()
	
	if outboundIP != nil && len(outboundIP) > 0 {
		fmt.Printf("本机出站IP地址: %s\n", outboundIP.String())
	} else {
		fmt.Println("获取出站IP地址失败")
	}
}
```

### 使用自定义DNS

```go
package main

import (
	"fmt"
	"github.com/moweilong/mo/ip"
)

func main() {
	// 使用Google DNS获取本机出站IP
	ipWithGoogleDNS := ip.GetOutboundIPWithDNS("8.8.8.8:53")
	if ipWithGoogleDNS != nil && len(ipWithGoogleDNS) > 0 {
		fmt.Printf("使用Google DNS获取的IP地址: %s\n", ipWithGoogleDNS.String())
	} else {
		fmt.Println("使用Google DNS获取IP地址失败")
	}
	
	// 使用OpenDNS获取本机出站IP
	ipWithOpenDNS := ip.GetOutboundIPWithDNS("208.67.222.222:53")
	if ipWithOpenDNS != nil && len(ipWithOpenDNS) > 0 {
		fmt.Printf("使用OpenDNS获取的IP地址: %s\n", ipWithOpenDNS.String())
	} else {
		fmt.Println("使用OpenDNS获取IP地址失败")
	}
}
```

### 结合其他网络操作

```go
package main

import (
	"fmt"
	"github.com/moweilong/mo/ip"
	"net/http"
	"time"
)

func main() {
	// 获取本机出站IP
	localIP := ip.GetOutboundIP()
	fmt.Printf("本机出站IP: %s\n", localIP)
	
	// 创建HTTP客户端并发送请求
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("HTTP请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Println("HTTP请求成功，服务器看到的IP可能与上面的出站IP一致")
}
```

## 测试

该包包含简单的测试用例，可以通过以下命令运行：

```bash
cd /path/to/github.com/moweilong/mo/ip
go test -v
```

测试用例主要验证：
1. 是否能获取到有效的 IP 地址
2. 获取的 IP 地址格式是否正确
3. 网络错误情况下的处理

## 注意事项

- 该包中的函数依赖于网络连接，需要能够访问外部网络
- GetOutboundIP函数默认使用阿里云公共DNS服务器（223.5.5.5:80）
- GetOutboundIPWithDNS函数允许指定自定义的DNS服务器地址
- 所有函数使用UDP协议连接到指定的DNS服务器，但不会实际发送数据
- **重要说明**：由于UDP协议的无连接特性，即使指定的DNS地址无效或不可达，操作系统也可能会立即分配本地地址，因此函数可能仍会返回一个IP地址
- 在某些网络环境下（如防火墙限制），可能无法获取到正确的出站IP
- 选择DNS服务器时，建议使用稳定可靠的公共DNS服务，如Google DNS（8.8.8.8:53）、OpenDNS（208.67.222.222:53）等
- 如果获取失败，函数会打印错误信息并返回空字节切片