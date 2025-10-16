# store - 通用数据存储层

store 是一个基于 GORM 的通用数据访问层封装，提供了统一的 CRUD 操作接口和查询条件构建功能，支持泛型和可配置的日志记录。

## 功能概述

- **泛型支持**：使用 Go 泛型提供类型安全的数据访问操作
- **统一接口**：封装了标准的 CRUD 操作接口
- **灵活查询**：提供强大的查询条件构建功能，支持分页、过滤等
- **可配置日志**：支持自定义日志记录器，默认提供空日志和基于 mo/log 的日志器
- **模型管理**：提供模型注册和数据库迁移功能

## 目录结构

```
store/
├── logger/
│   ├── empty/      # 空日志实现，不执行任何日志操作
│   │   └── logger.go
│   └── mo/         # 基于 mo/log 的日志实现
│       └── logger.go
├── registry/       # 模型注册和迁移功能
│   └── registry.go
├── where/          # 查询条件构建功能
│   └── where.go
├── logger.go       # 日志接口定义
└── store.go        # 核心存储接口和实现
```

## 核心接口和结构体

### DBProvider
```go
// DBProvider 定义了提供数据库连接的接口
type DBProvider interface {
    // DB 返回给定上下文中的数据库实例
    DB(ctx context.Context, wheres ...where.Where) *gorm.DB
}
```

### Store
```go
// Store 表示具有日志功能的通用数据存储
type Store[T any] struct {
    logger  Logger
    storage DBProvider
}
```

### Logger
```go
// Logger 定义了记录错误和上下文信息的接口
type Logger interface {
    // Error 记录错误消息及相关上下文
    Error(ctx context.Context, err error, message string, kvs ...any)
}
```

### Where
```go
// Where 定义了可修改 GORM 数据库查询的接口
type Where interface {
    Where(db *gorm.DB) *gorm.DB
}
```

### Options
```go
// Options 包含 GORM Where 查询条件的选项
type Options struct {
    Offset  int               // 分页起始点
    Limit   int               // 每页最大结果数
    Filters map[any]any       // 过滤条件键值对
    Clauses []clause.Expression // 自定义子句
    Queries []Query           // 查询条件列表
}
```

## 主要功能

### 基本 CRUD 操作

#### 创建 Store 实例
```go
import (
    "github.com/moweilong/mo/store"
    "github.com/moweilong/mo/store/logger/mo"
)

// 创建一个使用 mo/log 的日志器
logger := mo.NewLogger()

// 创建 Store 实例（假设 dbProvider 是一个实现了 DBProvider 接口的对象）
userStore := store.NewStore[User](dbProvider, logger)
```

#### 插入数据
```go
user := &User{Name: "张三", Age: 30}
err := userStore.Create(ctx, user)
if err != nil {
    // 处理错误
}
```

#### 更新数据
```go
user.Name = "李四"
err := userStore.Update(ctx, user)
if err != nil {
    // 处理错误
}
```

#### 删除数据
```go
// 创建删除条件（ID等于10的记录）
opts := where.NewWhere(
    where.WithQuery("id = ?", 10),
)
err := userStore.Delete(ctx, opts)
if err != nil {
    // 处理错误
}
```

#### 获取单条数据
```go
// 创建查询条件（ID等于10的记录）
opts := where.NewWhere(
    where.WithQuery("id = ?", 10),
)
user, err := userStore.Get(ctx, opts)
if err != nil {
    // 处理错误
}
```

#### 获取列表数据
```go
// 创建带分页的查询条件
opts := where.NewWhere(
    where.WithQuery("age > ?", 18),
    where.WithPage(1, 10), // 第1页，每页10条
)
count, users, err := userStore.List(ctx, opts)
if err != nil {
    // 处理错误
}
// count 是符合条件的总记录数
// users 是当前页的数据列表
```

### 查询条件构建

#### 基本条件查询
```go
// 使用函数选项模式创建条件
opts := where.NewWhere(
    where.WithQuery("status = ?", "active"),
    where.WithFilter(map[any]any{"category": "premium"}),
)

// 链式调用方式创建条件
opts := where.NewWhere().
    Q("status = ?", "active").
    F("category", "premium")
```

#### 分页查询
```go
// 使用函数选项模式
opts := where.NewWhere(
    where.WithPage(2, 20), // 第2页，每页20条
)

// 链式调用方式
opts := where.NewWhere().P(2, 20)

// 直接设置偏移量和限制
opts := where.NewWhere(
    where.WithOffset(20),
    where.WithLimit(20),
)

// 链式调用方式
opts := where.NewWhere().O(20).L(20)
```

#### 自定义 SQL 子句
```go
import (
    "gorm.io/gorm/clause"
)

// 添加 ORDER BY 子句
opts := where.NewWhere(
    where.WithClauses(clause.OrderBy{Expression: clause.Expr{SQL: "created_at DESC"}}),
)

// 链式调用方式
opts := where.NewWhere().C(
    clause.OrderBy{Expression: clause.Expr{SQL: "created_at DESC"}},
)
```

### 模型注册与迁移

```go
import (
    "github.com/moweilong/mo/store/registry"
    "gorm.io/gorm"
)

// 注册模型
registry.Register(&User{})
registry.Register(&Product{})

// 执行数据库迁移
var db *gorm.DB // 假设这是一个有效的 GORM DB 实例
err := registry.Migrate(db)
if err != nil {
    // 处理迁移错误
}
```

### 日志配置

store 包支持多种日志记录方式：

#### 使用空日志器（不记录任何日志）
```go
import (
    "github.com/moweilong/mo/store/logger/empty"
)

logger := empty.NewLogger()
userStore := store.NewStore[User](dbProvider, logger)
```

#### 使用 mo/log 日志器
```go
import (
    "github.com/moweilong/mo/store/logger/mo"
)

logger := mo.NewLogger()
userStore := store.NewStore[User](dbProvider, logger)
```

#### 自定义日志器
```go
import (
    "context"
    "log"
)

// 实现自定义日志器
type customLogger struct{}

func (l *customLogger) Error(ctx context.Context, err error, msg string, kvs ...any) {
    // 实现自定义日志逻辑
    log.Printf("Error: %s, Message: %s", err, msg)
}

// 使用自定义日志器
logger := &customLogger{}
userStore := store.NewStore[User](dbProvider, logger)
```

## 使用示例

### 完整示例：用户管理

```go
import (
    "context"
    
    "github.com/moweilong/mo/store"
    "github.com/moweilong/mo/store/logger/mo"
    "github.com/moweilong/mo/store/where"
)

// 定义用户模型
type User struct {
    ID        uint   `gorm:"primarykey"`
    Name      string `gorm:"size:100"`
    Email     string `gorm:"size:100;uniqueIndex"`
    Age       int
    Status    string `gorm:"size:20"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 创建 DBProvider 实现
type MyDBProvider struct {
    db *gorm.DB
}

func (p *MyDBProvider) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
    dbInstance := p.db.WithContext(ctx)
    for _, whr := range wheres {
        if whr != nil {
            dbInstance = whr.Where(dbInstance)
        }
    }
    return dbInstance
}

// 使用示例
func main() {
    // 假设 db 是一个已初始化的 GORM DB 实例
    dbProvider := &MyDBProvider{db: db}
    logger := mo.NewLogger()
    userStore := store.NewStore[User](dbProvider, logger)
    ctx := context.Background()
    
    // 创建用户
    user := &User{Name: "张三", Email: "zhangsan@example.com", Age: 30, Status: "active"}
    err := userStore.Create(ctx, user)
    
    // 查询用户
    opts := where.NewWhere().
        Q("status = ?", "active").
        P(1, 10)
    count, users, err := userStore.List(ctx, opts)
}
```

## 特性与优势

1. **类型安全**：利用 Go 泛型提供类型安全的数据访问操作
2. **统一接口**：标准化的 CRUD 接口简化了数据访问层的实现
3. **灵活配置**：支持自定义日志器和数据库提供者
4. **强大查询**：丰富的查询条件构建功能，支持复杂查询场景
5. **可扩展性**：模块化设计，易于扩展和定制

## 注意事项

1. 使用 `where.NewWhere()` 创建查询条件时，注意参数类型要与数据库字段类型匹配
2. 日志记录可能会包含敏感数据，请根据实际需求配置适当的日志级别和字段过滤
3. 在高并发场景下，建议合理设置分页参数，避免一次性查询大量数据
4. 使用自定义 `DBProvider` 时，确保正确处理数据库连接和事务上下文

## 依赖

- gorm.io/gorm：Go 的 ORM 库
- github.com/moweilong/mo/log：可选的日志库，用于 mo 日志器实现