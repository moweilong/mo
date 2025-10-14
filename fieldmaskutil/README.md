# fieldmaskutil

`fieldmaskutil` 是一个用于处理 Protocol Buffers 字段掩码（FieldMask）的工具包，提供了简单而强大的方法来过滤、裁剪和覆盖 Proto 消息中的字段。

## 功能特性

- **过滤字段**：保留消息中指定的字段，清除其余所有字段
- **裁剪字段**：清除消息中指定的字段，保留其余所有字段
- **覆盖字段**：使用源消息中的值覆盖目标消息中指定的字段
- **支持嵌套字段**：可以处理多层嵌套的消息结构
- **支持集合类型**：支持处理映射（Map）和列表（List）类型的字段

## 安装

```bash
go get github.com/moweilong/mo/fieldmaskutil
```

## 快速开始

### 基本用法

```go
import (
    "github.com/moweilong/mo/fieldmaskutil"
    "google.golang.org/protobuf/proto"
    "yourproject/pbgen/yourproto"
)

// 创建 Proto 消息
msg := &yourproto.User{
    Id:    123,
    Name:  "John Doe",
    Email: "john@example.com",
    Profile: &yourproto.Profile{
        Age:     30,
        Address: "123 Main St",
    },
}

// 只保留指定的字段
fieldmaskutil.Filter(msg, []string{"name", "profile.age"})
// 现在 msg 中只剩下 name 和 profile.age 字段

// 清除指定的字段
msg2 := &yourproto.User{
    Id:    123,
    Name:  "John Doe",
    Email: "john@example.com",
}
fieldmaskutil.Prune(msg2, []string{"email"})
// 现在 msg2 中的 email 字段被清除

// 从源消息覆盖目标消息的指定字段
src := &yourproto.User{
    Name: "Jane Doe",
    Profile: &yourproto.Profile{
        Age: 25,
    },
}
dest := &yourproto.User{
    Id:    456,
    Name:  "Old Name",
    Email: "old@example.com",
    Profile: &yourproto.Profile{
        Age:     99,
        Address: "Old Address",
    },
}
fieldmaskutil.Overwrite(src, dest, []string{"name", "profile.age"})
// 现在 dest 中的 name 和 profile.age 字段被 src 中的值覆盖
```

### 高级用法（复用 NestedMask）

当需要多次使用相同的字段掩码处理多个消息时，可以直接使用 `NestedMask` 来提高性能：

```go
// 创建字段掩码
mask := fieldmaskutil.NestedMaskFromPaths([]string{"name", "profile.age"})

// 多次使用同一掩码
mask.Filter(msg1)
mask.Filter(msg2)
mask.Filter(msg3)

// 或者进行其他操作
mask.Prune(msg4)
mask.Overwrite(srcMsg, destMsg)
```

## API 参考

### 函数

#### `Filter(msg proto.Message, paths []string)`
保留消息中在 paths 列表中列出的字段，清除其余所有字段。

- **参数**：
  - `msg`：要处理的 Proto 消息
  - `paths`：要保留的字段路径列表

#### `Prune(msg proto.Message, paths []string)`
清除消息中在 paths 列表中列出的字段，保留其余所有字段。

- **参数**：
  - `msg`：要处理的 Proto 消息
  - `paths`：要清除的字段路径列表

#### `Overwrite(src, dest proto.Message, paths []string)`
使用源消息中的值覆盖目标消息中在 paths 列表中列出的字段。

- **参数**：
  - `src`：源 Proto 消息
  - `dest`：目标 Proto 消息
  - `paths`：要覆盖的字段路径列表

#### `NestedMaskFromPaths(paths []string) NestedMask`
从路径列表创建 NestedMask 实例。

- **参数**：
  - `paths`：字段路径列表
- **返回值**：
  - 创建的 NestedMask 实例

### 类型

#### `NestedMask map[string]NestedMask`
代表字段掩码的递归 map 结构。

- **方法**：
  - `Filter(msg proto.Message)`：保留消息中在掩码中列出的字段
  - `Prune(msg proto.Message)`：清除消息中在掩码中列出的字段
  - `Overwrite(src, dest proto.Message)`：使用源消息中的值覆盖目标消息中在掩码中列出的字段

## 路径格式

字段路径使用点号（.）分隔的格式表示嵌套字段，例如：

- `name`：表示顶级字段 name
- `profile.age`：表示嵌套在 profile 字段下的 age 字段
- `address.street.name`：表示多层嵌套的字段

## 注意事项

1. 路径被假定为有效且已规范化，否则函数可能会 panic
2. 空掩码（空的 NestedMask 或空的路径列表）不会对消息进行任何修改
3. 对于嵌套消息字段，如果在目标消息中该字段为 nil，Overwrite 操作会先初始化该字段
4. 如果源消息中的字段是空值，Overwrite 操作会清除目标消息中的对应字段

## 示例

```go
// 过滤嵌套字段
user := &User{
    Name: "John",
    Contacts: []*Contact{
        {Type: "email", Value: "john@example.com"},
        {Type: "phone", Value: "123-456-7890"},
    },
}
fieldmaskutil.Filter(user, []string{"contacts.value"})
// 现在 user 中的 contacts 数组只保留了 value 字段

// 处理映射字段
config := &Config{
    Settings: map[string]*Setting{
        "api": {Enabled: true, Timeout: 30},
        "db":  {Enabled: true, Timeout: 60},
    },
}
fieldmaskutil.Filter(config, []string{"settings.api.enabled"})
// 现在 config 中的 settings 映射只保留了 api.enabled 字段
```