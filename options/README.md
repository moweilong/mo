# options 包

options 包提供了一套统一的配置选项管理机制，用于处理各种服务组件的配置，包括数据库连接、HTTP服务、TLS安全连接、健康检查等。该包定义了统一的接口和辅助函数，使配置管理更加规范化和便捷。

## 功能概述

- 提供统一的配置选项接口 `IOptions`
- 支持多种服务组件的配置定义（MySQL、PostgreSQL、Redis、HTTP、TLS、健康检查）
- 提供默认配置初始化方法
- 支持命令行参数绑定
- 集成配置验证功能
- 提供与其他包（如 gormx、cache）的便捷集成

## 文件说明

| 文件名 | 说明 |
|-------|------|
| options.go | 核心接口定义（IOptions） |
| helper.go | 辅助函数（地址验证、监听器创建等） |
| mysql_options.go | MySQL 数据库连接配置 |
| postgresql_options.go | PostgreSQL 数据库连接配置 |
| redis_options.go | Redis 缓存连接配置 |
| http_options.go | HTTP 服务器配置 |
| tls_options.go | TLS 安全连接配置 |
| health_options.go | 健康检查服务配置 |

## 核心接口

`IOptions` 接口是所有配置选项结构体的基础：

```go
// IOptions 定义了实现通用选项的方法
type IOptions interface {
    // Validate 验证所有必需的选项，也可用于完成选项
    Validate() []error

    // AddFlags 将与给定标志集相关的标志添加到指定的 FlagSet
    AddFlags(fs *pflag.FlagSet, prefixes ...string)
}
```

## 配置选项详细说明

### MySQL 配置

`MySQLOptions` 结构体用于配置 MySQL 数据库连接：

```go
type MySQLOptions struct {
    Addr                  string        // 数据库地址（host:port）
    Username              string        // 用户名
    Password              string        // 密码
    Database              string        // 数据库名
    MaxIdleConnections    int           // 最大空闲连接数
    MaxOpenConnections    int           // 最大打开连接数
    MaxConnectionLifeTime time.Duration // 连接最大生命周期
    LogLevel              int           // 日志级别
}
```

#### 主要方法

```go
// NewMySQLOptions 创建带有默认值的 MySQL 选项
func NewMySQLOptions() *MySQLOptions

// DSN 返回 MySQL 连接字符串
func (o *MySQLOptions) DSN() string

// NewDB 创建 MySQL 数据库连接
func (o *MySQLOptions) NewDB() (*gorm.DB, error)
```

### PostgreSQL 配置

`PostgreSQLOptions` 结构体用于配置 PostgreSQL 数据库连接：

```go
type PostgreSQLOptions struct {
    Addr                  string        // 数据库地址（host:port）
    Username              string        // 用户名
    Password              string        // 密码
    Database              string        // 数据库名
    MaxIdleConnections    int           // 最大空闲连接数
    MaxOpenConnections    int           // 最大打开连接数
    MaxConnectionLifeTime time.Duration // 连接最大生命周期
    LogLevel              int           // 日志级别
}
```

#### 主要方法

```go
// NewPostgreSQLOptions 创建带有默认值的 PostgreSQL 选项
func NewPostgreSQLOptions() *PostgreSQLOptions

// NewDB 创建 PostgreSQL 数据库连接
func (o *PostgreSQLOptions) NewDB() (*gorm.DB, error)
```

### Redis 配置

`RedisOptions` 结构体用于配置 Redis 缓存连接：

```go
type RedisOptions struct {
    Addr         string        // Redis 服务器地址（ip:port）
    Username     string        // 用户名
    Password     string        // 密码
    Database     int           // 数据库索引
    MaxRetries   int           // 最大重试次数
    MinIdleConns int           // 最小空闲连接数
    DialTimeout  time.Duration // 拨号超时时间
    ReadTimeout  time.Duration // 读取超时时间
    WriteTimeout time.Duration // 写入超时时间
    PoolTimeout  time.Duration // 连接池超时时间
    PoolSize     int           // 连接池大小
    EnableTrace  bool          // 是否启用跟踪
}
```

#### 主要方法

```go
// NewRedisOptions 创建带有默认值的 Redis 选项
func NewRedisOptions() *RedisOptions

// NewClient 创建 Redis 客户端连接
func (o *RedisOptions) NewClient() (*redis.Client, error)
```

### HTTP 配置

`HTTPOptions` 结构体用于配置 HTTP 服务器：

```go
type HTTPOptions struct {
    Network string        // 网络类型（如 tcp）
    Addr    string        // 服务器地址（ip:port）
    Timeout time.Duration // 服务器超时时间
}
```

#### 主要方法

```go
// NewHTTPOptions 创建带有默认值的 HTTP 选项
func NewHTTPOptions() *HTTPOptions

// Complete 填充未设置但需要有效数据的字段
func (s *HTTPOptions) Complete() error
```

### TLS 配置

`TLSOptions` 结构体用于配置 TLS 安全连接：

```go
type TLSOptions struct {
    UseTLS             bool   // 是否使用 TLS
    InsecureSkipVerify bool   // 是否跳过服务器证书验证
    CaCert             string // CA 证书路径
    Cert               string // 证书路径
    Key                string // 密钥路径
}
```

#### 主要方法

```go
// NewTLSOptions 创建 TLS 选项
func NewTLSOptions() *TLSOptions

// MustTLSConfig 获取 TLS 配置，如果出错返回默认配置
func (o *TLSOptions) MustTLSConfig() *tls.Config

// TLSConfig 获取 TLS 配置
func (o *TLSOptions) TLSConfig() (*tls.Config, error)

// Scheme 根据 TLS 配置返回 URL 协议（http 或 https）
func (o *TLSOptions) Scheme() string
```

### 健康检查配置

`HealthOptions` 结构体用于配置健康检查服务：

```go
type HealthOptions struct {
    HTTPProfile        bool   // 是否启用 HTTP 性能分析
    HealthCheckPath    string // 健康检查路径
    HealthCheckAddress string // 健康检查绑定地址
}
```

#### 主要方法

```go
// NewHealthOptions 创建带有默认值的健康检查选项
func NewHealthOptions() *HealthOptions

// ServeHealthCheck 启动健康检查服务
func (o *HealthOptions) ServeHealthCheck()
```

## 辅助函数

`helper.go` 文件提供了以下辅助函数：

```go
// ValidateAddress 验证地址是否有效（:port 或 ip:port 格式）
func ValidateAddress(addr string) error

// CreateListener 创建网络监听器并返回它和端口号
func CreateListener(addr string) (net.Listener, int, error)
```

## 使用示例

### 基本用法

以下是如何使用 options 包的基本示例：

```go
import (
    "github.com/spf13/pflag"
    "github.com/moweilong/mo/options"
)

// 创建配置选项
mysqlOpts := options.NewMySQLOptions()
redisOpts := options.NewRedisOptions()

// 添加命令行标志
fs := pflag.NewFlagSet("example", pflag.ExitOnError)
mysqlOpts.AddFlags(fs)
redisOpts.AddFlags(fs)

// 解析命令行参数
fs.Parse(os.Args[1:])

// 验证配置
if errs := mysqlOpts.Validate(); len(errs) > 0 {
    // 处理错误
}

// 使用配置创建客户端
mysqlDB, err := mysqlOpts.NewDB()
if err != nil {
    // 处理错误
}

redisClient, err := redisOpts.NewClient()
if err != nil {
    // 处理错误
}
```

### 带前缀的命令行标志

可以为命令行标志添加前缀，适用于多组件配置场景：

```go
// 为不同组件添加带前缀的标志
primaryDBOpts.AddFlags(fs, "primary")
secondaryDBOpts.AddFlags(fs, "secondary")

// 这样会生成如下命令行标志：
// --primary.mysql.host
// --primary.mysql.username
// ...
// --secondary.mysql.host
// --secondary.mysql.username
// ...
```

### 启动健康检查服务

```go
import "github.com/moweilong/mo/options"

// 创建健康检查选项
healthOpts := options.NewHealthOptions()
healthOpts.HealthCheckAddress = "0.0.0.0:8080"
healthOpts.HealthCheckPath = "/healthz"
healthOpts.HTTPProfile = true // 启用性能分析

// 在单独的 goroutine 中启动健康检查服务
go healthOpts.ServeHealthCheck()

// 健康检查服务将提供以下端点：
// - /healthz - 健康检查
// - /debug/pprof/... - 性能分析（如果启用）
```

### 使用 TLS 配置

```go
import "github.com/moweilong/mo/options"

// 创建 TLS 选项
tlsOpts := options.NewTLSOptions()
tlsOpts.UseTLS = true
tlsOpts.Cert = "/path/to/cert.pem"
tlsOpts.Key = "/path/to/key.pem"
tlsOpts.CaCert = "/path/to/ca.pem"

// 获取 TLS 配置
tlsConfig, err := tlsOpts.TLSConfig()
if err != nil {
    // 处理错误
}

// 使用 TLS 配置
httpServer := &http.Server{
    Addr:      ":443",
    TLSConfig: tlsConfig,
}
```

## 注意事项

1. 所有选项结构体都实现了 `IOptions` 接口，确保了统一的使用方式
2. 每个选项结构体都提供了 `NewXXXOptions()` 方法来创建带有默认值的实例
3. 使用 `AddFlags()` 方法可以方便地将配置选项绑定到命令行参数
4. 在使用配置前，应调用 `Validate()` 方法进行验证
5. 一些配置选项结构体提供了便捷方法（如 `NewDB()`、`NewClient()`）来直接创建相应的客户端

## 依赖说明

- github.com/spf13/pflag - 命令行参数解析
- gorm.io/gorm - 数据库 ORM 框架
- github.com/redis/go-redis/v9 - Redis 客户端
- github.com/moweilong/mo/gormx - GORM 扩展包
- github.com/moweilong/mo/cache - 缓存包
- github.com/moweilong/mo/log - 日志包