package strto

// Uints 将字符串切片转换为 uint 切片。
func (f facade) Uints(values []string) (result []uint) {
	for _, value := range values {
		result = append(result, f.Uint(value))
	}
	return result
}

// Uint8s 将字符串切片转换为 uint8 切片。
func (f facade) Uint8s(values []string) (result []uint8) {
	for _, value := range values {
		result = append(result, f.Uint8(value))
	}
	return result
}

// Uint16s 将字符串切片转换为 uint16 切片。
func (f facade) Uint16s(values []string) (result []uint16) {
	for _, value := range values {
		result = append(result, f.Uint16(value))
	}
	return result
}

// Uint32s 将字符串切片转换为 uint32 切片。
func (f facade) Uint32s(values []string) (result []uint32) {
	for _, value := range values {
		result = append(result, f.Uint32(value))
	}
	return result
}

// Uint64s 将字符串切片转换为 uint64 切片。
func (f facade) Uint64s(values []string) (result []uint64) {
	for _, value := range values {
		result = append(result, f.Uint64(value))
	}
	return result
}

// Deprecated: 使用 StrTo.Uints。
func ArrStrToUint(values []string) []uint { return StrTo.Uints(values) }

// Deprecated: 使用 StrTo.Uint8s。
func ArrStrToUint8(values []string) []uint8 { return StrTo.Uint8s(values) }

// Deprecated: 使用 StrTo.Uint16s。
func ArrStrToUint16(values []string) []uint16 { return StrTo.Uint16s(values) }

// Deprecated: 使用 StrTo.Uint32s。
func ArrStrToUint32(values []string) []uint32 { return StrTo.Uint32s(values) }

// Deprecated: 使用 StrTo.Uint64s。
func ArrStrToUint64(values []string) []uint64 { return StrTo.Uint64s(values) }
