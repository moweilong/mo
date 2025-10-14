# entx/update

`entx/update` 是 entx 包下的一个子模块，提供了一组用于构建和执行 SQL 更新操作的工具函数，特别适用于处理 NULL 值和 JSON 字段的更新场景。

## 功能特性

- **NULL 值处理**：提供便捷的方法将数据库字段设置为 NULL
- **JSON 字段更新**：支持对 JSON 类型字段进行部分更新，包括设置具体值和删除特定路径
- **字段命名转换**：支持驼峰命名到蛇形命名的自动转换，适配不同数据库命名规范
- **与 Protocol Buffers 集成**：可以直接从 Proto 消息中提取字段值进行更新

## 函数列表

### NULL 值处理函数

#### `BuildSetNullUpdate(u *sql.UpdateBuilder, fields []string)`
构建一个 SQL UPDATE 语句，将指定的字段设置为 NULL。

- **参数**：
  - `u`：SQL 更新构建器实例
  - `fields`：要设置为 NULL 的字段名称列表

- **示例**：
  ```go
  import (
      "entgo.io/ent/dialect/sql"
      "github.com/moweilong/mo/entx/update"
  )
  
  builder := sql.Dialect(dialect.MySQL).Update("users")
  update.BuildSetNullUpdate(builder, []string{"email", "phone"})
  // 生成的 SQL: UPDATE `users` SET `email` = NULL, `phone` = NULL
  ```

#### `BuildSetNullUpdater(fields []string) func(u *sql.UpdateBuilder)`
返回一个函数，该函数用于将指定字段设置为 NULL。适用于需要延迟构建更新操作的场景。

- **参数**：
  - `fields`：要设置为 NULL 的字段名称列表
- **返回值**：
  - 一个接受 `sql.UpdateBuilder` 参数的函数

- **示例**：
  ```go
  updater := update.BuildSetNullUpdater([]string{"profile_picture", "bio"})
  if userDeleted {
      updater(userBuilder)
  }
  ```

### JSON 字段处理函数

#### `ExtractJsonFieldKeyValues(msg proto.Message, paths []string, needToSnakeCase bool) []string`
从 Proto 消息中提取指定路径字段的键值对，用于构建 JSON 字段更新。

- **参数**：
  - `msg`：Proto 消息实例
  - `paths`：要提取的字段路径列表
  - `needToSnakeCase`：是否需要将字段名转换为蛇形命名
- **返回值**：
  - 包含键值对的字符串切片

#### `SetJsonNullFieldUpdateBuilder(fieldName string, msg proto.Message, paths []string) func(u *sql.UpdateBuilder)`
创建一个函数，用于从 JSON 字段中删除指定路径的值（设置为 NULL）。

- **参数**：
  - `fieldName`：JSON 字段的名称
  - `msg`：Proto 消息实例
  - `paths`：要删除的 JSON 路径列表
- **返回值**：
  - 一个接受 `sql.UpdateBuilder` 参数的函数

- **示例**：
  ```go
  // 从 user_profile JSON 字段中删除 address 和 phone 路径
  jsonUpdater := update.SetJsonNullFieldUpdateBuilder("user_profile", userProto, []string{"address", "phone"})
  if jsonUpdater != nil {
      jsonUpdater(builder)
  }
  ```

#### `SetJsonFieldValueUpdateBuilder(fieldName string, msg proto.Message, paths []string, needToSnakeCase bool) func(u *sql.UpdateBuilder)`
创建一个函数，用于更新 JSON 字段中的特定路径值。

- **参数**：
  - `fieldName`：JSON 字段的名称
  - `msg`：Proto 消息实例
  - `paths`：要更新的 JSON 路径列表
  - `needToSnakeCase`：是否需要将字段名转换为蛇形命名
- **返回值**：
  - 一个接受 `sql.UpdateBuilder` 参数的函数

- **示例**：
  ```go
  // 更新 user_preferences JSON 字段中的 theme 和 language 路径
  jsonUpdater := update.SetJsonFieldValueUpdateBuilder("user_preferences", userProto, []string{"theme", "language"}, true)
  if jsonUpdater != nil {
      jsonUpdater(builder)
  }
  ```

## 使用场景

### 1. 部分字段清空

当需要将记录中的部分字段设置为 NULL 时，可以使用 `BuildSetNullUpdate` 或 `BuildSetNullUpdater` 函数：

```go
// 直接构建 NULL 更新
update.BuildSetNullUpdate(builder, []string{"last_login_time", "session_token"})

// 或使用延迟执行的方式
updater := update.BuildSetNullUpdater([]string{"last_login_time", "session_token"})
// ... 其他逻辑 ...
updater(builder)
```

### 2. JSON 字段部分更新

对于存储在 JSON 类型字段中的复杂数据，可以使用 JSON 字段处理函数进行部分更新：

```go
// 设置 JSON 字段的特定路径值
builder := sql.Dialect(dialect.Postgres).Update("users")
jsonUpdater := update.SetJsonFieldValueUpdateBuilder("preferences", userProto, []string{"notification.email", "theme.dark_mode"}, true)
if jsonUpdater != nil {
    jsonUpdater(builder)
}

// 从 JSON 字段中删除特定路径
nullUpdater := update.SetJsonNullFieldUpdateBuilder("profile", userProto, []string{"temporary_data"})
if nullUpdater != nil {
    nullUpdater(builder)
}
```

### 3. 结合 FieldMask 使用

可以结合 `fieldmaskutil` 包一起使用，根据字段掩码动态确定要更新的字段：

```go
// 假设 fieldMask 是从请求中提取的字段掩码
paths := fieldmaskutil.GetPathsFromFieldMask(fieldMask)

// 只更新指定的字段
builder := sql.Dialect(dialect.MySQL).Update("users").Where(sql.EQ("id", userID))

// 处理 JSON 字段更新
jsonUpdater := update.SetJsonFieldValueUpdateBuilder("metadata", userProto, paths, true)
if jsonUpdater != nil {
    jsonUpdater(builder)
}
```

## 注意事项

1. 函数会自动处理不同数据库方言（如 MySQL、PostgreSQL）的 SQL 语法差异
2. 当 `needToSnakeCase` 设为 `true` 时，会自动将驼峰命名的字段转换为蛇形命名
3. 对于不存在的字段路径，函数会安全地跳过而不会引发错误
4. 当没有要更新的字段时，部分函数会返回 `nil`，使用前请检查返回值
5. JSON 字段操作主要针对 PostgreSQL 中的 JSONB 类型优化，其他数据库可能需要调整

## 依赖关系

- `entgo.io/ent/dialect/sql`：提供 SQL 构建器功能
- `github.com/moweilong/mo/fieldmaskutil`：用于处理字段掩码
- `github.com/moweilong/mo/stringcase`：用于字段命名转换
- `google.golang.org/protobuf`：用于处理 Protocol Buffers 消息