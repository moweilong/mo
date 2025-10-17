# mapper - 类型映射工具包

mapper 包提供了灵活的数据结构映射工具，用于简化不同类型或结构之间的数据转换，特别是在数据库实体与业务模型、API请求/响应模型之间的转换场景。它可以减少手动赋值的代码量，提高开发效率，并确保数据转换的一致性和可靠性。

## 目录结构

```
mapper/
├── README.md           # 文档
├── enum_converter.go   # 枚举类型转换器
├── interface.go        # 映射器接口定义
├── mapper.go           # 基于Copier的映射器实现
└── mapper_test.go      # 测试文件
```

## 核心接口

### Mapper 接口

```go
// Mapper 定义了在数据传输对象(DTO)和数据库实体之间进行转换的接口
type Mapper[DTO any, ENTITY any] interface {
	// ToEntity 将DTO转换为数据库实体
	ToEntity(*DTO) *ENTITY

	// ToDTO 将数据库实体转换为DTO
	ToDTO(*ENTITY) *DTO
}
```

## 实现类

### CopierMapper

基于 [jinzhu/copier](https://github.com/jinzhu/copier) 库实现的通用对象映射器，支持结构体间的字段复制。

```go
type CopierMapper[DTO any, ENTITY any] struct {
	copierOption copier.Option
}
```

主要方法：
- `NewCopierMapper[DTO, ENTITY]() *CopierMapper[DTO, ENTITY]` - 创建新的CopierMapper实例
- `AppendConverter(converter copier.TypeConverter)` - 添加单个类型转换器
- `AppendConverters(converters []copier.TypeConverter)` - 添加多个类型转换器
- `ToEntity(dto *DTO) *ENTITY` - 将DTO转换为实体
- `ToDTO(entity *ENTITY) *DTO` - 将实体转换为DTO

### EnumTypeConverter

专门用于枚举类型转换的映射器，特别适用于ent与protobuf之间的枚举类型映射。

```go
type EnumTypeConverter[DTO ~int32, ENTITY ~string] struct {
	nameMap  map[int32]string
	valueMap map[string]int32
}
```

主要方法：
- `NewEnumTypeConverter[DTO, ENTITY](nameMap, valueMap) *EnumTypeConverter[DTO, ENTITY]` - 创建枚举类型转换器
- `ToEntity(dto *DTO) *ENTITY` - 将DTO枚举值转换为实体枚举名称
- `ToDTO(entity *ENTITY) *DTO` - 将实体枚举名称转换为DTO枚举值
- `NewConverterPair() []copier.TypeConverter` - 创建可用于Copier的类型转换器对

## 使用示例

### 1. 基于Copier的类型映射器

```go
package main

import (
	"github.com/moweilong/mo/mapper"
)

func main() {
	// 定义DTO和实体类型
	type DtoType struct {
		Name string
		Age  int
	}

	type EntityType struct {
		Name string
		Age  int
	}

	// 创建映射器实例
	dataMapper := mapper.NewCopierMapper[DtoType, EntityType]()

	// 测试 ToEntity 方法：DTO 转 实体
	dto := &DtoType{Name: "Alice", Age: 25}
	entity := dataMapper.ToEntity(dto)

	// 测试 ToEntity 方法，处理 nil 输入
	entityNil := dataMapper.ToEntity(nil)

	// 测试 ToDTO 方法：实体 转 DTO
	entity = &EntityType{Name: "Bob", Age: 30}
	dtoResult := dataMapper.ToDTO(entity)

	// 测试 ToDTO 方法，处理 nil 输入
	dtoNil := dataMapper.ToDTO(nil)
}
```

### 2. 枚举类型映射器（ent 与 protobuf）

```go
package main

import "github.com/moweilong/mo/mapper"

func main() {
	// 定义枚举类型（protobuf风格和ent风格）
	type DtoType int32 // 类似protobuf的int32枚举
	type EntityType string // 类似ent的string枚举

	// DTO枚举值
	const (
		DtoTypeOne DtoType = 1
		DtoTypeTwo DtoType = 2
	)

	// 实体枚举值
	const (
		EntityTypeOne EntityType = "One"
		EntityTypeTwo EntityType = "Two"
	)

	// 创建映射表
	nameMap := map[int32]string{
		1: "One",
		2: "Two",
	}
	valueMap := map[string]int32{
		"One": 1,
		"Two": 2,
	}

	// 创建枚举类型转换器
	converter := mapper.NewEnumTypeConverter[DtoType, EntityType](nameMap, valueMap)

	// 测试 ToEntity 方法：int32枚举转string枚举
	dto := DtoTypeOne
	entity := converter.ToEntity(&dto)

	// 测试 ToEntity 方法，处理不存在的值
	dtoInvalid := DtoType(3)
	entityInvalid := converter.ToEntity(&dtoInvalid) // 返回nil

	// 测试 ToDTO 方法：string枚举转int32枚举
	tmpEntityTwo := EntityTypeTwo
	entity = &tmpEntityTwo
	dtoResult := converter.ToDTO(entity)

	// 测试 ToDTO 方法，处理不存在的值
	tmpEntityThree := EntityType("Three")
	entityInvalid = &tmpEntityThree
	dtoInvalidResult := converter.ToDTO(entityInvalid) // 返回nil
}
```

### 3. 自定义类型转换器

```go
package main

import (
	"github.com/moweilong/mo/mapper"
	"github.com/jinzhu/copier"
)

func main() {
	// 定义自定义类型
	type CustomType string
	type StandardType int

	// 创建自定义类型转换器对
	converter := mapper.NewGenericTypeConverterPair(
		CustomType(""),
		StandardType(0),
		func(src CustomType) StandardType {
			// 从CustomType转换到StandardType的逻辑
			if src == "one" {
				return 1
			}
			return 0
		},
		func(src StandardType) CustomType {
			// 从StandardType转换到CustomType的逻辑
			if src == 1 {
				return "one"
			}
			return "default"
		},
	)

	// 将自定义转换器添加到CopierMapper
	type Source struct {
		Field CustomType
	}

	type Target struct {
		Field StandardType
	}

	dataMapper := mapper.NewCopierMapper[Source, Target]()
	dataMapper.AppendConverters(converter)

	// 现在映射器可以自动处理自定义类型之间的转换
}
```

## 特性优势

- **泛型支持**：利用Go泛型提供类型安全的映射功能
- **接口抽象**：通过Mapper接口定义统一的数据转换规范
- **灵活配置**：支持自定义类型转换器，处理复杂的类型转换逻辑
- **空值安全**：对nil输入有良好处理，避免空指针异常
- **枚举类型支持**：专门的枚举类型映射器，简化枚举值与枚举名称的转换
- **无缝集成**：可与现有的结构体和类型系统无缝集成

## 注意事项

1. 在生产环境中，应考虑替换`mapper.go`中的`panic(err)`为更合适的错误处理机制
2. 使用`NewEnumTypeConverter`时，确保nameMap和valueMap的映射是完整且一致的
3. 对于复杂的类型转换，建议实现自定义的`copier.TypeConverter`以获得更好的控制
4. 当结构体字段名称或类型不匹配时，可能需要额外的配置或自定义转换器
5. 使用`AppendConverters`方法添加自定义类型转换器时，注意添加的顺序可能会影响转换结果
