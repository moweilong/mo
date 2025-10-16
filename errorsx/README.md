# errorsx

errorsx 是一个增强型错误处理包，提供了结构化、可扩展的错误定义和处理机制，特别适合需要在微服务架构中传递详细错误信息的场景。该包支持 HTTP 和 gRPC 协议的错误转换，并提供了丰富的错误元数据管理功能。

## 核心功能

- 结构化错误定义，包含错误码、错误原因和详细信息
- 支持错误元数据管理，可以附加额外的上下文信息
- 与 HTTP 和 gRPC 协议的错误转换互操作
- 提供预定义的标准错误类型
- 支持错误链处理（Is、As、Unwrap）

## 核心类型

### ErrorX 结构体

```go
// ErrorX 定义了 OneX 项目体系中使用的错误类型，用于描述错误的详细信息.
type ErrorX struct {
    // Code 表示错误的 HTTP 状态码，用于与客户端进行交互时标识错误的类型.
    Code int `json:"code,omitempty"`

    // Reason 表示错误发生的原因，通常为业务错误码，用于精准定位问题.
    Reason string `json:"reason,omitempty"`

    // Message 表示简短的错误信息，通常可直接暴露给用户查看.
    Message string `json:"message,omitempty"`

    // Metadata 用于存储与该错误相关的额外元信息，可以包含上下文或调试信息.
    Metadata map[string]string `json:"metadata,omitempty"`
}
```

## 主要函数和方法

### 创建错误

```go
// New 创建一个新的错误.
func New(code int, reason string, format string, args ...any) *ErrorX
```

### 错误方法

```go
// Error 实现 error 接口中的 `Error` 方法.
func (err *ErrorX) Error() string

// WithMessage 设置错误的 Message 字段.
func (err *ErrorX) WithMessage(format string, args ...any) *ErrorX

// WithMetadata 设置元数据.
func (err *ErrorX) WithMetadata(md map[string]string) *ErrorX

// KV 使用 key-value 对设置元数据.
func (err *ErrorX) KV(kvs ...string) *ErrorX

// WithRequestID 设置请求 ID.
func (err *ErrorX) WithRequestID(requestID string) *ErrorX

// Is 判断当前错误是否与目标错误匹配.
func (err *ErrorX) Is(target error) bool

// GRPCStatus 返回 gRPC 状态表示.
func (err *ErrorX) GRPCStatus() *status.Status
```

### 错误工具函数

```go
// FromError 尝试将一个通用的 error 转换为自定义的 *ErrorX 类型.
func FromError(err error) *ErrorX

// Code 返回错误的 HTTP 代码.
func Code(err error) int

// Reason 返回特定错误的原因.
func Reason(err error) string

// Is 报告 err 链中是否有任何错误与目标匹配.
func Is(err, target error) bool

// As 在 err 的链中查找第一个与目标匹配的错误.
func As(err error, target interface{}) bool

// Unwrap 返回在 err 上调用 Unwrap 方法的结果.
func Unwrap(err error) error
```

## 预定义错误

```go
var (
    // OK 代表请求成功.
    OK = &ErrorX{Code: http.StatusOK, Message: ""}

    // ErrInternal 表示所有未知的服务器端错误.
    ErrInternal = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError", Message: "Internal server error."}

    // ErrNotFound 表示资源未找到.
    ErrNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound", Message: "Resource not found."}

    // ErrBind 表示请求体绑定错误.
    ErrBind = &ErrorX{Code: http.StatusBadRequest, Reason: "BindError", Message: "Error occurred while binding the request body to the struct."}

    // ErrInvalidArgument 表示参数验证失败.
    ErrInvalidArgument = &ErrorX{Code: http.StatusBadRequest, Reason: "InvalidArgument", Message: "Argument verification failed."}

    // ErrUnauthenticated 表示认证失败.
    ErrUnauthenticated = &ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthenticated", Message: "Unauthenticated."}

    // ErrPermissionDenied 表示请求没有权限.
    ErrPermissionDenied = &ErrorX{Code: http.StatusForbidden, Reason: "PermissionDenied", Message: "Permission denied. Access to the requested resource is forbidden."}

    // ErrOperationFailed 表示操作失败.
    ErrOperationFailed = &ErrorX{Code: http.StatusConflict, Reason: "OperationFailed", Message: "The requested operation has failed. Please try again later."}
)
```

## 使用示例

### 创建和使用基本错误

```go
import (
    "net/http"
    "github.com/moweilong/mo/errorsx"
)

// 创建一个新错误
err := errorsx.New(http.StatusBadRequest, "InvalidInput", "Invalid input for field %s", "username")

// 添加元数据
err = err.WithMetadata(map[string]string{
    "field": "username",
    "validation": "required",
})

// 或者使用键值对方式添加元数据
err = err.KV("field", "username", "validation", "required")

// 添加请求ID
err = err.WithRequestID("req-12345")

// 更改错误消息
err = err.WithMessage("新的错误消息")
```

### 错误转换

```go
// 将普通错误转换为 ErrorX
plainErr := errors.New("Something went wrong")
errx := errorsx.FromError(plainErr) // 将返回一个包含默认值的 ErrorX

// 获取错误码
code := errorsx.Code(errx) // 返回 HTTP 状态码

// 获取错误原因
reason := errorsx.Reason(errx) // 返回错误原因

// 错误比较
if errorsx.Is(err, errorsx.ErrNotFound) {
    // 处理资源未找到的情况
}
```

### gRPC 错误互操作

```go
// ErrorX 转换为 gRPC 错误
status := errx.GRPCStatus()
grpcErr := status.Err()

// gRPC 错误转换为 ErrorX
errx := errorsx.FromError(grpcErr)
```

## 注意事项

1. 错误比较（Is 方法）基于 Code 和 Reason 字段，不考虑 Message 和 Metadata
2. 当将普通错误转换为 ErrorX 时，默认使用 ErrInternal 的 Code 和 Reason
3. 预定义错误是全局变量，使用时应避免直接修改其字段值
4. WithMetadata 方法会替换整个元数据映射，而 KV 方法会添加或更新键值对
5. 在微服务间传递错误时，建议使用 FromError 函数确保错误信息正确转换