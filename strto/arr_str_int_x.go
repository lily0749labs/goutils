package strto

// Ints 将字符串切片转换为 int 切片。
func (f facade) Ints(values []string) (result []int) {
	for _, value := range values {
		result = append(result, f.Int(value))
	}
	return result
}

// Int8s 将字符串切片转换为 int8 切片。
func (f facade) Int8s(values []string) (result []int8) {
	for _, value := range values {
		result = append(result, f.Int8(value))
	}
	return result
}

// Int16s 将字符串切片转换为 int16 切片。
func (f facade) Int16s(values []string) (result []int16) {
	for _, value := range values {
		result = append(result, f.Int16(value))
	}
	return result
}

// Int32s 将字符串切片转换为 int32 切片。
func (f facade) Int32s(values []string) (result []int32) {
	for _, value := range values {
		result = append(result, f.Int32(value))
	}
	return result
}

// Int64s 将字符串切片转换为 int64 切片。
func (f facade) Int64s(values []string) (result []int64) {
	for _, value := range values {
		result = append(result, f.Int64(value))
	}
	return result
}

// Deprecated: 使用 StrTo.Ints。
func ArrStrToInt(values []string) []int { return StrTo.Ints(values) }

// Deprecated: 使用 StrTo.Int8s。
func ArrStrToInt8(values []string) []int8 { return StrTo.Int8s(values) }

// Deprecated: 使用 StrTo.Int16s。
func ArrStrToInt16(values []string) []int16 { return StrTo.Int16s(values) }

// Deprecated: 使用 StrTo.Int32s。
func ArrStrToInt32(values []string) []int32 { return StrTo.Int32s(values) }

// Deprecated: 使用 StrTo.Int64s。
func ArrStrToInt64(values []string) []int64 { return StrTo.Int64s(values) }
