# netutil

netutil 是一个提供网络相关功能的工具包，主要提供获取本地机器出站 IP 地址的功能。

## 功能概述

该包主要实现了以下功能：
- 获取本地机器的出站 IP 地址（即用于连接外部网络的 IP 地址）
- 支持使用默认 DNS 服务器或自定义 DNS 服务器地址获取出站 IP

## 文件说明

- `netutil.go`：实现了获取出站 IP 的核心功能
- `netutil_test.go`：包含了对功能的单元测试

## 主要函数

### GetOutboundIP

```go
func GetOutboundIP() net.IP
```

使用默认的阿里云公共 DNS (223.5.5.5:80) 获取本地出口 IP 地址。

**返回值**：
- `net.IP`：本地机器的出站 IP 地址，如果获取失败则返回空字节数组 []byte{}

### GetOutboundIPWithDNS

```go
func GetOutboundIPWithDNS(dnsAddr string) net.IP
```

使用自定义 DNS 地址获取本地出口 IP 地址。

**参数**：
- `dnsAddr`：DNS 服务器地址，格式为 "host:port"，例如 "8.8.8.8:53"

**返回值**：
- `net.IP`：本地机器的出站 IP 地址，如果获取失败则返回空字节数组 []byte{}

## 使用示例

以下是如何使用 netutil 包获取本地出站 IP 地址的示例：

```go
import (
    "fmt"
    "github.com/moweilong/mo/netutil"
)

func main() {
    // 使用默认 DNS 获取出站 IP
    defaultIP := netutil.GetOutboundIP()
    fmt.Printf("使用默认 DNS 获取的出站 IP: %s\n", defaultIP.String())
    
    // 使用自定义 DNS 获取出站 IP
    customIP := netutil.GetOutboundIPWithDNS("8.8.8.8:53")
    fmt.Printf("使用 Google DNS 获取的出站 IP: %s\n", customIP.String())
}
```

## 注意事项

1. 获取出站 IP 的方法基于 UDP 连接，由于 UDP 是无连接协议，即使指定的 DNS 服务器地址不可达，函数通常也能返回本地 IP 地址。

2. 如果获取 IP 失败，函数会在控制台打印错误信息并返回空字节数组。在实际应用中，建议检查返回值是否为空。

3. 该包依赖于 Go 标准库中的 `net` 包，不需要额外的第三方依赖。

## 依赖

- Go 标准库：net
- Go 标准库：fmt

## 测试

包中包含完整的单元测试，测试了正常情况下获取 IP、使用自定义 DNS 获取 IP 以及使用无效 DNS 时的行为。可以使用以下命令运行测试：

```bash
go test -v ./netutil
```