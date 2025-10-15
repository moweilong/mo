//go:build go1.18
// +build go1.18

package trans

import "time"

// ToPtr 返回值的指针
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr 使用默认值提供者将指针转换为值
func FromPtr[T any](p *T, provider IDefaultValue[T]) T {
	if p == nil {
		return provider.DefaultValue()
	}
	return *p
}

// StringValue 将字符串指针转换为值
func StringValue(a *string) string {
	return FromPtr(a, StringDefaultValue{})
}

// IntValue 将整数指针转换为值
func IntValue(a *int) int {
	return FromPtr(a, IntDefaultValue{})
}

// Int8Value 将整数指针转换为值
func Int8Value(a *int8) int8 {
	return FromPtr(a, Int8DefaultValue{})
}

// Int16Value 将整数指针转换为值
func Int16Value(a *int16) int16 {
	return FromPtr(a, Int16DefaultValue{})
}

// Int32Value 将整数指针转换为值
func Int32Value(a *int32) int32 {
	return FromPtr(a, Int32DefaultValue{})
}

// Int64Value 将整数指针转换为值
func Int64Value(a *int64) int64 {
	return FromPtr(a, Int64DefaultValue{})
}

// UintValue 将无符号整数指针转换为值
func UintValue(a *uint) uint {
	return FromPtr(a, UintDefaultValue{})
}

// Uint8Value 将无符号整数指针转换为值
func Uint8Value(a *uint8) uint8 {
	return FromPtr(a, Uint8DefaultValue{})
}

// Uint16Value 将无符号整数指针转换为值
func Uint16Value(a *uint16) uint16 {
	return FromPtr(a, Uint16DefaultValue{})
}

// Uint32Value 将无符号整数指针转换为值
func Uint32Value(a *uint32) uint32 {
	return FromPtr(a, Uint32DefaultValue{})
}

// Uint64Value 将无符号整数指针转换为值
func Uint64Value(a *uint64) uint64 {
	return FromPtr(a, Uint64DefaultValue{})
}

// BoolValue 将布尔指针转换为值
func BoolValue(a *bool) bool {
	return FromPtr(a, BoolDefaultValue{})
}

// Float32Value 将浮点数指针转换为值
func Float32Value(a *float32) float32 {
	return FromPtr(a, Float32DefaultValue{})
}

// Float64Value 将浮点数指针转换为值
func Float64Value(a *float64) float64 {
	return FromPtr(a, Float64DefaultValue{})
}

// TimeValue 将时间指针转换为值
func TimeValue(a *time.Time) time.Time {
	return FromPtr(a, TimeDefaultValue{})
}

// BatchPtrs 批量将值转换为指针
func BatchPtrs[T any](values ...T) []*T {
	result := make([]*T, len(values))
	for i := range values {
		result[i] = &values[i]
	}
	return result
}

// IDefaultValue 为每种类型提供默认值处理器的接口
type IDefaultValue[T any] interface {
	DefaultValue() T
}

// StringDefaultValue 字符串类型的默认值提供者
type StringDefaultValue struct{}

func (s StringDefaultValue) DefaultValue() string {
	return ""
}

// IntDefaultValue 整数类型的默认值提供者
type IntDefaultValue struct{}

func (i IntDefaultValue) DefaultValue() int {
	return 0
}

// Int8DefaultValue 整数类型的默认值提供者
type Int8DefaultValue struct{}

func (i Int8DefaultValue) DefaultValue() int8 {
	return 0
}

// Int16DefaultValue 整数类型的默认值提供者
type Int16DefaultValue struct{}

func (i Int16DefaultValue) DefaultValue() int16 {
	return 0
}

// Int32DefaultValue 整数类型的默认值提供者
type Int32DefaultValue struct{}

func (i Int32DefaultValue) DefaultValue() int32 {
	return 0
}

// Int64DefaultValue 整数类型的默认值提供者
type Int64DefaultValue struct{}

func (i Int64DefaultValue) DefaultValue() int64 {
	return 0
}

// UintDefaultValue 无符号整数类型的默认值提供者
type UintDefaultValue struct{}

func (u UintDefaultValue) DefaultValue() uint {
	return 0
}

// Uint8DefaultValue 无符号整数类型的默认值提供者
type Uint8DefaultValue struct{}

func (u Uint8DefaultValue) DefaultValue() uint8 {
	return 0
}

// Uint16DefaultValue 无符号整数类型的默认值提供者
type Uint16DefaultValue struct{}

func (u Uint16DefaultValue) DefaultValue() uint16 {
	return 0
}

// Uint32DefaultValue 无符号整数类型的默认值提供者
type Uint32DefaultValue struct{}

func (u Uint32DefaultValue) DefaultValue() uint32 {
	return 0
}

// Uint64DefaultValue 无符号整数类型的默认值提供者
type Uint64DefaultValue struct{}

func (u Uint64DefaultValue) DefaultValue() uint64 {
	return 0
}

// BoolDefaultValue 布尔类型的默认值提供者
type BoolDefaultValue struct{}

func (b BoolDefaultValue) DefaultValue() bool {
	return false
}

// Float32DefaultValue 浮点数类型的默认值提供者
type Float32DefaultValue struct{}

func (f Float32DefaultValue) DefaultValue() float32 {
	return 0.0
}

// Float64DefaultValue 浮点数类型的默认值提供者
type Float64DefaultValue struct{}

func (f Float64DefaultValue) DefaultValue() float64 {
	return 0.0
}

// TimeDefaultValue 时间类型的默认值提供者
type TimeDefaultValue struct{}

func (t TimeDefaultValue) DefaultValue() time.Time {
	return time.Now()
}
