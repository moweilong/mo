# trans 包

`trans` 是一个 Go 语言工具包，提供了丰富的类型转换函数，特别是用于指针与值、集合类型之间的转换，以及自定义默认值处理。该包需要 Go 1.18+ 版本，因为它广泛使用了泛型功能。

## 功能特性

- **基础类型转换**：提供指针与值之间的相互转换函数
- **集合类型转换**：支持切片和映射的指针与值之间的转换
- **UUID 处理**：提供 UUID 相关的解析和转换功能
- **自定义默认值**：支持自定义默认值提供者，处理 nil 指针的默认值
- **泛型支持**：利用 Go 1.18+ 的泛型特性，提供类型安全的转换函数

## 安装

```bash
go get github.com/moweilong/mo/trans
```

## 使用示例

### 基础类型转换

```go
import "github.com/moweilong/mo/trans"

// 将值转换为指针
strPtr := trans.ToPtr("hello")
intPtr := trans.ToPtr(42)

// 将指针转换为值，处理 nil 情况
str := trans.FromPtr(strPtr, trans.StringDefaultValue{})
intVal := trans.FromPtr(intPtr, trans.IntDefaultValue{})

// 使用专用的转换函数
strVal := trans.StringValue(strPtr) // 等价于 FromPtr(strPtr, StringDefaultValue{})
intVal2 := trans.IntValue(intPtr)   // 等价于 FromPtr(intPtr, IntDefaultValue{})
```

### 集合类型转换

```go
// 切片值转指针切片
intSlice := []int{1, 2, 3}
intPtrSlice := trans.SliceToPtrs(intSlice)

// 指针切片转值切片，处理 nil 指针
intValueSlice := trans.SliceFromPtrs(intPtrSlice, trans.IntDefaultValue{})

// 提取 map 的键和值
m := map[string]int{"a": 1, "b": 2}
keys := trans.MapKeys(m)    // []string{"a", "b"}（顺序不确定）
values := trans.MapValues(m) // []int{1, 2}（顺序不确定）

// 指针值 map 转值类型 map
ptrMap := map[string]*int{"a": &a, "b": &b}
valueMap := trans.MapFromPtrs(ptrMap, trans.IntDefaultValue{})
```

### UUID 处理

```go
import (
	"github.com/google/uuid"
	"github.com/moweilong/mo/trans"
)

// 字符串转 UUID
uid, err := trans.ToUuidE("123e4567-e89b-12d3-a456-426614174000")

// 字符串转 UUID 指针
uidPtr, err := trans.ToUuidPtrE("123e4567-e89b-12d3-a456-426614174000")

// UUID 转字符串指针
strPtr := trans.ToStringPtr(uuid.New())

// UUID 指针转 UUID 值
uidVal := trans.UUIDValue(uidPtr)

// UUID 切片转换
uuidSlice := []uuid.UUID{uuid.New(), uuid.New()}
uuidPtrSlice := trans.SliceToUUIDPtrs(uuidSlice)
valueSlice := trans.SliceFromUUIDPtrs(uuidPtrSlice)
```

### 自定义默认值提供者

```go
// 定义自定义类型
 type Person struct {
	Name string
	Age  int
}

// 实现 IDefaultValue 接口
 type PersonDefaultValue struct{}

func (p PersonDefaultValue) DefaultValue() Person {
	return Person{
		Name: "Unknown",
		Age:  0,
	}
}

// 使用自定义默认值提供者
var nilPersonPtr *Person
person := trans.FromPtr(nilPersonPtr, PersonDefaultValue{})
// person 将是 Person{Name: "Unknown", Age: 0}
```

## API 参考

### 基础类型转换函数

- `ToPtr[T any](v T) *T` - 将值转换为指针
- `FromPtr[T any](p *T, provider IDefaultValue[T]) T` - 将指针转换为值，使用默认值提供者处理 nil 情况
- `StringValue(a *string) string` - 将字符串指针转换为值
- `IntValue(a *int) int` - 将整数指针转换为值
- `BoolValue(a *bool) bool` - 将布尔指针转换为值
- `Float64Value(a *float64) float64` - 将浮点数指针转换为值
- 支持的其他基础类型：Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Time

### 批量指针转换

- `BatchPtrs[T any](vs []T) []*T` - 将值切片转换为指针切片（已弃用，推荐使用 SliceToPtrs）

### 集合类型转换函数

- `SliceToPtrs[T any](values []T) []*T` - 将值切片转换为指针切片
- `SliceFromPtrs[T any](ptrs []*T, provider IDefaultValue[T]) []T` - 将指针切片转换为值切片，处理 nil 指针
- `MapKeys[K comparable, V any](m map[K]V) []K` - 提取 map 的所有键
- `MapValues[K comparable, V any](m map[K]V) []V` - 提取 map 的所有值
- `MapFromPtrs[K comparable, V any](m map[K]*V, provider IDefaultValue[V]) map[K]V` - 将指针值 map 转换为值类型 map

### UUID 相关函数

- `ToUuidE(s string) (uuid.UUID, error)` - 将字符串转换为 UUID，返回错误信息
- `ToUuidPtrE(s *string) (*uuid.UUID, error)` - 将字符串指针转换为 UUID 指针，返回错误信息
- `ToUuid(s string) uuid.UUID` - 将字符串转换为 UUID，失败时返回 uuid.Nil
- `ToUuidPtr(s string) *uuid.UUID` - 将字符串转换为 UUID 指针，失败时返回 nil
- `ToStringPtr(u uuid.UUID) *string` - 将 UUID 转换为字符串指针
- `UUIDValue(a *uuid.UUID) uuid.UUID` - 将 UUID 指针转换为 UUID 值
- `SliceToUUIDPtrs(uuids []uuid.UUID) []*uuid.UUID` - 将 UUID 切片转换为 UUID 指针切片
- `SliceFromUUIDPtrs(ptrs []*uuid.UUID) []uuid.UUID` - 将 UUID 指针切片转换为 UUID 值切片

### 默认值提供者

`trans` 包为所有支持的基础类型提供了默认值提供者：

- `StringDefaultValue{}` - 提供空字符串默认值
- `IntDefaultValue{}`, `Int8DefaultValue{}`, 等 - 提供整数类型的 0 默认值
- `BoolDefaultValue{}` - 提供 false 默认值
- `Float32DefaultValue{}`, `Float64DefaultValue{}` - 提供浮点数类型的 0 默认值
- `TimeDefaultValue{}` - 提供当前时间作为默认值
- `UUIDDefaultValue{}` - 提供 uuid.Nil 作为默认值

## 实现原理

该包的核心是 `IDefaultValue[T any]` 接口，它定义了如何为特定类型提供默认值：

```go
// IDefaultValue 接口定义了获取类型默认值的方法
 type IDefaultValue[T any] interface {
	DefaultValue() T
}
```

所有转换函数都利用这个接口来处理 nil 指针情况，确保转换结果始终是有效的值。

## 测试

该包包含全面的单元测试，覆盖了所有主要功能和边缘情况：

```bash
go test -v ./trans/...
```

## 版本要求

- Go 1.18 或更高版本（因为使用了泛型特性）

## 依赖项

- `github.com/google/uuid` - 用于 UUID 相关功能
- `github.com/stretchr/testify` - 仅用于测试

## 注意事项

- 对于时间类型，`TimeValue` 函数在处理 nil 指针时会返回当前时间，而不是零值时间
- 对于集合类型的转换，处理 nil 切片会返回 nil，处理空切片会返回空切片
- 使用自定义默认值提供者时，请确保正确实现 `IDefaultValue[T]` 接口