// +build go1.18

package trans

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	// 测试基本类型
	str := "hello"
	strPtr := ToPtr(str)
	assert.NotNil(t, strPtr)
	assert.Equal(t, str, *strPtr)

	// 测试数值类型
	num := 42
	numPtr := ToPtr(num)
	assert.NotNil(t, numPtr)
	assert.Equal(t, num, *numPtr)

	// 测试布尔类型
	b := true
	bPtr := ToPtr(b)
	assert.NotNil(t, bPtr)
	assert.Equal(t, b, *bPtr)

	// 测试复杂类型
	now := time.Now()
	nowPtr := ToPtr(now)
	assert.NotNil(t, nowPtr)
	assert.Equal(t, now, *nowPtr)
}

func TestFromPtr(t *testing.T) {
	// 测试非nil指针
	s := "test"
	sPtr := &s
	result := FromPtr(sPtr, StringDefaultValue{})
	assert.Equal(t, s, result)

	// 测试nil指针
	var nilPtr *string
	result = FromPtr(nilPtr, StringDefaultValue{})
	assert.Equal(t, "", result)

	// 测试自定义默认值 - 使用StringDefaultValue
	assert.Equal(t, "", FromPtr(nilPtr, StringDefaultValue{}))
}

func TestStringValue(t *testing.T) {
	// 测试非nil指针
	s := "test"
	sPtr := &s
	result := StringValue(sPtr)
	assert.Equal(t, s, result)

	// 测试nil指针
	var nilPtr *string
	result = StringValue(nilPtr)
	assert.Equal(t, "", result)
}

func TestIntValue(t *testing.T) {
	// 测试非nil指针
	n := 42
	nPtr := &n
	result := IntValue(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *int
	result = IntValue(nilPtr)
	assert.Equal(t, 0, result)
}

func TestBoolValue(t *testing.T) {
	// 测试非nil指针 - true值
	b := true
	bPtr := &b
	result := BoolValue(bPtr)
	assert.Equal(t, b, result)

	// 测试非nil指针 - false值
	b = false
	bPtr = &b
	result = BoolValue(bPtr)
	assert.Equal(t, b, result)

	// 测试nil指针
	var nilPtr *bool
	result = BoolValue(nilPtr)
	assert.Equal(t, false, result)
}

func TestInt8Value(t *testing.T) {
	// 测试非nil指针
	var n int8 = 42
	var nPtr *int8 = &n
	result := Int8Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *int8
	result = Int8Value(nilPtr)
	assert.Equal(t, int8(0), result)
}

func TestInt16Value(t *testing.T) {
	// 测试非nil指针
	var n int16 = 42
	var nPtr *int16 = &n
	result := Int16Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *int16
	result = Int16Value(nilPtr)
	assert.Equal(t, int16(0), result)
}

func TestInt32Value(t *testing.T) {
	// 测试非nil指针
	var n int32 = 42
	var nPtr *int32 = &n
	result := Int32Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *int32
	result = Int32Value(nilPtr)
	assert.Equal(t, int32(0), result)
}

func TestInt64Value(t *testing.T) {
	// 测试非nil指针
	var n int64 = 42
	var nPtr *int64 = &n
	result := Int64Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *int64
	result = Int64Value(nilPtr)
	assert.Equal(t, int64(0), result)
}

func TestUintValue(t *testing.T) {
	// 测试非nil指针
	var n uint = 42
	var nPtr *uint = &n
	result := UintValue(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *uint
	result = UintValue(nilPtr)
	assert.Equal(t, uint(0), result)
}

func TestUint8Value(t *testing.T) {
	// 测试非nil指针
	var n uint8 = 42
	var nPtr *uint8 = &n
	result := Uint8Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *uint8
	result = Uint8Value(nilPtr)
	assert.Equal(t, uint8(0), result)
}

func TestUint16Value(t *testing.T) {
	// 测试非nil指针
	var n uint16 = 42
	var nPtr *uint16 = &n
	result := Uint16Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *uint16
	result = Uint16Value(nilPtr)
	assert.Equal(t, uint16(0), result)
}

func TestUint32Value(t *testing.T) {
	// 测试非nil指针
	var n uint32 = 42
	var nPtr *uint32 = &n
	result := Uint32Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *uint32
	result = Uint32Value(nilPtr)
	assert.Equal(t, uint32(0), result)
}

func TestUint64Value(t *testing.T) {
	// 测试非nil指针
	var n uint64 = 42
	var nPtr *uint64 = &n
	result := Uint64Value(nPtr)
	assert.Equal(t, n, result)

	// 测试nil指针
	var nilPtr *uint64
	result = Uint64Value(nilPtr)
	assert.Equal(t, uint64(0), result)
}

func TestFloat32Value(t *testing.T) {
	// 测试非nil指针
	var f float32 = 3.14
	var fPtr *float32 = &f
	result := Float32Value(fPtr)
	assert.Equal(t, f, result)

	// 测试nil指针
	var nilPtr *float32
	result = Float32Value(nilPtr)
	assert.Equal(t, float32(0.0), result)
}

func TestFloat64Value(t *testing.T) {
	// 测试非nil指针
	var f float64 = 3.14
	var fPtr *float64 = &f
	result := Float64Value(fPtr)
	assert.Equal(t, f, result)

	// 测试nil指针
	var nilPtr *float64
	result = Float64Value(nilPtr)
	assert.Equal(t, float64(0.0), result)
}

func TestTimeValue(t *testing.T) {
	// 测试非nil指针
	tNow := time.Now()
	var tPtr *time.Time = &tNow
	result := TimeValue(tPtr)
	assert.Equal(t, tNow, result)

	// 测试nil指针
	var nilPtr *time.Time
	result = TimeValue(nilPtr)
	// 检查返回的是当前时间，不是nil值
	assert.NotZero(t, result)
}

func TestBatchPtrs(t *testing.T) {
	// 测试字符串切片
	strs := []string{"a", "b", "c"}
	strPtrs := BatchPtrs(strs...)
	assert.Len(t, strPtrs, 3)
	for i, ptr := range strPtrs {
		assert.NotNil(t, ptr)
		assert.Equal(t, strs[i], *ptr)
	}

	// 测试整数切片
	nums := []int{1, 2, 3}
	numPtrs := BatchPtrs(nums...)
	assert.Len(t, numPtrs, 3)
	for i, ptr := range numPtrs {
		assert.NotNil(t, ptr)
		assert.Equal(t, nums[i], *ptr)
	}

	// 测试空参数
	emptyPtrs := BatchPtrs[int]()
	assert.Len(t, emptyPtrs, 0)
}

func TestDefaultValueProviders(t *testing.T) {
	// 测试StringDefaultValue
	strProvider := StringDefaultValue{}
	assert.Equal(t, "", strProvider.DefaultValue())

	// 测试IntDefaultValue
	intProvider := IntDefaultValue{}
	assert.Equal(t, 0, intProvider.DefaultValue())

	// 测试BoolDefaultValue
	boolProvider := BoolDefaultValue{}
	assert.Equal(t, false, boolProvider.DefaultValue())

	// 测试Float32DefaultValue
	float32Provider := Float32DefaultValue{}
	assert.Equal(t, float32(0.0), float32Provider.DefaultValue())

	// 测试Float64DefaultValue
	float64Provider := Float64DefaultValue{}
	assert.Equal(t, float64(0.0), float64Provider.DefaultValue())

	// 测试TimeDefaultValue
	timeProvider := TimeDefaultValue{}
	timeVal := timeProvider.DefaultValue()
	assert.NotZero(t, timeVal)
}

// 自定义类型定义
type CustomType struct {
	Field1 string
	Field2 int
}

type CustomTypeDefaultValue struct{}

func (c CustomTypeDefaultValue) DefaultValue() CustomType {
	return CustomType{
		Field1: "default",
		Field2: -1,
	}
}

// 自定义类型测试

func TestFromPtrWithCustomType(t *testing.T) {
	// 测试非nil指针
	customVal := CustomType{"test", 42}
	customPtr := &customVal
	result := FromPtr(customPtr, CustomTypeDefaultValue{})
	assert.Equal(t, customVal, result)

	// 测试nil指针
	var nilPtr *CustomType
	result = FromPtr(nilPtr, CustomTypeDefaultValue{})
	assert.Equal(t, CustomType{"default", -1}, result)
}