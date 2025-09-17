# mlog

mlog 是一个基于 zap 日志库的日志库，提供了简单的日志记录功能。

## 安装

```bash
go get github.com/moweilong/mo/pkg/mlog
```

## 使用

```go
import (
	"context"
	"log"
	"os"

	"github.com/moweilong/mo/pkg/mlog"
)

func main() {
	// 初始化日志记录器
	mlog.Init(&mlog.Options{
		Level:  "debug",
		Format: "console",
	})

	// 记录日志
	mlog.Info("这是一条 info 日志")
	mlog.Debug("这是一条 debug 日志")
}
```

## 配置

mlog 支持通过 `Options` 结构体配置日志记录器，包括日志级别、日志格式、日志输出位置等。

## 日志格式

mlog 支持两种日志格式：console 和 json。console 格式输出可读性较高的日志，json 格式输出符合 JSON 格式的日志。

