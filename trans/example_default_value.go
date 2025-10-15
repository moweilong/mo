//go:build ignore
// +build ignore

package trans

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// 这个文件包含了 IDefaultValue 接口的使用示例

func ExampleDefaultValue() {
	// 1. 基本类型的默认值处理
	var strPtr *string
	var intPtr *int
	var boolPtr *bool
	var int64Ptr *int64
	var uintPtr *uint
	var float32Ptr *float32
	var float64Ptr *float64
	var timePtr *time.Time

	// 使用 IDefaultValue 接口获取默认值
	fmt.Println("String value:", StringValue(strPtr)) // 输出: ""
	fmt.Println("Int value:", IntValue(intPtr))       // 输出: 0
	fmt.Println("Bool value:", BoolValue(boolPtr))    // 输出: false
	fmt.Println("Int64 value:", Int64Value(int64Ptr))  // 输出: 0
	fmt.Println("Uint value:", UintValue(uintPtr))    // 输出: 0
	fmt.Println("Float32 value:", Float32Value(float32Ptr)) // 输出: 0.0
	fmt.Println("Float64 value:", Float64Value(float64Ptr)) // 输出: 0.0
	fmt.Println("Time value:", TimeValue(timePtr))     // 输出: 当前时间

	// 直接使用 FromPtr 函数和默认值提供者
	fmt.Println("String from ptr:", FromPtr(strPtr, StringDefaultValue{})) // 输出: ""
	fmt.Println("Int from ptr:", FromPtr(intPtr, IntDefaultValue{}))       // 输出: 0

	// 2. 切片的默认值处理
	strPtrs := []*string{ToStringPtr("hello"), nil, ToStringPtr("world")}
	intPtrs := []*int{ToPtr(10), nil, ToPtr(20)}

	// 将指针切片转换为值切片，处理 nil 指针
	strValues := SliceFromPtrs(strPtrs, StringDefaultValue{})
	intValues := SliceFromPtrs(intPtrs, IntDefaultValue{})
	fmt.Println("String slice values:", strValues) // 输出: [hello  world]
	fmt.Println("Int slice values:", intValues)     // 输出: [10 0 20]

	// 将值切片转换为指针切片
	numValues := []int{1, 2, 3}
	numPtrs := SliceToPtrs(numValues)
	fmt.Println("Number pointers:", numPtrs) // 输出: [0x... 0x... 0x...]

	// 使用 BatchPtrs 批量创建指针
	batchPtrs := BatchPtrs("a", "b", "c")
	fmt.Println("Batch pointers:", batchPtrs) // 输出: [0x... 0x... 0x...]

	// 3. UUID 类型的默认值处理
	var uuidPtr *uuid.UUID

	// 获取 UUID 默认值
	fmt.Println("UUID value:", UUIDValue(uuidPtr)) // 输出: 00000000-0000-0000-0000-000000000000

	// UUID 转换函数示例
	uuidStr := "123e4567-e89b-12d3-a456-426614174000"
	uuidVal := ToUuid(uuidStr)
	uuidPtrVal := ToUuidPtr(uuidStr)
	strPtrVal := ToStringPtr(uuidVal)
	fmt.Println("UUID from string:", uuidVal)
	fmt.Println("UUID pointer from string:", uuidPtrVal)
	fmt.Println("String pointer from UUID:", *strPtrVal)

	// 4. UUID 切片的默认值处理
	validUUID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	uuidPtrs := []*uuid.UUID{&validUUID, nil}

	// 处理 UUID 切片中的 nil 指针
	uuidValues := SliceFromUUIDPtrs(uuidPtrs)
	fmt.Println("UUID slice values:", uuidValues) // 输出: [123e4567-e89b-12d3-a456-426614174000 00000000-0000-0000-0000-000000000000]

	// UUID 切片和指针切片的转换
	uuidSlice := []uuid.UUID{validUUID, uuid.Nil}
	uuidPtrSlice := SliceToUUIDPtrs(uuidSlice)
	fmt.Println("UUID pointer slice:", uuidPtrSlice) // 输出: [0x... 0x...]

	// 5. Map 的默认值处理
	intMap := map[string]*int{
		"key1": ToPtr(42),
		"key2": nil,
	}

	// 将指针值的 map 转换为值类型的 map，处理 nil 指针
	intValuesMap := MapFromPtrs(intMap, IntDefaultValue{})
	fmt.Println("Map values:", intValuesMap) // 输出: map[key1:42 key2:0]

	// Map 键值提取
	keys := MapKeys(intValuesMap)
	values := MapValues(intValuesMap)
	fmt.Println("Map keys:", keys)       // 输出: [key1 key2]
	fmt.Println("Map values:", values)   // 输出: [42 0]
}

// 自定义类型的默认值提供者示例

type Person struct {
	Name string
	Age  int
}

// PersonDefaultValue 自定义类型的默认值提供者
type PersonDefaultValue struct{}

func (p PersonDefaultValue) DefaultValue() Person {
	return Person{
		Name: "Unknown",
		Age:  0,
	}
}

// 自定义默认值提供者函数
func GetPersonDefaultProvider() IDefaultValue[Person] {
	return PersonDefaultValue{}
}

func ExampleCustomTypeDefaultValue() {
	var personPtr *Person

	// 使用自定义默认值提供者
	person := FromPtr(personPtr, GetPersonDefaultProvider())
	fmt.Println("Default person:", person) // 输出: {Unknown 0}

	// 直接使用自定义默认值提供者类型
	person2 := FromPtr(personPtr, PersonDefaultValue{})
	fmt.Println("Default person (direct):", person2) // 输出: {Unknown 0}

	// 自定义类型切片的默认值处理
	personPtrs := []*Person{nil, {Name: "Alice", Age: 30}}
	persons := SliceFromPtrs(personPtrs, PersonDefaultValue{})
	fmt.Println("Person slice values:", persons) // 输出: [{Unknown 0} {Alice 30}]
}

// 高级使用示例：自定义默认值

// CustomIntDefaultValue 自定义整数默认值提供者
type CustomIntDefaultValue struct {
	Default int
}

func (c CustomIntDefaultValue) DefaultValue() int {
	return c.Default
}

func ExampleCustomDefaultValue() {
	var intPtr *int

	// 使用自定义的默认值
	customProvider := CustomIntDefaultValue{Default: -1}
	value := FromPtr(intPtr, customProvider)
	fmt.Println("Custom default value:", value) // 输出: -1

	// 在切片转换中使用自定义默认值
	intPtrs := []*int{nil, ToPtr(10), nil}
	intValues := SliceFromPtrs(intPtrs, customProvider)
	fmt.Println("Slice with custom defaults:", intValues) // 输出: [-1 10 -1]

	// 在 Map 转换中使用自定义默认值
	intMap := map[string]*int{
		"key1": nil,
		"key2": ToPtr(20),
	}
	intValuesMap := MapFromPtrs(intMap, customProvider)
	fmt.Println("Map with custom defaults:", intValuesMap) // 输出: map[key1:-1 key2:20]
}
