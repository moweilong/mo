# EntX Mixin 组件库

EntX Mixin 是一组基于 [Ent](https://entgo.io/) 框架的可复用 Schema 组件，旨在简化数据库模型定义，提供常见功能的标准化实现。

## 目录结构

```
├── autoincrement_id.go  # 自增ID组件
├── creator_id.go        # 创建者ID组件
├── operator.go          # 操作人相关组件
├── remark.go            # 备注字段组件
├── snowflake_id.go      # 雪花算法ID组件
├── string_id.go         # 字符串ID组件
├── switch_status.go     # 开关状态组件
├── time.go              # 时间字段组件（time.Time类型）
├── timestamp.go         # 时间戳组件（int64类型）
└── uuid_id.go           # UUID类型ID组件
```

## 组件分类与说明

### 1. ID 生成组件

ID 生成组件提供了多种 ID 类型的实现，可根据项目需求选择合适的 ID 策略。

#### 1.1 AutoIncrementId

自增 ID 组件，适用于需要简单递增序列的场景。

**特性：**
- 基于 uint32 类型
- 支持 MySQL 和 PostgreSQL 数据库类型映射
- 配置为正数、不可变、唯一
- 支持 entproto 注解

**使用示例：**
```go
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.AutoIncrementId{},
    }
}
```

#### 1.2 SnowflackId

雪花算法 ID 组件，适用于分布式系统中需要全局唯一 ID 的场景。

**特性：**
- 基于 uint64 类型
- 使用 Sonyflake 算法生成 ID
- 配置为正数、不可变
- 支持 MySQL 和 PostgreSQL 数据库类型映射

**使用示例：**
```go
func (Order) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.SnowflackId{},
    }
}
```

#### 1.3 StringId

字符串 ID 组件，适用于需要自定义字符串格式 ID 的场景。

**特性：**
- 基于 string 类型
- 最大长度 25 字符
- 支持字母、数字、下划线和连字符
- 配置为非空、唯一、不可变

**使用示例：**
```go
func (Product) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.StringId{},
    }
}
```

#### 1.4 UuidId

UUID ID 组件，适用于需要全局唯一标识符的场景。

**特性：**
- 基于 uuid.UUID 类型
- 自动生成新的 UUID 值
- 配置为唯一、不可变

**使用示例：**
```go
func (Session) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.UuidId{},
    }
}
```

### 2. 时间相关组件

时间相关组件提供了两种时间表示方式：基于 `time.Time` 的时间字段和基于 `int64` 的时间戳字段。

#### 2.1 基于 time.Time 的时间组件

**组件列表：**
- `CreatedAt`: 创建时间字段（created_at）
- `UpdatedAt`: 更新时间字段（updated_at）
- `DeletedAt`: 删除时间字段（deleted_at）
- `TimeAt`: 组合了以上三个字段
- `CreateTime`: 创建时间字段（create_time）- 替代命名风格
- `UpdateTime`: 更新时间字段（update_time）- 替代命名风格
- `DeleteTime`: 删除时间字段（delete_time）- 替代命名风格
- `Time`: 组合了以上三个替代命名风格字段

**特性：**
- 基于 time.Time 类型
- CreatedAt/CreateTime 配置为不可变
- 所有字段均配置为可选、可为空

**使用示例：**
```go
// 使用驼峰命名风格
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.TimeAt{}, // 包含 created_at, updated_at, deleted_at
    }
}

// 或使用下划线命名风格
func (Product) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{}, // 包含 create_time, update_time, delete_time
    }
}
```

#### 2.2 基于 int64 的时间戳组件

**组件列表：**
- `CreateTimestamp`: 创建时间戳（create_time）
- `UpdateTimestamp`: 更新时间戳（update_time）
- `DeleteTimestamp`: 删除时间戳（delete_time）
- `Timestamp`: 组合了以上三个字段
- `CreatedAtTimestamp`: 创建时间戳（created_at）- 替代命名风格
- `UpdatedAtTimestamp`: 更新时间戳（updated_at）- 替代命名风格
- `DeletedAtTimestamp`: 删除时间戳（deleted_at）- 替代命名风格
- `TimestampAt`: 组合了以上三个替代命名风格字段

**特性：**
- 基于 int64 类型，存储毫秒级时间戳
- 自动设置默认值和更新默认值
- CreateTimestamp/CreatedAtTimestamp 配置为不可变
- 所有字段均配置为可选、可为空

**使用示例：**
```go
// 使用驼峰命名风格
func (Log) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.TimestampAt{}, // 包含 created_at, updated_at, deleted_at
    }
}

// 或使用下划线命名风格
func (Event) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Timestamp{}, // 包含 create_time, update_time, delete_time
    }
}
```

### 3. 操作人相关组件

操作人相关组件用于记录数据操作的用户信息。

#### 3.1 Operator 组件

**组件列表：**
- `CreateBy`: 创建者ID字段（create_by）
- `UpdateBy`: 更新者ID字段（update_by）
- `DeleteBy`: 删除者ID字段（delete_by）
- `CreatedBy`: 创建者ID字段（created_by）- 替代命名风格
- `UpdatedBy`: 更新者ID字段（updated_by）- 替代命名风格
- `DeletedBy`: 删除者ID字段（deleted_by）- 替代命名风格

**特性：**
- 基于 uint32 类型
- 所有字段均配置为可选、可为空

**使用示例：**
```go
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.CreatedBy{},
        mixin.UpdatedBy{},
    }
}
```

#### 3.2 CreatorId 组件

创建者ID组件，专为仅需要记录创建者信息的场景设计。

**特性：**
- 基于 uint64 类型
- 配置为不可变、可选、可为空

**使用示例：**
```go
func (Article) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.CreatorId{},
    }
}
```

### 4. 其他功能组件

#### 4.1 Remark 组件

备注字段组件，用于添加额外说明信息。

**特性：**
- 基于 string 类型
- 默认值为空字符串
- 配置为可选、可为空

**使用示例：**
```go
func (Setting) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Remark{},
    }
}
```

#### 4.2 SwitchStatus 组件

开关状态组件，用于表示二元状态。

**特性：**
- 基于枚举类型
- 支持 "OFF" 和 "ON" 两个值
- 默认值为 "ON"
- 配置为可选、可为空

**注意：** 在 PostgreSQL 数据库中使用时，需要手动创建对应的枚举类型。

**使用示例：**
```go
func (Feature) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.SwitchStatus{},
    }
}
```

## 最佳实践

1. **命名风格一致性**
   - 在同一个项目中，建议统一使用一种命名风格（驼峰式或下划线式）的mixin组件
   - 例如：统一使用 `CreatedAt`、`UpdatedAt` 或统一使用 `CreateTime`、`UpdateTime`

2. **ID 策略选择**
   - 小型项目或单机部署可选择 `AutoIncrementId`
   - 分布式系统建议使用 `SnowflackId` 或 `UuidId`
   - 需要自定义ID格式时选择 `StringId`

3. **时间表示选择**
   - 需要精确时间操作和数据库原生时间类型时，选择基于 `time.Time` 的组件
   - 对性能要求高、需要跨系统交互或存储优化时，选择基于 `int64` 的时间戳组件

4. **组件组合使用**
   - 根据业务需求组合使用不同类型的组件
   - 例如：一个完整的用户模型可能同时包含 ID、时间戳和操作人组件

```go
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.AutoIncrementId{},
        mixin.TimeAt{},
        mixin.CreatedBy{},
        mixin.UpdatedBy{},
        mixin.Remark{},
    }
}
```