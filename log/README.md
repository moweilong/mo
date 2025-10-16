# log 包

log 包是一个基于 zap 日志库封装的日志模块，提供了统一的日志记录接口，并集成了 GORM 和 Kratos 框架的日志系统。

## 功能概述

- 支持多级别日志记录（Debug、Info、Warn、Error、Panic、Fatal）
- 支持格式化输出（console、json）
- 支持上下文（context）信息提取
- 集成 GORM 框架日志系统
- 集成 Kratos 框架日志系统
- 支持自定义配置选项

## 文件说明

| 文件名 | 说明 |
|-------|------|
| log.go | 核心日志接口定义和实现 |
| options.go | 日志配置选项定义 |
| gorm.go | GORM 框架日志接口实现 |
| kratos.go | Kratos 框架日志接口实现 |
| logger/kratos/kratos.go | Kratos 服务日志记录器创建 |
| logger/store/logger.go | 存储层错误日志记录器 |

## 配置选项

Options 结构体提供了以下配置选项：

```go
type Options struct {
    // DisableCaller 指定是否在日志中包含调用者信息
    DisableCaller bool
    // DisableStacktrace 指定是否禁止记录 panic 及以上级别的堆栈信息
    DisableStacktrace bool
    // EnableColor 指定是否输出彩色日志
    EnableColor bool
    // Level 指定最低日志级别，有效值：debug、info、warn、error、dpanic、panic、fatal
    Level string
    // Format 指定日志输出格式，有效值：console、json
    Format string
    // OutputPaths 指定日志输出路径
    OutputPaths []string
}
```

## 主要函数

### 创建日志记录器

```go
// NewLogger 根据传入的选项创建日志记录器
func NewLogger(opts *Options, options ...Option) *zapLogger

// NewOptions 创建带有默认值的配置选项
func NewOptions() *Options
```

### 全局日志记录器

```go
// Default 返回全局日志记录器
func Default() Logger

// Init 使用指定选项初始化全局日志记录器
func Init(opts *Options, options ...Option)
```

### 日志记录方法

支持多种日志记录方法，包括：

```go
// 格式化日志记录
Debugf(format string, args ...any)
Infof(format string, args ...any)
Warnf(format string, args ...any)
Errorf(format string, args ...any)
Panicf(format string, args ...any)
Fatalf(format string, args ...any)

// 结构化日志记录
Debugw(msg string, keyvals ...any)
Infow(msg string, keyvals ...any)
Warnw(msg string, keyvals ...any)
Errorw(err error, msg string, keyvals ...any)
Panicw(msg string, keyvals ...any)
Fatalw(msg string, keyvals ...any)
```

### 上下文相关

```go
// W 返回带有上下文的日志记录器
W(ctx context.Context) Logger

// WithContextExtractor 添加自定义上下文提取逻辑
func WithContextExtractor(contextExtractors ContextExtractors) Option
```

## 集成功能

### GORM 集成

log 包实现了 `gormlogger.Interface` 接口，可以直接用于 GORM 框架的日志配置：

```go
import (
    "gorm.io/gorm"
    "github.com/moweilong/mo/log"
)

func setupDB() {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: log.Default(), // 使用 log 包的日志记录器
    })
}
```

### Kratos 集成

log 包实现了 `krtlog.Logger` 接口，可以用于 Kratos 框架的日志配置：

```go
import (
    "github.com/go-kratos/kratos/v2"
    "github.com/moweilong/mo/log"
    "github.com/moweilong/mo/log/logger/kratos"
)

func newApp() *kratos.App {
    // 创建带有服务信息的日志记录器
    logger := kratos.NewLogger("service-id", "service-name", "v1.0.0")
    
    return kratos.New(
        kratos.Logger(logger),
        // 其他选项
    )
}
```

## 使用示例

### 基本使用

```go
import "github.com/moweilong/mo/log"

// 使用全局日志记录器
log.Infof("Hello, %s!", "world")
log.Errorw(err, "operation failed")

// 带字段的结构化日志
log.Infow("user login", "user_id", 123, "ip", "192.168.1.1")

// 带上下文的日志
log.W(ctx).Infow("processing request")
```

### 自定义配置

```go
import "github.com/moweilong/mo/log"

// 创建自定义配置
opts := log.NewOptions()
opts.Level = "debug"
opts.Format = "json"
opts.OutputPaths = []string{"stdout", "/var/log/app.log"}
opts.EnableColor = true

// 初始化全局日志记录器
log.Init(opts)

// 使用自定义的日志记录器
logger := log.NewLogger(opts)
logger.Debugf("Debug message")
```

### 上下文提取器

```go
import (
    "context"
    "github.com/moweilong/mo/log"
)

// 定义上下文提取器
extractors := log.ContextExtractors{
    "request_id": func(ctx context.Context) string {
        if rid, ok := ctx.Value("request_id").(string); ok {
            return rid
        }
        return ""
    },
    "user_id": func(ctx context.Context) string {
        if uid, ok := ctx.Value("user_id").(string); ok {
            return uid
        }
        return ""
    },
}

// 创建带有上下文提取器的日志记录器
logger := log.NewLogger(log.NewOptions(), log.WithContextExtractor(extractors))

// 使用
ctx := context.WithValue(context.Background(), "request_id", "req-123")
logger.W(ctx).Infow("processing request") // 日志会自动包含 request_id 字段
```

## 存储层日志

log 包提供了一个简单的存储层日志记录器，可以方便地记录存储操作中的错误：

```go
import (
    "context"
    "github.com/moweilong/mo/log/logger/store"
)

// 创建存储日志记录器
logger := store.NewLogger()

// 记录错误
ctx := context.Background()
err := someStorageOperation()
if err != nil {
    logger.Error(ctx, err, "storage operation failed", "operation", "read")
}
```

## 注意事项

1. 在应用程序退出前，建议调用 `log.Sync()` 来确保所有日志都被刷新到输出目标。
2. 对于性能敏感的应用，可以考虑关闭调用者信息记录（`DisableCaller: true`）。
3. 在生产环境中，建议使用 JSON 格式输出日志，便于日志收集和分析。
4. 对于长时间运行的服务，建议将日志输出到文件，并配置日志轮转。

## 依赖说明

- github.com/go-kratos/kratos/v2/log - Kratos 日志接口
- go.uber.org/zap - 底层日志库
- gorm.io/gorm/logger - GORM 日志接口
- github.com/spf13/pflag - 命令行参数解析