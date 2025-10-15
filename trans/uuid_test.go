// +build go1.18

package trans

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToUuidE(t *testing.T) {
	// 测试有效的UUID字符串
	validUUIDStr := "123e4567-e89b-12d3-a456-426614174000"
	uid, err := ToUuidE(validUUIDStr)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, uid)
	assert.Equal(t, validUUIDStr, uid.String())

	// 测试无效的UUID字符串
	invalidUUIDStr := "not-a-uuid"
	invalidUid, err := ToUuidE(invalidUUIDStr)
	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, invalidUid)

	// 测试空字符串
	emptyUUIDStr := ""
	emptyUid, err := ToUuidE(emptyUUIDStr)
	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, emptyUid)
}

func TestToUuidPtrE(t *testing.T) {
	// 测试有效的UUID字符串指针
	validUUIDStr := "123e4567-e89b-12d3-a456-426614174000"
	uidPtr, err := ToUuidPtrE(&validUUIDStr)
	assert.NoError(t, err)
	assert.NotNil(t, uidPtr)
	assert.Equal(t, validUUIDStr, uidPtr.String())

	// 测试无效的UUID字符串指针
	invalidUUIDStr := "not-a-uuid"
	invalidUidPtr, err := ToUuidPtrE(&invalidUUIDStr)
	assert.Error(t, err)
	assert.Nil(t, invalidUidPtr)

	// 测试nil字符串指针
	var nilStrPtr *string
	nilUidPtr, err := ToUuidPtrE(nilStrPtr)
	assert.NoError(t, err)
	assert.Nil(t, nilUidPtr)
}

func TestToUuid(t *testing.T) {
	// 测试有效的UUID字符串
	validUUIDStr := "123e4567-e89b-12d3-a456-426614174000"
	uid := ToUuid(validUUIDStr)
	assert.NotEqual(t, uuid.Nil, uid)
	assert.Equal(t, validUUIDStr, uid.String())

	// 测试无效的UUID字符串
	invalidUUIDStr := "not-a-uuid"
	invalidUid := ToUuid(invalidUUIDStr)
	// 注意：ToUuid在解析失败时返回nil值，但uuid.Nil不是nil指针，而是一个特殊的UUID值
	assert.Equal(t, uuid.Nil, invalidUid)

	// 测试空字符串
	emptyUUIDStr := ""
	emptyUid := ToUuid(emptyUUIDStr)
	assert.Equal(t, uuid.Nil, emptyUid)
}

func TestToUuidPtr(t *testing.T) {
	// 测试有效的UUID字符串
	validUUIDStr := "123e4567-e89b-12d3-a456-426614174000"
	uidPtr := ToUuidPtr(validUUIDStr)
	assert.NotNil(t, uidPtr)
	assert.Equal(t, validUUIDStr, uidPtr.String())

	// 测试无效的UUID字符串
	invalidUUIDStr := "not-a-uuid"
	invalidUidPtr := ToUuidPtr(invalidUUIDStr)
	assert.Nil(t, invalidUidPtr)

	// 测试空字符串
	emptyUUIDStr := ""
	emptyUidPtr := ToUuidPtr(emptyUUIDStr)
	assert.Nil(t, emptyUidPtr)
}

func TestToStringPtr(t *testing.T) {
	// 测试有效的UUID
	validUUID := uuid.New()
	strPtr := ToStringPtr(validUUID)
	assert.NotNil(t, strPtr)
	assert.Equal(t, validUUID.String(), *strPtr)

	// 测试nil值UUID
	nilUUID := uuid.Nil
	nilStrPtr := ToStringPtr(nilUUID)
	assert.NotNil(t, nilStrPtr)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", *nilStrPtr)
}

func TestUUIDValue(t *testing.T) {
	// 测试非nil UUID指针
	validUUID := uuid.New()
	validUUIDPtr := &validUUID
	uid := UUIDValue(validUUIDPtr)
	assert.Equal(t, validUUID, uid)

	// 测试nil UUID指针
	var nilUUIDPtr *uuid.UUID
	nilUid := UUIDValue(nilUUIDPtr)
	assert.Equal(t, uuid.Nil, nilUid)
}

func TestSliceToUUIDPtrs(t *testing.T) {
	// 测试UUID切片
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuid3 := uuid.New()
	uuidSlice := []uuid.UUID{uuid1, uuid2, uuid3}
	ptrSlice := SliceToUUIDPtrs(uuidSlice)
	assert.Len(t, ptrSlice, 3)
	for i, ptr := range ptrSlice {
		assert.NotNil(t, ptr)
		assert.Equal(t, uuidSlice[i], *ptr)
	}

	// 测试空切片
	emptyUUIDSlice := []uuid.UUID{}
	emptyPtrSlice := SliceToUUIDPtrs(emptyUUIDSlice)
	assert.Len(t, emptyPtrSlice, 0)

	// 测试nil切片
	var nilUUIDSlice []uuid.UUID
	nilPtrSlice := SliceToUUIDPtrs(nilUUIDSlice)
	assert.Nil(t, nilPtrSlice)

	// 测试包含nil值的UUID切片
	nilUUID := uuid.Nil
	nilValueSlice := []uuid.UUID{uuid1, nilUUID, uuid3}
	nilValuePtrSlice := SliceToUUIDPtrs(nilValueSlice)
	assert.Len(t, nilValuePtrSlice, 3)
	assert.Equal(t, uuid1, *nilValuePtrSlice[0])
	assert.Equal(t, nilUUID, *nilValuePtrSlice[1])
	assert.Equal(t, uuid3, *nilValuePtrSlice[2])
}

func TestSliceFromUUIDPtrs(t *testing.T) {
	// 测试非nil指针切片
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuid3 := uuid.New()
	uuidPtrSlice := []*uuid.UUID{&uuid1, &uuid2, &uuid3}
	valueSlice := SliceFromUUIDPtrs(uuidPtrSlice)
	assert.Len(t, valueSlice, 3)
	assert.Equal(t, []uuid.UUID{uuid1, uuid2, uuid3}, valueSlice)

	// 测试包含nil指针的切片
	var nilUUIDPtr *uuid.UUID
	mixedPtrSlice := []*uuid.UUID{&uuid1, nilUUIDPtr, &uuid3}
	mixedValueSlice := SliceFromUUIDPtrs(mixedPtrSlice)
	assert.Len(t, mixedValueSlice, 3)
	assert.Equal(t, uuid1, mixedValueSlice[0])
	assert.Equal(t, uuid.Nil, mixedValueSlice[1])
	assert.Equal(t, uuid3, mixedValueSlice[2])

	// 测试空切片
	emptyPtrSlice := []*uuid.UUID{}
	emptyValueSlice := SliceFromUUIDPtrs(emptyPtrSlice)
	assert.Len(t, emptyValueSlice, 0)

	// 测试nil切片
	var nilPtrSlice []*uuid.UUID
	nilValueSlice := SliceFromUUIDPtrs(nilPtrSlice)
	assert.Nil(t, nilValueSlice)
}

func TestUUIDDefaultValue(t *testing.T) {
	// 测试UUIDDefaultValue
	uuidProvider := UUIDDefaultValue{}
	defaultUUID := uuidProvider.DefaultValue()
	assert.Equal(t, uuid.Nil, defaultUUID)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", defaultUUID.String())
}

func TestIntegrationUUIDFunctions(t *testing.T) {
	// 集成测试：从字符串到UUID指针再返回
	originalStr := "123e4567-e89b-12d3-a456-426614174000"
	uuidPtr := ToUuidPtr(originalStr)
	resultStrPtr := ToStringPtr(*uuidPtr)
	assert.NotNil(t, uuidPtr)
	assert.NotNil(t, resultStrPtr)
	assert.Equal(t, originalStr, *resultStrPtr)

	// 集成测试：处理包含nil的UUID切片
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	var nilUUIDPtr *uuid.UUID
	mixedPtrSlice := []*uuid.UUID{&uuid1, nilUUIDPtr, &uuid2}
	valueSlice := SliceFromUUIDPtrs(mixedPtrSlice)
	assert.Len(t, valueSlice, 3)
	assert.Equal(t, uuid1, valueSlice[0])
	assert.Equal(t, uuid.Nil, valueSlice[1])
	assert.Equal(t, uuid2, valueSlice[2])

	// 再转换回指针切片
	convertedPtrSlice := SliceToUUIDPtrs(valueSlice)
	assert.Len(t, convertedPtrSlice, 3)
	assert.Equal(t, uuid1, *convertedPtrSlice[0])
	assert.Equal(t, uuid.Nil, *convertedPtrSlice[1])
	assert.Equal(t, uuid2, *convertedPtrSlice[2])
}