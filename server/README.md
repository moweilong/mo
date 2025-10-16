# server 包

`server` 包是一个统一的服务器管理模块，提供了多种服务器类型的实现和统一的生命周期管理。

## 功能概述

该包主要提供了以下功能：
- 定义统一的 `Server` 接口，规范服务器的启动和优雅关闭行为
- 提供通用的 `Serve` 函数，管理服务器的生命周期
- 实现 HTTP 服务器，支持 TLS
- 实现 gRPC 服务器，支持 TLS 和健康检查
- 实现 Kratos 框架服务器，支持服务注册（etcd/consul）
- 实现 gRPC 网关服务器，支持 HTTP 到 gRPC 的反向代理

## 文件说明

- **server.go**: 定义统一的服务器接口和通用服务函数
- **http_server.go**: HTTP 服务器实现
- **grpc_server.go**: gRPC 服务器实现
- **kratos_server.go**: Kratos 框架服务器实现
- **reverse_proxy_server.go**: gRPC 网关服务器实现

## 核心接口

### Server 接口

```go
// Server 定义所有服务器类型的接口
type Server interface {
    // RunOrDie 运行服务器，如果运行失败会退出程序
    RunOrDie()
    // GracefulStop 方法用来优雅关停服务器
    GracefulStop(ctx context.Context)
}
```

## 主要函数和结构体

### 通用服务函数

```go
// Serve starts the server and blocks until the context is canceled
func Serve(ctx context.Context, srv Server) error
```

### HTTP 服务器

```go
// HTTPServer 代表一个 HTTP 服务器
type HTTPServer struct {
    srv *http.Server
}

// NewHTTPServer 创建一个新的 HTTP 服务器实例
func NewHTTPServer(httpOptions *genericoptions.HTTPOptions, tlsOptions *genericoptions.TLSOptions, handler http.Handler) *HTTPServer
```

### gRPC 服务器

```go
// GRPCServer 代表一个 GRPC 服务器
type GRPCServer struct {
    srv *grpc.Server
    lis net.Listener
}

// NewGRPCServer 创建一个新的 GRPC 服务器实例
func NewGRPCServer(
    grpcOptions *genericoptions.GRPCOptions,
    tlsOptions *genericoptions.TLSOptions,
    serverOptions []grpc.ServerOption,
    registerServer func(grpc.ServiceRegistrar),
) (*GRPCServer, error)
```

### Kratos 服务器

```go
// KratosServer 代表一个 Kratos 框架服务器
type KratosServer struct {
    kapp *kratos.App
}

// NewKratosServer 创建一个新的 Kratos 服务器实例
func NewKratosServer(cfg KratosAppConfig, servers ...transport.Server) (*KratosServer, error)

// NewEtcdRegistrar 创建 etcd 服务注册器
func NewEtcdRegistrar(opts *genericoptions.EtcdOptions) registry.Registrar

// NewConsulRegistrar 创建 consul 服务注册器
func NewConsulRegistrar(opts *genericoptions.ConsulOptions) registry.Registrar
```

### gRPC 网关服务器

```go
// GRPCGatewayServer 代表一个 GRPC 网关服务器
type GRPCGatewayServer struct {
    srv *http.Server
}

// NewGRPCGatewayServer 创建一个新的 GRPC 网关服务器实例
func NewGRPCGatewayServer(
    httpOptions *genericoptions.HTTPOptions,
    grpcOptions *genericoptions.GRPCOptions,
    tlsOptions *genericoptions.TLSOptions,
    registerHandler func(mux *runtime.ServeMux, conn *grpc.ClientConn) error,
) (*GRPCGatewayServer, error)
```

## 使用示例

### HTTP 服务器示例

```go
httpOptions := &genericoptions.HTTPOptions{
    Addr: ":8080",
}

handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
})

httpServer := NewHTTPServer(httpOptions, nil, handler)

ctx, cancel := context.WithCancel(context.Background())

go func() {
    // 等待停止信号
    // ...
    cancel()
}()

if err := Serve(ctx, httpServer); err != nil {
    log.Fatalw(err, "Failed to serve HTTP server")
}
```

### gRPC 服务器示例

```go
grpcOptions := &genericoptions.GRPCOptions{
    Addr: ":9090",
}

// 注册 gRPC 服务的函数
registerServer := func(s grpc.ServiceRegistrar) {
    // 注册你的 gRPC 服务
    // protoexample.RegisterExampleServer(s, &exampleServer{})
}

// 创建 gRPC 服务器选项
serverOptions := []grpc.ServerOption{
    // 添加你的服务器选项
}

grpcServer, err := NewGRPCServer(grpcOptions, nil, serverOptions, registerServer)
if err != nil {
    log.Fatalw(err, "Failed to create gRPC server")
}

ctx, cancel := context.WithCancel(context.Background())

go func() {
    // 等待停止信号
    // ...
    cancel()
}()

if err := Serve(ctx, grpcServer); err != nil {
    log.Fatalw(err, "Failed to serve gRPC server")
}
```

### gRPC 网关服务器示例

```go
httpOptions := &genericoptions.HTTPOptions{
    Addr: ":8080",
}

grpcOptions := &genericoptions.GRPCOptions{
    Addr: ":9090",
}

// 注册 gRPC 网关处理器的函数
registerHandler := func(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
    // 注册你的 gRPC 网关处理器
    // return protoexample.RegisterExampleHandler(conn, mux)
    return nil
}

gatewayServer, err := NewGRPCGatewayServer(httpOptions, grpcOptions, nil, registerHandler)
if err != nil {
    log.Fatalw(err, "Failed to create gRPC gateway server")
}

ctx, cancel := context.WithCancel(context.Background())

go func() {
    // 等待停止信号
    // ...
    cancel()
}()

if err := Serve(ctx, gatewayServer); err != nil {
    log.Fatalw(err, "Failed to serve gRPC gateway server")
}
```

### Kratos 服务器示例

```go
import (
    "time"
    "github.com/go-kratos/kratos/v2/transport/http"
    "github.com/go-kratos/kratos/v2/transport/grpc"
)

// 创建 etcd 注册器
etcdOptions := &genericoptions.EtcdOptions{
    Endpoints:   []string{"localhost:2379"},
    DialTimeout: 5 * time.Second,
}
registrar := NewEtcdRegistrar(etcdOptions)

// 创建 Kratos HTTP 服务器
httpSrv := http.NewServer(http.Address(":8080"))
// 创建 Kratos gRPC 服务器
grpcSrv := grpc.NewServer(grpc.Address(":9090"))

// 创建 Kratos 应用配置
appConfig := KratosAppConfig{
    ID:       "my-app",
    Name:     "my-service",
    Version:  "1.0.0",
    Metadata: map[string]string{"env": "development"},
    Registrar: registrar,
}

// 创建 Kratos 服务器
kratosServer, err := NewKratosServer(appConfig, httpSrv, grpcSrv)
if err != nil {
    log.Fatalw(err, "Failed to create Kratos server")
}

ctx, cancel := context.WithCancel(context.Background())

go func() {
    // 等待停止信号
    // ...
    cancel()
}()

if err := Serve(ctx, kratosServer); err != nil {
    log.Fatalw(err, "Failed to serve Kratos server")
}
```

## 注意事项

1. 所有服务器都实现了 `Server` 接口，可以使用统一的 `Serve` 函数进行管理
2. TLS 配置可以通过 `genericoptions.TLSOptions` 传递给各个服务器
3. gRPC 服务器自动集成了健康检查服务
4. Kratos 服务器支持 etcd 和 consul 两种服务注册方式
5. gRPC 网关服务器默认配置了 Protobuf JSON 序列化选项，枚举类型以数字格式输出
6. 所有服务器都支持优雅关闭，确保在关闭时处理完所有请求

## 依赖

- 标准库：`context`, `net/http`, `crypto/tls`, `net` 等
- gRPC 相关：`google.golang.org/grpc`, `google.golang.org/grpc/health` 等
- gRPC 网关：`github.com/grpc-ecosystem/grpc-gateway/v2`
- Kratos 框架：`github.com/go-kratos/kratos/v2`
- 服务注册：`go.etcd.io/etcd/client/v3`, `github.com/hashicorp/consul/api`
- 项目内部：`github.com/moweilong/mo/log`, `github.com/moweilong/mo/options`