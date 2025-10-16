# validation - 验证工具包

## 功能概述

validation 是一个提供请求参数和结构体字段验证功能的Go工具包。该包主要实现了两种验证模式：

1. 基于类型的验证器 - 自动注册和调用与请求类型匹配的验证方法
2. 基于规则的验证器 - 通过定义验证规则对结构体字段进行验证

## 目录结构

```
validation/
├── validation.go    # 基于类型的验证器实现
└── validator.go     # 基于规则的验证器实现
```

## 核心接口和结构体

### 1. Validator 结构体

```go
// Validator implements the validate.IValidator interface.
type Validator struct {
    registry map[string]reflect.Value
}
```

`Validator` 是基于类型的验证器实现，它维护了一个验证方法注册表，根据请求类型自动选择匹配的验证方法。

### 2. 验证函数类型

```go
// 定义验证函数类型
type ValidatorFunc func(value any) error
```

`ValidatorFunc` 是一个函数类型，用于定义对单个值进行验证的函数。

### 3. 验证规则类型

```go
// 定义验证规则类型
type Rules map[string]ValidatorFunc
```

`Rules` 是一个映射类型，用于将字段名映射到对应的验证函数。

## 主要函数

### 1. 基于类型的验证

#### NewValidator

```go
// NewValidator creates and initializes a custom validator.
func NewValidator(customValidator any) *Validator
```

创建一个新的验证器实例，自动从提供的自定义验证器中提取验证方法。

**参数**：
- `customValidator` - 包含验证方法的自定义验证器实例

**返回值**：
- 初始化后的 `Validator` 实例

#### Validate

```go
// Validate validates the request using the appropriate validation method.
func (v *Validator) Validate(ctx context.Context, request any) error
```

根据请求类型调用匹配的验证方法进行验证。

**参数**：
- `ctx` - 上下文对象
- `request` - 要验证的请求对象

**返回值**：
- 如果验证失败，返回错误；如果验证成功或没有找到匹配的验证方法，返回 nil

#### ValidRequired

```go
// ValidRequired 验证结构体中的必需字段是否存在且不为空.
func ValidRequired(obj any, requiredFields ...string) error
```

验证结构体中的必需字段是否存在且不为空。

**参数**：
- `obj` - 要验证的结构体或结构体指针
- `requiredFields` - 需要验证的必需字段列表

**返回值**：
- 如果验证失败，返回错误；如果验证成功，返回 nil

### 2. 基于规则的验证

#### ValidateAllFields

```go
func ValidateAllFields(obj any, rules Rules) error
```

使用指定的验证规则验证结构体的所有字段。

**参数**：
- `obj` - 要验证的结构体或结构体指针
- `rules` - 验证规则映射

**返回值**：
- 如果验证失败，返回错误；如果验证成功，返回 nil

#### ValidateSelectedFields

```go
func ValidateSelectedFields(obj any, rules Rules, fields ...string) error
```

使用指定的验证规则验证结构体的选定字段。

**参数**：
- `obj` - 要验证的结构体或结构体指针
- `rules` - 验证规则映射
- `fields` - 要验证的字段列表

**返回值**：
- 如果验证失败，返回错误；如果验证成功，返回 nil

#### GetExportedFieldNames

```go
// GetExportedFieldNames 返回传入结构体中所有可导出的字段名字.
func GetExportedFieldNames(obj any) []string
```

获取结构体中所有可导出的字段名（字段名以大写字母开头）。

**参数**：
- `obj` - 结构体或结构体指针

**返回值**：
- 可导出字段名的切片

## 使用示例

### 1. 基于类型的验证示例

```go
import (
    "context"
    "errors"
    "github.com/moweilong/mo/validation"
)

// 定义自定义验证器
type CustomValidator struct{}

// 定义请求结构体
type CreateUserRequest struct {
    Username string
    Email    string
}

// 为 CreateUserRequest 实现验证方法
func (v *CustomValidator) ValidateCreateUserRequest(ctx context.Context, req *CreateUserRequest) error {
    if req.Username == "" {
        return errors.New("username is required")
    }
    if req.Email == "" {
        return errors.New("email is required")
    }
    return nil
}

// 创建验证器并使用
func main() {
    validator := validation.NewValidator(&CustomValidator{})
    req := &CreateUserRequest{Username: "user123", Email: "user@example.com"}
    err := validator.Validate(context.Background(), req)
    if err != nil {
        // 处理验证错误
    }
}
```

### 2. 基于规则的验证示例

```go
import (
    "errors"
    "github.com/moweilong/mo/validation"
)

// 定义用户结构体
type User struct {
    Username string
    Age      int
    Email    string
}

// 创建验证规则
func main() {
    rules := validation.Rules{
        "Username": func(val any) error {
            username := val.(string)
            if username == "" {
                return errors.New("username cannot be empty")
            }
            if len(username) < 3 {
                return errors.New("username must be at least 3 characters long")
            }
            return nil
        },
        "Age": func(val any) error {
            age := val.(int)
            if age < 18 {
                return errors.New("user must be at least 18 years old")
            }
            return nil
        },
    }

    user := &User{Username: "user", Age: 16, Email: "user@example.com"}
    err := validation.ValidateAllFields(user, rules)
    if err != nil {
        // 处理验证错误：Age is less than 18
    }
}
```

### 3. 使用 ValidRequired 验证必需字段

```go
import (
    "github.com/moweilong/mo/validation"
)

// 定义请求结构体
type LoginRequest struct {
    Username *string
    Password *string
}

func main() {
    username := "admin"
    req := &LoginRequest{Username: &username}
    
    // 验证 Username 和 Password 字段是否存在且不为空
    err := validation.ValidRequired(req, "Username", "Password")
    if err != nil {
        // 处理验证错误：Password must be provided
    }
}
```

## 特性和优势

1. **灵活的验证方式**：支持基于类型和基于规则两种验证方式，可以根据需求选择合适的验证方式
2. **自动注册验证方法**：基于类型的验证器可以自动从自定义验证器中提取验证方法
3. **反射支持**：使用Go的反射机制实现对任意结构体的字段验证
4. **可扩展的验证规则**：可以根据需要定义任意的验证函数
5. **支持必需字段验证**：提供了专门的函数验证结构体中的必需字段

## 注意事项

1. 基于类型的验证器要求验证方法名称遵循 `Validate{RequestTypeName}` 的命名约定
2. 验证方法必须接受 `context.Context` 和请求结构体指针作为参数，并返回 `error` 类型
3. 对于指针类型的字段，如果字段值为 nil，基于规则的验证器会跳过对该字段的验证
4. 基于规则的验证器只会验证结构体中存在且可导出的字段