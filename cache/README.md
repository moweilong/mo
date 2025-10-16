# cache 包

cache 包提供了缓存连接管理功能，目前主要实现了 Redis 客户端的创建和依赖注入支持。

## 功能概述

- 提供 Redis 客户端的统一创建接口
- 支持详细的 Redis 连接配置
- 集成 Google Wire 依赖注入
- 自动验证 Redis 连接可用性

## 文件说明

- `redis.go`: 定义 Redis 连接配置和客户端创建函数
- `wire.go`: 提供依赖注入支持

## Redis 连接配置

`RedisOptions` 结构体包含以下配置选项：

| 字段名 | 类型 | 说明 |
|-------|------|------|
| Addr | string | Redis 服务器地址，格式为 `host:port` |
| Username | string | Redis 用户名（用于 Redis 6.0+ ACL） |
| Password | string | Redis 密码 |
| Database | int | Redis 数据库索引 |
| MaxRetries | int | 最大重试次数 |
| MinIdleConns | int | 最小空闲连接数 |
| DialTimeout | time.Duration | 拨号超时时间 |
| ReadTimeout | time.Duration | 读取超时时间 |
| WriteTimeout | time.Duration | 写入超时时间 |
| PoolTimeout | time.Duration | 连接池超时时间 |
| PoolSize | int | 连接池大小 |

## 主要函数

### NewRedis

```go
func NewRedis(opts *RedisOptions) (*redis.Client, error)
```

创建并返回一个 Redis 客户端实例。函数会使用提供的选项配置客户端，并通过 Ping 操作验证连接是否正常。

#### 参数
- `opts`: Redis 连接配置选项

#### 返回值
- `*redis.Client`: Redis 客户端实例
- `error`: 错误信息，连接失败时返回

## 依赖注入

`ProviderSet` 变量提供了用于 Google Wire 依赖注入的提供者集合：

```go
var ProviderSet = wire.NewSet(
    NewRedis,
    wire.Bind(new(redis.UniversalClient), new(*redis.Client)),
)
```

该集合会：
1. 提供 `NewRedis` 函数作为 Redis 客户端创建器
2. 将 `*redis.Client` 实现绑定到 `redis.UniversalClient` 接口

## 安装

确保在你的项目中添加以下依赖：

```bash
go get github.com/redis/go-redis/v9
go get github.com/google/wire
```

## 使用示例

### 基本用法

```go
import (
    "github.com/moweilong/mo/cache"
    "time"
)

// 创建 Redis 配置
opts := &cache.RedisOptions{
    Addr:         "localhost:6379",
    Password:     "your-password",
    Database:     0,
    MaxRetries:   3,
    MinIdleConns: 5,
    DialTimeout:  5 * time.Second,
    ReadTimeout:  3 * time.Second,
    WriteTimeout: 3 * time.Second,
    PoolSize:     10,
}

// 创建 Redis 客户端
client, err := cache.NewRedis(opts)
if err != nil {
    // 处理错误
}

// 使用客户端
val, err := client.Get(ctx, "key").Result()
```

### 与 Google Wire 集成

```go
import (
    "github.com/google/wire"
    "github.com/moweilong/mo/cache"
)

// 定义依赖注入提供者
var ProviderSet = wire.NewSet(
    // 提供 Redis 配置
    NewRedisConfig,
    // 包含 cache 包的提供者
    cache.ProviderSet,
)

// 创建 Redis 配置函数
func NewRedisConfig() *cache.RedisOptions {
    return &cache.RedisOptions{
        Addr:     "localhost:6379",
        Password: "your-password",
        Database: 0,
    }
}
```

## 注意事项

1. 确保在使用 Redis 客户端前检查错误返回值
2. 对于生产环境，建议合理配置连接池大小和超时参数
3. Redis 6.0+ 版本才支持用户名验证
4. 当不再使用客户端时，建议调用 `Close()` 方法释放资源

## 依赖

- [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) - Redis 客户端库
- [github.com/google/wire](https://github.com/google/wire) - 依赖注入工具