# rid - 资源标识符生成工具包

## 功能概述

rid 是一个提供资源标识符（Resource ID）生成和管理功能的Go工具包。该包主要用于创建带前缀的唯一标识符，适用于分布式系统中各种资源的唯一标识管理。

## 目录结构

```
rid/
├── rid.go          # 核心资源标识符实现
├── salt.go         # 机器ID和盐值生成
├── example_test.go # 使用示例
└── rid_test.go     # 单元测试
```

## 核心类型和函数

### 1. ResourceID 类型

```go
// ResourceID 表示资源的唯一标识符
type ResourceID string
```

`ResourceID` 是一个字符串类型的别名，用于表示不同类型资源的标识符前缀。

### 2. 主要函数和方法

#### NewResourceID

```go
// NewResourceID 创建一个新的资源标识符
func NewResourceID(prefix string) ResourceID
```

创建一个新的资源标识符实例。

**参数**：
- `prefix` - 资源标识符的前缀字符串

**返回值**：
- 初始化后的 `ResourceID` 实例

#### String

```go
// String 实现 Stringer 接口
func (rid ResourceID) String() string
```

将 `ResourceID` 转换为字符串。

**返回值**：
- 资源标识符的字符串表示

#### New

```go
// New 创建带前缀的唯一标识符.
func (rid ResourceID) New(counter uint64) string
```

使用计数器生成带前缀的唯一标识符。

**参数**：
- `counter` - 用于生成唯一标识符的计数器值

**返回值**：
- 格式为 "前缀-唯一字符串" 的完整资源标识符

#### Salt

```go
// Salt 计算机器 ID 的哈希值并返回一个 uint64 类型的盐值.
func Salt() uint64
```

计算基于机器 ID 的哈希值作为盐值，用于增强唯一标识符的唯一性。

**返回值**：
- 基于机器 ID 计算的 uint64 类型盐值

## 使用示例

### 1. 基本使用

```go
import (
    "fmt"
    "github.com/moweilong/mo/rid"
)

func main() {
    // 定义用户资源标识符
    userID := rid.NewResourceID("user")
    
    // 生成唯一标识符，传入一个递增的计数器值
    uniqueUserID := userID.New(1)
    fmt.Println(uniqueUserID) // 输出类似: user-abc123
    
    // 生成另一个唯一标识符
    anotherUserID := userID.New(2)
    fmt.Println(anotherUserID) // 输出类似: user-def456
}
```

### 2. 定义资源类型常量

```go
import (
    "fmt"
    "github.com/moweilong/mo/rid"
)

// 定义资源类型常量
const (
    UserID  rid.ResourceID = "user"
    PostID  rid.ResourceID = "post"
    OrderID rid.ResourceID = "order"
)

func main() {
    // 生成不同类型资源的唯一标识符
    userUID := UserID.New(1001)
    postUID := PostID.New(2001)
    orderUID := OrderID.New(3001)
    
    fmt.Println(userUID)  // 输出类似: user-xyz789
    fmt.Println(postUID)  // 输出类似: post-uvw456
    fmt.Println(orderUID) // 输出类似: order-rst123
}
```

### 3. 获取资源标识符前缀

```go
import (
    "fmt"
    "github.com/moweilong/mo/rid"
)

const UserID rid.ResourceID = "user"

func main() {
    // 获取资源标识符的前缀字符串
    prefix := UserID.String()
    fmt.Println(prefix) // 输出: user
}
```

## 特性和优势

1. **带前缀的标识符**：生成的唯一标识符包含资源类型前缀，便于识别资源类型
2. **固定长度**：生成的唯一标识符部分固定为 6 个字符，整体格式统一
3. **分布式友好**：使用机器 ID 作为盐值，增强了分布式环境下的唯一性
4. **简单易用**：API 设计简洁明了，易于集成到各种项目中
5. **基于计数器**：通过传入计数器值，保证了相同前缀下的标识符唯一性

## 实现原理

rid 包的核心实现基于以下原理：

1. **前缀标识**：使用 `ResourceID` 作为资源类型的前缀标识
2. **唯一字符串生成**：依赖 `idx.NewCode` 函数生成唯一字符串部分
   - 使用指定的字符集（默认为大小写字母和数字）
   - 固定长度为 6 个字符
   - 使用基于机器 ID 的盐值增强唯一性
3. **机器 ID 获取**：通过读取系统特定文件（如 `/etc/machine-id` 或 `/sys/class/dmi/id/product_uuid`）获取机器 ID，如果无法获取则使用主机名或生成随机 ID

## 注意事项

1. 生成唯一标识符时，应确保传入的计数器值在相同前缀下是唯一的
2. 在分布式环境中，建议为不同节点分配不同的计数器范围，以避免冲突
3. 机器 ID 的获取依赖于系统特定文件，在某些容器环境或特殊系统中可能会有所不同
4. 生成的唯一标识符包含固定长度的随机部分，在极高并发场景下仍有极小的碰撞概率