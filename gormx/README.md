# gormx 包

gormx 包提供了 GORM 数据库框架的扩展功能，包括数据库连接管理、性能跟踪插件和依赖注入支持。

## 功能概述

- 提供 MySQL 和 PostgreSQL 数据库的统一连接创建接口
- 支持详细的数据库连接配置和连接池设置
- 包含 SQL 执行性能跟踪插件
- 集成 Google Wire 依赖注入
- 自动设置合理的默认配置值

## 文件说明

- `plugin.go`: 实现 SQL 执行性能跟踪插件
- `mysql.go`: 提供 MySQL 数据库连接创建功能
- `postgresql.go`: 提供 PostgreSQL 数据库连接创建功能
- `wire.go`: 定义依赖注入提供者集合

## 数据库连接配置

### MySQL 配置

`MySQLOptions` 结构体包含以下配置选项：

| 字段名 | 类型 | 说明 | 默认值 |
|-------|------|------|-------|
| Addr | string | MySQL 服务器地址，格式为 `host:port` | 127.0.0.1:3306 |
| Username | string | 数据库用户名 | - |
| Password | string | 数据库密码 | - |
| Database | string | 数据库名称 | - |
| MaxIdleConnections | int | 最大空闲连接数 | 100 |
| MaxOpenConnections | int | 最大打开连接数 | 100 |
| MaxConnectionLifeTime | time.Duration | 连接最大生命周期 | 10秒 |
| Logger | logger.Interface | GORM 日志记录器 | logger.Default |

### PostgreSQL 配置

`PostgreSQLOptions` 结构体包含以下配置选项：

| 字段名 | 类型 | 说明 | 默认值 |
|-------|------|------|-------|
| Addr | string | PostgreSQL 服务器地址，格式为 `host:port` | 127.0.0.1:5432 |
| Username | string | 数据库用户名 | - |
| Password | string | 数据库密码 | - |
| Database | string | 数据库名称 | - |
| MaxIdleConnections | int | 最大空闲连接数 | 100 |
| MaxOpenConnections | int | 最大打开连接数 | 100 |
| MaxConnectionLifeTime | time.Duration | 连接最大生命周期 | 10秒 |
| Logger | logger.Interface | GORM 日志记录器 | logger.Default |

## 主要函数

### NewMySQL

```go
func NewMySQL(opts *MySQLOptions) (*gorm.DB, error)
```

创建并返回一个 MySQL 数据库的 GORM 客户端实例。函数会自动应用默认配置，设置连接池参数，并启用预编译语句以提高性能。

#### 参数
- `opts`: MySQL 连接配置选项

#### 返回值
- `*gorm.DB`: GORM 数据库实例
- `error`: 错误信息，连接失败时返回

### NewPostgreSQL

```go
func NewPostgreSQL(opts *PostgreSQLOptions) (*gorm.DB, error)
```

创建并返回一个 PostgreSQL 数据库的 GORM 客户端实例。函数会自动应用默认配置，设置连接池参数，并启用预编译语句以提高性能。

#### 参数
- `opts`: PostgreSQL 连接配置选项

#### 返回值
- `*gorm.DB`: GORM 数据库实例
- `error`: 错误信息，连接失败时返回

### MustRawDB

```go
func MustRawDB(db *gorm.DB) *sql.DB
```

从 GORM 数据库实例中获取原始的 SQL 数据库连接。如果获取失败，函数会 panic。

#### 参数
- `db`: GORM 数据库实例

#### 返回值
- `*sql.DB`: 原始 SQL 数据库连接

## 性能跟踪插件

`TracePlugin` 是一个 GORM 插件，用于跟踪 SQL 查询的执行时间。

```go
// TracePlugin defines gorm plugin used to trace sql.
type TracePlugin struct{}

// Name returns the name of trace plugin.
func (op *TracePlugin) Name() string

// Initialize initialize the trace plugin.
func (op *TracePlugin) Initialize(db *gorm.DB) (err error)
```

该插件会在 SQL 执行前后记录时间，并在执行完成后输出查询耗时，使用 klog 以 INFO 级别（V(4)）记录日志。

## 依赖注入

`ProviderSet` 变量提供了用于 Google Wire 依赖注入的提供者集合：

```go
var ProviderSet = wire.NewSet(
    NewMySQL,
)
```

目前该集合仅包含 MySQL 数据库连接创建器。如需包含 PostgreSQL 支持，可以扩展此集合。

## 安装

确保在你的项目中添加以下依赖：

```bash
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get gorm.io/driver/postgres
go get github.com/google/wire
go get k8s.io/klog/v2
```

## 使用示例

### MySQL 连接示例

```go
import (
    "github.com/moweilong/mo/gormx"
    "time"
)

// 创建 MySQL 配置
opts := &gormx.MySQLOptions{
    Addr:                  "localhost:3306",
    Username:              "root",
    Password:              "password",
    Database:              "testdb",
    MaxIdleConnections:    20,
    MaxOpenConnections:    100,
    MaxConnectionLifeTime: 30 * time.Second,
}

// 创建 MySQL 数据库连接
db, err := gormx.NewMySQL(opts)
if err != nil {
    // 处理错误
}

// 使用连接
var users []User
result := db.Find(&users)
```

### PostgreSQL 连接示例

```go
import (
    "github.com/moweilong/mo/gormx"
    "time"
)

// 创建 PostgreSQL 配置
opts := &gormx.PostgreSQLOptions{
    Addr:                  "localhost:5432",
    Username:              "postgres",
    Password:              "password",
    Database:              "testdb",
    MaxIdleConnections:    20,
    MaxOpenConnections:    100,
    MaxConnectionLifeTime: 30 * time.Second,
}

// 创建 PostgreSQL 数据库连接
pgdb, err := gormx.NewPostgreSQL(opts)
if err != nil {
    // 处理错误
}

// 使用连接
var products []Product
result := pgdb.Find(&products)
```

### 启用性能跟踪插件

```go
import (
    "github.com/moweilong/mo/gormx"
)

// 创建数据库连接
 db, err := gormx.NewMySQL(opts)
 if err != nil {
    // 处理错误
 }

// 注册性能跟踪插件
err = db.Use(&gormx.TracePlugin{})
if err != nil {
    // 处理错误
}
```

### 与 Google Wire 集成

```go
import (
    "github.com/google/wire"
    "github.com/moweilong/mo/gormx"
)

// 定义依赖注入提供者
var ProviderSet = wire.NewSet(
    // 提供数据库配置
    NewDatabaseConfig,
    // 包含 gormx 包的提供者
    gormx.ProviderSet,
)

// 创建数据库配置函数
func NewDatabaseConfig() *gormx.MySQLOptions {
    return &gormx.MySQLOptions{
        Addr:     "localhost:3306",
        Username: "root",
        Password: "password",
        Database: "testdb",
    }
}
```

## 注意事项

1. 确保在使用数据库连接前检查错误返回值
2. 对于生产环境，建议根据实际需求调整连接池配置参数
3. 性能跟踪插件使用 klog 记录日志，需要配置合适的日志级别才能看到输出
4. 当不再使用数据库连接时，建议调用 `db.Close()` 方法释放资源
5. PostgreSQL 连接默认设置 `sslmode=disable` 和 `TimeZone=Asia/Shanghai`，如有特殊需求可修改代码

## 依赖

- [gorm.io/gorm](https://gorm.io/) - GORM 数据库框架
- [gorm.io/driver/mysql](https://gorm.io/docs/connecting_to_the_database.html#MySQL) - MySQL 数据库驱动
- [gorm.io/driver/postgres](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL) - PostgreSQL 数据库驱动
- [github.com/google/wire](https://github.com/google/wire) - 依赖注入工具
- [k8s.io/klog/v2](https://github.com/kubernetes/klog) - 日志记录库