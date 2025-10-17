# copierutil - 类型转换器工具包

copierutil 包提供了一组基于 [jinzhu/copier](https://github.com/jinzhu/copier) 的类型转换器工具，用于简化不同数据类型之间的自动转换，特别是在结构体映射过程中处理自定义类型转换的场景。

## 目录结构

```
copierutil/
├── README.md          # 文档
├── converters.go      # 类型转换器定义和实现
├── converters_test.go # 测试文件
├── go.mod             # Go模块定义
└── go.sum             # 依赖校验和
```

## 核心功能

copierutil 包主要提供以下核心功能：

1. **预定义类型转换器**：为常见类型（如时间、字符串）提供现成的转换器
2. **转换器创建函数**：提供灵活的API来创建自定义类型转换器
3. **泛型支持**：利用Go泛型提供类型安全的转换功能
4. **错误处理**：支持带错误返回的类型转换函数
5. **转换器对**：自动创建双向转换的转换器对

## 预定义转换器

### 时间相关转换器

| 转换器名称 | 源类型 | 目标类型 | 功能描述 |
|------------|--------|----------|----------|
| TimeToStringConverter | `*time.Time` | `*string` | 将时间转换为字符串 |
| StringToTimeConverter | `*string` | `*time.Time` | 将字符串转换为时间 |
| TimeToTimestamppbConverter | `*time.Time` | `*timestamppb.Timestamp` | 将时间转换为protobuf时间戳 |
| TimestamppbToTimeConverter | `*timestamppb.Timestamp` | `*time.Time` | 将protobuf时间戳转换为时间 |

## 核心API

### 单个转换器创建

```go
// NewTypeConverter 创建单个类型转换器
func NewTypeConverter(srcType, dstType interface{}, fn func(src interface{}) (interface{}, error)) copier.TypeConverter
```

### 转换器对创建

```go
// NewTypeConverterPair 创建一对类型转换器（双向）
func NewTypeConverterPair(srcType, dstType interface{}, fromFn, toFn func(src interface{}) (interface{}, error)) []copier.TypeConverter

// NewGenericTypeConverterPair 使用泛型创建类型安全的转换器对
func NewGenericTypeConverterPair[A interface{}, B interface{}](srcType A, dstType B, fromFn func(src A) B, toFn func(src B) A) []copier.TypeConverter

// NewErrorHandlingGenericTypeConverterPair 使用泛型创建支持错误处理的类型安全转换器对
func NewErrorHandlingGenericTypeConverterPair[A interface{}, B interface{}](srcType A, dstType B, fromFn func(src A) (B, error), toFn func(src B) (A, error)) []copier.TypeConverter
```

### 特定类型转换器对

```go
// NewTimeStringConverterPair 创建时间和字符串之间的转换器对
func NewTimeStringConverterPair() []copier.TypeConverter

// NewTimeTimestamppbConverterPair 创建时间和protobuf时间戳之间的转换器对
func NewTimeTimestamppbConverterPair() []copier.TypeConverter
```

### 辅助函数

```go
// TimeToString 将时间转换为ISO8601格式的字符串
func TimeToString(tm *time.Time) *string
```

## 使用示例

### 1. 使用预定义的时间转换器

```go
package main

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/moweilong/mo/copierutil"
)

func main() {
	// 创建带有时间转换器的copier选项
	options := copier.Option{
		Converters: []copier.TypeConverter{
			copierutil.TimeToStringConverter,
			copierutil.StringToTimeConverter,
		},
	}

	// 定义源结构体和目标结构体
	src := struct {
		CreatedAt time.Time
	}{
		CreatedAt: time.Now(),
	}

	var dst struct {
		CreatedAt string
	}

	// 使用带转换器的copier进行复制
	copier.CopyWithOption(&dst, src, options)
}
```

### 2. 创建自定义类型转换器

```go
package main

import (
	"github.com/jinzhu/copier"
	"github.com/moweilong/mo/copierutil"
)

func main() {
	// 定义自定义类型
	type CustomInt int
	type CustomString string

	// 创建自定义转换器对
	converters := copierutil.NewGenericTypeConverterPair(
		CustomInt(0),
		CustomString(""),
		func(src CustomInt) CustomString {
			return CustomString("prefix-" + string(rune(src+'0')))
		},
		func(src CustomString) CustomInt {
			if len(src) > 7 {
				return CustomInt(src[7] - '0')
			}
			return 0
		},
	)

	// 创建带转换器的copier选项
	options := copier.Option{
		Converters: converters,
	}

	// 使用自定义转换器进行复制
	// ...
}
```

### 3. 使用支持错误处理的转换器

```go
package main

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/moweilong/mo/copierutil"
)

func main() {
	// 定义自定义类型
	type ValidatedInt int
	type ValidatedString string

	// 创建支持错误处理的转换器对
	converters := copierutil.NewErrorHandlingGenericTypeConverterPair(
		ValidatedInt(0),
		ValidatedString(""),
		func(src ValidatedInt) (ValidatedString, error) {
			if src < 0 {
				return "", errors.New("negative value not allowed")
			}
			return ValidatedString(fmt.Sprintf("valid-%d", src)), nil
		},
		func(src ValidatedString) (ValidatedInt, error) {
			// 实现带错误检查的转换逻辑
			// ...
			return 0, nil
		},
	)

	// 创建带转换器的copier选项
	options := copier.Option{
		Converters: converters,
	}

	// 使用转换器进行复制，错误会向上传递
	// ...
}
```

### 4. 与其他库集成

```go
package main

import (
	"github.com/jinzhu/copier"
	"github.com/moweilong/mo/copierutil"
	"github.com/moweilong/mo/mapper"
)

func main() {
	// 定义DTO和实体类型
	type UserDTO struct {
		CreatedAt string
	}

	type UserEntity struct {
		CreatedAt time.Time
	}

	// 创建带有时间转换器的映射器
	userMapper := mapper.NewCopierMapper[UserDTO, UserEntity]()
	timeConverters := copierutil.NewTimeStringConverterPair()
	userMapper.AppendConverters(timeConverters)

	// 现在映射器可以自动处理时间和字符串之间的转换
	// ...
}
```

## 特性优势

- **类型安全**：利用Go泛型提供编译时类型检查
- **灵活扩展**：支持自定义任意类型的转换逻辑
- **错误处理**：提供带错误返回的转换器变体
- **预定义转换器**：内置常用的时间和Protobuf类型转换器
- **与copier集成**：无缝配合jinzhu/copier进行结构体映射
- **双向转换**：自动创建支持双向转换的转换器对

## 依赖关系

- [github.com/jinzhu/copier](https://github.com/jinzhu/copier) - 核心复制功能
- [google.golang.org/protobuf](https://github.com/golang/protobuf) - Protobuf支持
- [github.com/moweilong/mo/timeutil](https://github.com/moweilong/mo/tree/main/timeutil) - 时间处理工具
- [github.com/moweilong/mo/trans](https://github.com/moweilong/mo/tree/main/trans) - 类型转换辅助工具

## 注意事项

1. 使用泛型转换器时，请确保传入的类型参数与实际类型匹配
2. 处理可能失败的转换时，应使用`NewErrorHandlingGenericTypeConverterPair`
3. 转换器对中的两个方向转换应该保持逻辑一致性，避免数据丢失
4. 对于复杂的嵌套结构体，可能需要为每个自定义类型定义单独的转换器
5. 在处理指针类型时，确保适当地处理nil值情况