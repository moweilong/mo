//go:build go1.18
// +build go1.18

package trans

import "github.com/google/uuid"

// ToUuidE 将字符串解析为UUID，返回错误信息
func ToUuidE(str string) (uuid.UUID, error) {
	return uuid.Parse(str)
}

// ToUuidPtrE 将字符串指针解析为UUID指针，返回错误信息
func ToUuidPtrE(str *string) (*uuid.UUID, error) {
	if str == nil {
		return nil, nil
	}
	id, err := uuid.Parse(*str)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

// ToUuid 将字符串转换为UUID，失败时返回nil值
func ToUuid(s string) uuid.UUID {
	uid, _ := uuid.Parse(s)
	return uid
}

// ToUuidPtr 将字符串转换为UUID指针，失败时返回nil
func ToUuidPtr(s string) *uuid.UUID {
	uid, err := uuid.Parse(s)
	if err != nil {
		return nil
	}
	return &uid
}

// ToStringPtr 将UUID转换为字符串指针
func ToStringPtr(uid uuid.UUID) *string {
	s := uid.String()
	return &s
}

// UUIDValue 将UUID指针转换为值
func UUIDValue(u *uuid.UUID) uuid.UUID {
	if u == nil {
		return UUIDDefaultValue{}.DefaultValue()
	}
	return *u
}

// SliceToUUIDPtrs 将UUID切片转换为UUID指针切片
func SliceToUUIDPtrs(slice []uuid.UUID) []*uuid.UUID {
	return SliceToPtrs(slice)
}

// SliceFromUUIDPtrs 将UUID指针切片转换为UUID切片
func SliceFromUUIDPtrs(slice []*uuid.UUID) []uuid.UUID {
	return SliceFromPtrs(slice, UUIDDefaultValue{})
}

// UUIDDefaultValue UUID类型的默认值提供者
type UUIDDefaultValue struct{}

func (u UUIDDefaultValue) DefaultValue() uuid.UUID {
	return uuid.Nil
}
