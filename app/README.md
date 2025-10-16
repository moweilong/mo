# app 包

`app` 包是一个功能完备的命令行应用程序框架，基于 Cobra 构建，提供了配置管理、日志设置、命令行解析等基础功能，可以用于快速构建一个功能完备的命令行应用程序。

## 功能概述

该包主要提供了以下功能：
- 统一的应用程序结构定义和生命周期管理
- 命令行参数解析和验证
- 配置文件加载和监控（支持多种格式：JSON、TOML、YAML、HCL 等）
- 环境变量自动映射
- 日志系统初始化和配置
- 健康检查服务集成
- 应用程序版本信息管理

## 文件说明

- **app.go**: 定义应用程序的核心结构和主要功能
- **config.go**: 提供配置文件加载和管理功能
- **help.go**: 实现帮助命令和帮助信息显示
- **options.go**: 定义选项验证和标志集相关接口

## 核心结构和接口

### App 结构体

```go
// App is the main structure of a cli application.
type App struct {
    name        string
    shortDesc   string
    description string
    run         RunFunc
    cmd         *cobra.Command
    args        cobra.PositionalArgs
    healthCheckFunc HealthCheckFunc
    options     any
    silence     bool
    noConfig    bool
    watch       bool
    contextExtractors map[string]func(context.Context) string
}
```

### 核心接口

```go
// RunFunc defines the application's startup callback function.
type RunFunc func() error

// HealthCheckFunc defines the health check function for the application.
type HealthCheckFunc func() error

// Option defines optional parameters for initializing the application structure.
type Option func(*App)

// OptionsValidator provides methods to complete and validate options.
type OptionsValidator interface {
    Complete() error
    Validate() error
}

// NamedFlagSetOptions provides access to server-specific flag sets.
type NamedFlagSetOptions interface {
    Flags() cliflag.NamedFlagSets
    OptionsValidator
}

// FlagSetOptions defines an interface for command-line options.
type FlagSetOptions interface {
    AddFlags(fs *pflag.FlagSet)
    OptionsValidator
}
```

## 主要函数

### 应用程序创建和配置

```go
// NewApp creates a new application instance.
func NewApp(name string, shortDesc string, opts ...Option) *App

// Run is used to launch the application.
func (app *App) Run()

// Command returns cobra command instance inside the application.
func (app *App) Command() *cobra.Command
```

### 配置管理

```go
// AddConfigFlag adds flags for a specific server to the specified FlagSet object.
func AddConfigFlag(fs *pflag.FlagSet, name string, watch bool)

// PrintConfig prints all configuration values.
func PrintConfig()
```

### 帮助功能

```go
// helpCommand creates a help command.
func helpCommand(name string) *cobra.Command

// addHelpFlag adds help flags for an application.
func addHelpFlag(name string, fs *pflag.FlagSet)

// addHelpCommandFlag adds help flags for a specific command.
func addHelpCommandFlag(usage string, fs *pflag.FlagSet)
```

## 应用选项

`app` 包提供了一系列 Option 函数来配置应用程序的行为：

```go
// WithOptions 启用从命令行或配置文件读取参数的功能
func WithOptions(opts any) Option

// WithRunFunc 设置应用程序启动回调函数
func WithRunFunc(run RunFunc) Option

// WithDescription 设置应用程序的描述
func WithDescription(desc string) Option

// WithHealthCheckFunc 设置应用程序的健康检查函数
func WithHealthCheckFunc(fn HealthCheckFunc) Option

// WithDefaultHealthCheckFunc 设置默认的健康检查函数
func WithDefaultHealthCheckFunc() Option

// WithSilence 设置应用程序为静默模式
func WithSilence() Option

// WithNoConfig 设置应用程序不提供配置标志
func WithNoConfig() Option

// WithValidArgs 设置非标志参数的验证函数
func WithValidArgs(args cobra.PositionalArgs) Option

// WithDefaultValidArgs 设置默认的非标志参数验证函数
func WithDefaultValidArgs() Option

// WithWatchConfig 启用配置文件监控和重新读取
func WithWatchConfig() Option

// WithLoggerContextExtractor 设置日志上下文提取器
func WithLoggerContextExtractor(contextExtractors map[string]func(context.Context) string) Option
```

## 使用示例

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/moweilong/mo/app"
)

func main() {
    // 创建应用程序
    myapp := app.NewApp("myapp", "A simple CLI application",
        app.WithRunFunc(func() error {
            fmt.Println("Hello, World!")
            return nil
        }),
    )
    
    // 运行应用程序
    myapp.Run()
}
```

### 带配置选项的用法

```go
package main

import (
    "github.com/moweilong/mo/app"
    "github.com/moweilong/mo/options"
)

// 定义自定义选项
type MyOptions struct {
    Server *options.HTTPOptions
    Log    *options.LogOptions
}

// 实现 OptionsValidator 接口
func (o *MyOptions) Complete() error {
    return nil
}

func (o *MyOptions) Validate() error {
    if err := o.Server.Validate(); err != nil {
        return err
    }
    return nil
}

// 实现 FlagSetOptions 接口
func (o *MyOptions) AddFlags(fs *pflag.FlagSet) {
    o.Server.AddFlags(fs)
    o.Log.AddFlags(fs)
}

func main() {
    // 创建选项实例
    opts := &MyOptions{
        Server: options.NewHTTPOptions(),
        Log:    options.NewLogOptions(),
    }
    
    // 创建应用程序
    myapp := app.NewApp("myapp", "A CLI application with options",
        app.WithOptions(opts),
        app.WithRunFunc(func() error {
            // 使用配置选项
            fmt.Printf("Server address: %s\n", opts.Server.Addr)
            fmt.Printf("Log level: %s\n", opts.Log.Level)
            return nil
        }),
        app.WithDefaultValidArgs(),
    )
    
    // 运行应用程序
    myapp.Run()
}
```

### 带健康检查的用法

```go
package main

import (
    "github.com/moweilong/mo/app"
    "github.com/moweilong/mo/log"
)

func main() {
    // 创建应用程序
    myapp := app.NewApp("myapp", "A CLI application with health check",
        app.WithRunFunc(func() error {
            // 应用程序主逻辑
            return nil
        }),
        app.WithDefaultHealthCheckFunc(), // 使用默认健康检查
    )
    
    // 运行应用程序
    myapp.Run()
}
```

### 带配置文件监控的用法

```go
package main

import (
    "github.com/moweilong/mo/app"
)

func main() {
    // 创建应用程序
    myapp := app.NewApp("myapp", "A CLI application with config watch",
        app.WithRunFunc(func() error {
            // 应用程序主逻辑
            return nil
        }),
        app.WithWatchConfig(), // 启用配置文件监控
    )
    
    // 运行应用程序
    myapp.Run()
}
```

## 配置文件

`app` 包支持多种格式的配置文件（JSON、TOML、YAML、HCL 等），配置文件可以通过 `-c` 或 `--config` 标志指定，也可以放在以下位置：
- 当前目录
- `~/.应用名/` 目录
- `/etc/应用名/` 目录

配置文件的名称默认为应用程序的名称。

## 环境变量

`app` 包会自动将环境变量映射到配置中，环境变量的命名规则是：
- 前缀为应用程序名称的大写形式（连字符 `-` 替换为下划线 `_`）
- 配置键中的点 `.` 和连字符 `-` 也替换为下划线 `_`

例如，配置键 `server.addr` 对应的环境变量为 `MYAPP_SERVER_ADDR`（假设应用程序名称为 myapp）。

## 注意事项

1. 建议使用 `app.NewApp()` 函数创建应用程序实例
2. 应用程序的运行逻辑应该通过 `WithRunFunc` 选项设置
3. 配置选项应该实现 `OptionsValidator` 接口以支持验证
4. 可以通过 `WithSilence()` 选项禁用启动信息和配置信息的打印
5. 可以通过 `WithNoConfig()` 选项禁用配置文件功能
6. 日志系统会自动根据配置进行初始化

## 依赖

- Cobra: `github.com/spf13/cobra` - 命令行框架
- Viper: `github.com/spf13/viper` - 配置管理
- Kubernetes 组件基础: `k8s.io/component-base` - 提供 CLI 工具和标志支持
- 项目内部: `github.com/moweilong/mo/log`, `github.com/moweilong/mo/options`, `github.com/moweilong/mo/version`