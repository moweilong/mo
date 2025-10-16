# version

version 是一个提供版本信息管理功能的工具包，主要用于管理和显示应用程序的版本信息，支持命令行标志和动态版本设置。

## 功能概述

该包主要实现了以下功能：
- 获取和管理应用程序的版本信息（Git版本、构建日期、Git提交ID等）
- 支持以多种格式输出版本信息（字符串、JSON、表格）
- 支持动态设置和验证版本信息
- 提供命令行标志支持（--version）
- 包含语义化版本解析和比较工具

## 文件说明

- `version.go`：定义了版本信息的核心结构体和基本操作
- `dynamic.go`：提供动态设置和验证版本的功能
- `flag.go`：提供命令行标志支持，用于处理--version参数
- `util/version.go`：提供语义化版本的解析和比较工具

## 核心类型和函数

### Info 结构体

```go
type Info struct {
    GitVersion   string `json:"gitVersion"`
    GitCommit    string `json:"gitCommit"`
    GitTreeState string `json:"gitTreeState"`
    BuildDate    string `json:"buildDate"`
    GoVersion    string `json:"goVersion"`
    Compiler     string `json:"compiler"`
    Platform     string `json:"platform"`
}
```

包含应用程序的完整版本信息。

### Get 函数

```go
func Get() Info
```

返回详尽的代码库版本信息，标明二进制文件由哪个版本的代码构建。

**返回值**：
- `Info`：包含完整版本信息的结构体

### Info 结构体方法

#### String

```go
func (info Info) String() string
```

返回人性化的版本信息字符串（仅返回GitVersion）。

#### ToJSON

```go
func (info Info) ToJSON() string
```

以JSON格式返回版本信息。

#### Text

```go
func (info Info) Text() string
```

将版本信息编码为UTF-8格式的表格文本并返回。

### 动态版本设置

#### SetDynamicVersion

```go
func SetDynamicVersion(dynamicVersion string)
```

覆盖Get()函数返回的GitVersion版本。指定的版本必须非空、是有效的语义化版本，并且必须匹配默认gitVersion的major/minor/patch版本。

**参数**：
- `dynamicVersion`：要设置的动态版本字符串

**返回值**：
- `error`：如果版本无效则返回错误

#### ValidateDynamicVersion

```go
func ValidateDynamicVersion(dynamicVersion string) error
```

验证给定的版本是否非空、是有效的语义化版本，并且匹配默认gitVersion的major/minor/patch版本。

**参数**：
- `dynamicVersion`：要验证的版本字符串

**返回值**：
- `error`：如果版本无效则返回错误

### 命令行标志支持

#### VersionVar

```go
func VersionVar(p *versionValue, name string, value versionValue, usage string)
```

定义一个具有指定名称和用法的版本标志。

#### Version

```go
func Version(name string, value versionValue, usage string) *versionValue
```

包装了VersionVar函数，创建并返回一个版本标志。

#### AddFlags

```go
func AddFlags(fs *pflag.FlagSet)
```

在任意FlagSet上注册这个包的标志，使它们指向与全局标志相同的值。

#### PrintAndExitIfRequested

```go
func PrintAndExitIfRequested()
```

检查是否传递了--version标志，如果是，则打印版本并退出程序。

## 使用示例

### 基本用法

```go
import (
    "fmt"
    "github.com/moweilong/mo/version"
)

func main() {
    // 获取版本信息
    v := version.Get()
    
    // 以不同格式输出版本信息
    fmt.Println("字符串格式:", v.String())
    fmt.Println("JSON格式:", v.ToJSON())
    fmt.Println("表格格式:\n", v.Text())
}
```

### 动态设置版本

```go
import (
    "fmt"
    "github.com/moweilong/mo/version"
)

func main() {
    // 尝试动态设置版本
    err := version.SetDynamicVersion("v1.2.3-custom")
    if err != nil {
        fmt.Printf("设置动态版本失败: %v\n", err)
    } else {
        fmt.Println("设置动态版本成功")
        fmt.Println("新的版本信息:", version.Get().String())
    }
}
```

### 命令行标志支持

```go
import (
    "github.com/moweilong/mo/version"
    "github.com/spf13/pflag"
)

func main() {
    // 创建自定义FlagSet
    flags := pflag.NewFlagSet("myapp", pflag.ExitOnError)
    
    // 添加版本标志
    version.AddFlags(flags)
    
    // 解析命令行参数
    flags.Parse(os.Args[1:])
    
    // 检查是否请求了版本信息
    version.PrintAndExitIfRequested()
    
    // 程序的主要逻辑...
}
```

### 使用语义化版本工具

```go
import (
    "fmt"
    "github.com/moweilong/mo/version/util"
)

func main() {
    // 解析语义化版本
    v1, err := util.ParseSemantic("v1.2.3-alpha+build.123")
    if err != nil {
        fmt.Printf("解析版本失败: %v\n", err)
        return
    }
    
    v2, err := util.ParseSemantic("v1.2.4")
    if err != nil {
        fmt.Printf("解析版本失败: %v\n", err)
        return
    }
    
    // 比较版本
    if v1.LessThan(v2) {
        fmt.Printf("%v 小于 %v\n", v1, v2)
    }
    
    // 获取版本组件
    fmt.Printf("主要版本: %d\n", v1.Major())
    fmt.Printf("次要版本: %d\n", v1.Minor())
    fmt.Printf("补丁版本: %d\n", v1.Patch())
    fmt.Printf("预发布版本: %s\n", v1.PreRelease())
    fmt.Printf("构建元数据: %s\n", v1.BuildMetadata())
}
```

## 注意事项

1. 版本信息通常在构建时通过 `-ldflags` 参数设置，例如：
   ```bash
   go build -ldflags "-X github.com/moweilong/mo/version.gitVersion=v1.0.0 -X github.com/moweilong/mo/version.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ') -X github.com/moweilong/mo/version.gitCommit=$(git rev-parse HEAD) -X github.com/moweilong/mo/version.gitTreeState=clean"
   ```

2. 动态设置版本时，必须遵循语义化版本规则，并且必须匹配默认版本的 major/minor/patch 版本号。

3. 命令行标志支持三种模式：
   - `--version=false`（默认）：不显示版本信息
   - `--version` 或 `--version=true`：显示简短的版本信息
   - `--version=raw`：显示详细的版本信息表格

4. util包中的语义化版本工具基于Kubernetes的版本解析实现，支持完整的语义化版本规范。

## 依赖

- `github.com/gosuri/uitable`：用于格式化版本信息表格输出
- `github.com/spf13/pflag`：用于命令行标志支持
- Go标准库：encoding/json, fmt, os, regexp, runtime, strconv, strings, sync/atomic

## 许可证

util目录下的代码基于Apache License 2.0许可，源自Kubernetes项目。