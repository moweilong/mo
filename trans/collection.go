//go:build go1.18
// +build go1.18

package trans

// SliceToPtrs 将值类型切片转换为指针类型切片
func SliceToPtrs[T any](slice []T) []*T {
	if slice == nil {
		return nil
	}
	result := make([]*T, len(slice))
	for i := range slice {
		result[i] = &slice[i]
	}
	return result
}

// SliceFromPtrs 将指针类型切片转换为值类型切片
func SliceFromPtrs[T any](slice []*T, defaultValue IDefaultValue[T]) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, len(slice))
	for i := range slice {
		if slice[i] != nil {
			result[i] = *slice[i]
		} else {
			result[i] = defaultValue.DefaultValue()
		}
	}
	return result
}

// MapKeysGen 获取map的键
func MapKeys[TKey comparable, TValue any](source map[TKey]TValue) []TKey {
	target := make([]TKey, 0, len(source)) // 预分配切片容量以提升性能
	for k := range source {
		target = append(target, k)
	}
	return target
}

// MapValuesGen 获取map的值
func MapValues[TKey comparable, TValue any](source map[TKey]TValue) []TValue {
	target := make([]TValue, 0, len(source)) // 预分配切片容量以提升性能
	for _, v := range source {
		target = append(target, v)
	}
	return target
}

// MapFromPtrs 将指针值的map转换为值类型的map
func MapFromPtrs[TKey comparable, TValue any](source map[TKey]*TValue, defaultValue IDefaultValue[TValue]) map[TKey]TValue {
	if source == nil {
		return nil
	}
	result := make(map[TKey]TValue, len(source))
	for k, v := range source {
		if v != nil {
			result[k] = *v
		} else {
			result[k] = defaultValue.DefaultValue()
		}
	}
	return result
}
