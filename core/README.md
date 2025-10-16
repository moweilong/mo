# core

core 包提供了项目的核心功能组件，包括请求处理、配置管理和对象复制等基础工具，用于简化 API 开发中的常见任务。

## 核心功能

- 请求参数绑定、验证和处理
- 统一的错误响应格式化
- 配置文件和环境变量管理
- 对象之间的高效复制（含类型转换支持）

## 目录结构

```
core/
├── core.go      # 核心请求处理功能
├── config.go    # 配置管理功能
└── copier.go    # 对象复制功能
```

## 核心类型

### 1. 请求处理相关类型

```go
// Validator 是验证函数的类型，用于对绑定的数据结构进行验证.
type Validator[T any] func(context.Context, *T) error

// Binder 定义绑定函数的类型，用于绑定请求数据到相应结构体.
type Binder func(any) error

// Handler 是处理函数的类型，用于处理已经绑定和验证的数据.
type Handler[T any, R any] func(ctx context.Context, req *T) (R, error)

// ErrorResponse 定义了错误响应的结构，用于 API 请求中发生错误时返回统一的格式化错误信息.
type ErrorResponse struct {
    Reason string `json:"reason,omitempty"`
    Message string `json:"message,omitempty"`
    Metadata map[string]string `json:"metadata,omitempty"`
}
```

## 主要函数

### 1. 请求处理函数

```go
// HandleJSONRequest 处理 JSON 请求的快捷函数.
func HandleJSONRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T])

// HandleQueryRequest 处理 Query 参数请求的快捷函数.
func HandleQueryRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T])

// HandleUriRequest 处理 URI 请求的快捷函数.
func HandleUriRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T])

// HandleRequest 通用的请求处理函数.
func HandleRequest[T any, R any](c *gin.Context, binder Binder, handler Handler[T, R], validators ...Validator[T])
```

### 2. 请求参数绑定函数

```go
// ShouldBindJSON 使用 JSON 格式的绑定函数绑定请求参数并执行验证.
func ShouldBindJSON[T any](c *gin.Context, rq *T, validators ...Validator[T]) error

// ShouldBindQuery 使用 Query 格式的绑定函数绑定请求参数并执行验证.
func ShouldBindQuery[T any](c *gin.Context, rq *T, validators ...Validator[T]) error

// ShouldBindUri 使用 URI 格式的绑定函数绑定请求参数并执行验证.
func ShouldBindUri[T any](c *gin.Context, rq *T, validators ...Validator[T]) error

// ReadRequest 用于绑定和验证请求数据的通用工具函数.
func ReadRequest[T any](c *gin.Context, rq *T, binder Binder, validators ...Validator[T]) error
```

### 3. 响应处理函数

```go
// WriteResponse 通用的响应函数，根据是否发生错误生成成功响应或标准化的错误响应.
func WriteResponse(c *gin.Context, data any, err error)
```

### 4. 配置管理函数

```go
// OnInitialize 设置需要读取的配置文件名、环境变量，并将其内容读取到 viper 中.
func OnInitialize(configFile *string, envPrefix string, loadDirs []string, defaultConfigName string) func()
```

### 5. 对象复制函数

```go
// TypeConverters 定义时间类型转换器，用于 copier 的深度拷贝.
func TypeConverters() []copier.TypeConverter

// CopyWithConverters 执行带类型转换器的对象复制.
func CopyWithConverters(to any, from any) error

// Copy 执行简单的对象复制.
func Copy(to any, from any) error
```

## 使用示例

### 1. 请求处理示例

```go
import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/moweilong/mo/core"
    "github.com/moweilong/mo/errorsx"
)

// 定义请求和响应结构体
type UserRequest struct {
    ID   int64  `json:"id" binding:"required,gt=0"`
    Name string `json:"name" binding:"required,min=2,max=20"`
}

type UserResponse struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 自定义验证函数
func validateUserRequest(ctx context.Context, req *UserRequest) error {
    if req.Name == "admin" {
        return errorsx.ErrInvalidArgument.WithMessage("名称不能为admin")
    }
    return nil
}

// 业务处理函数
func getUserHandler(ctx context.Context, req *UserRequest) (UserResponse, error) {
    // 实际业务逻辑
    return UserResponse{
        ID:   req.ID,
        Name: req.Name,
        Age:  30,
    }, nil
}

func main() {
    r := gin.Default()
    
    // 使用 HandleJSONRequest 处理 JSON 请求
    r.POST("/users", func(c *gin.Context) {
        core.HandleJSONRequest(c, getUserHandler, validateUserRequest)
    })
    
    // 启动服务器
    r.Run(":8080")
}
```

### 2. 配置初始化示例

```go
import (
    "github.com/moweilong/mo/core"
)

func main() {
    var configFile string
    
    // 创建配置初始化函数
    initConfig := core.OnInitialize(
        &configFile,          // 配置文件路径，如果为nil则从默认路径搜索
        "APP",                // 环境变量前缀
        []string{".", "/etc/app"},  // 配置文件搜索路径
        "config",             // 默认配置文件名
    )
    
    // 执行配置初始化
    initConfig()
    
    // 现在可以通过 viper 获取配置值了
    // value := viper.GetString("key")
}
```

### 3. 对象复制示例

```go
import (
    "time"
    "github.com/moweilong/mo/core"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}

type UserDTO struct {
    ID        int64              `json:"id"`
    Name      string             `json:"name"`
    CreatedAt *timestamppb.Timestamp `json:"created_at"`
}

func main() {
    user := User{
        ID:        1,
        Name:      "张三",
        CreatedAt: time.Now(),
    }
    
    var userDTO UserDTO
    
    // 使用带类型转换器的复制函数（支持 time.Time 和 timestamppb.Timestamp 转换）
    err := core.CopyWithConverters(&userDTO, &user)
    if err != nil {
        // 处理错误
    }
    
    // 或者使用简单复制（不支持自定义类型转换）
    // err := core.Copy(&userDTO, &user)
}
```

## 注意事项

1. 请求处理函数使用了泛型，可以处理不同类型的请求和响应
2. ReadRequest 函数会自动检测请求结构体是否实现了 Default() 方法，如果实现了则会调用该方法设置默认值
3. WriteResponse 函数会自动将 errorsx.ErrorX 类型的错误转换为标准化的错误响应
4. CopyWithConverters 函数支持 time.Time 和 timestamppb.Timestamp 之间的自动转换
5. 配置管理基于 viper 库，支持多种配置格式和环境变量覆盖