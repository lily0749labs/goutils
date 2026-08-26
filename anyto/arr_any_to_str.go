package anyto

// AnySliceToStrings 将任意值切片转换为字符串切片。
func (f anyto) AnySliceToStrings(values []any) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// IntSliceToStrings 将 int 切片转换为字符串切片。
func (f anyto) IntSliceToStrings(values []int) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Int8SliceToStrings 将 int8 切片转换为字符串切片。
func (f anyto) Int8SliceToStrings(values []int8) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Int16SliceToStrings 将 int16 切片转换为字符串切片。
func (f anyto) Int16SliceToStrings(values []int16) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Int32SliceToStrings 将 int32 切片转换为字符串切片。
func (f anyto) Int32SliceToStrings(values []int32) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Int64SliceToStrings 将 int64 切片转换为字符串切片。
func (f anyto) Int64SliceToStrings(values []int64) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// UintSliceToStrings 将 uint 切片转换为字符串切片。
func (f anyto) UintSliceToStrings(values []uint) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Uint8SliceToStrings 将 uint8 切片转换为字符串切片。
func (f anyto) Uint8SliceToStrings(values []uint8) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Uint16SliceToStrings 将 uint16 切片转换为字符串切片。
func (f anyto) Uint16SliceToStrings(values []uint16) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Uint32SliceToStrings 将 uint32 切片转换为字符串切片。
func (f anyto) Uint32SliceToStrings(values []uint32) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Uint64SliceToStrings 将 uint64 切片转换为字符串切片。
func (f anyto) Uint64SliceToStrings(values []uint64) (result []string) {
	for _, value := range values {
		result = append(result, f.String(value))
	}
	return result
}

// Deprecated: 使用 AnyTo.AnySliceToStrings。
func ArrAnyToStr(values []any) []string { return AnyTo.AnySliceToStrings(values) }

// Deprecated: 使用 AnyTo.IntSliceToStrings。
func ArrIntToStr(values []int) []string { return AnyTo.IntSliceToStrings(values) }

// Deprecated: 使用 AnyTo.Int8SliceToStrings。
func ArrInt8ToStr(values []int8) []string { return AnyTo.Int8SliceToStrings(values) }

// Deprecated: 使用 AnyTo.Int16SliceToStrings。
func ArrInt16ToStr(values []int16) []string { return AnyTo.Int16SliceToStrings(values) }

// Deprecated: 使用 AnyTo.Int32SliceToStrings。
func ArrInt32ToStr(values []int32) []string { return AnyTo.Int32SliceToStrings(values) }

// Deprecated: 使用 AnyTo.Int64SliceToStrings。
func ArrInt64ToStr(values []int64) []string { return AnyTo.Int64SliceToStrings(values) }

// Deprecated: 使用 AnyTo.UintSliceToStrings。
func ArrUintToStr(values []uint) []string { return AnyTo.UintSliceToStrings(values) }

// Deprecated: 使用 AnyTo.Uint8SliceToStrings。
func ArrUint8ToStr(values []uint8) []string { return AnyTo.Uint8SliceToStrings(values) }

// Deprecated: 使用 AnyTo.Uint16SliceToStrings。
func ArrUint16ToStr(values []uint16) []string { return AnyTo.Uint16SliceToStrings(values) }

// Deprecated: 使用 AnyTo.Uint32SliceToStrings。
func ArrUint32ToStr(values []uint32) []string { return AnyTo.Uint32SliceToStrings(values) }

// Deprecated: 使用 AnyTo.Uint64SliceToStrings。
func ArrUint64ToStr(values []uint64) []string { return AnyTo.Uint64SliceToStrings(values) }
