package strto

import stdtime "time"

// convertSliceOrZero 逐项转换字符串切片，失败元素写入目标类型的零值。
// 为保持既有行为，nil 和非 nil 空切片都返回 nil。
func convertSliceOrZero[T any](values []string, convert func(string) T) []T {
	if len(values) == 0 {
		return nil
	}
	result := make([]T, len(values))
	for index, value := range values {
		result[index] = convert(value)
	}
	return result
}

// parseSlice 解析字符串切片；任一元素失败时返回 ElementError，不返回部分结果。
func parseSlice[T any](values []string, parse func(string) (T, error)) ([]T, error) {
	if values == nil {
		return nil, nil
	}
	result := make([]T, len(values))
	for index, value := range values {
		converted, err := parse(value)
		if err != nil {
			return nil, &ElementError{Index: index, Value: value, Err: err}
		}
		result[index] = converted
	}
	return result, nil
}

// BoolsOrFalse 逐项解析字符串切片；失败元素写入 false。
func (s strto) BoolsOrFalse(values []string) []bool {
	return convertSliceOrZero(values, s.BoolOrFalse)
}

// IntsOrZero 逐项解析字符串切片为 int；失败元素写入 0。
func (s strto) IntsOrZero(values []string) []int {
	return convertSliceOrZero(values, s.IntOrZero)
}

// Int8sOrZero 逐项解析字符串切片为 int8；失败元素写入 0。
func (s strto) Int8sOrZero(values []string) []int8 {
	return convertSliceOrZero(values, s.Int8OrZero)
}

// Int16sOrZero 逐项解析字符串切片为 int16；失败元素写入 0。
func (s strto) Int16sOrZero(values []string) []int16 {
	return convertSliceOrZero(values, s.Int16OrZero)
}

// Int32sOrZero 逐项解析字符串切片为 int32；失败元素写入 0。
func (s strto) Int32sOrZero(values []string) []int32 {
	return convertSliceOrZero(values, s.Int32OrZero)
}

// Int64sOrZero 逐项解析字符串切片为 int64；失败元素写入 0。
func (s strto) Int64sOrZero(values []string) []int64 {
	return convertSliceOrZero(values, s.Int64OrZero)
}

// UintsOrZero 逐项解析字符串切片为 uint；失败元素写入 0。
func (s strto) UintsOrZero(values []string) []uint {
	return convertSliceOrZero(values, s.UintOrZero)
}

// Uint8sOrZero 逐项解析字符串切片为 uint8；失败元素写入 0。
func (s strto) Uint8sOrZero(values []string) []uint8 {
	return convertSliceOrZero(values, s.Uint8OrZero)
}

// Uint16sOrZero 逐项解析字符串切片为 uint16；失败元素写入 0。
func (s strto) Uint16sOrZero(values []string) []uint16 {
	return convertSliceOrZero(values, s.Uint16OrZero)
}

// Uint32sOrZero 逐项解析字符串切片为 uint32；失败元素写入 0。
func (s strto) Uint32sOrZero(values []string) []uint32 {
	return convertSliceOrZero(values, s.Uint32OrZero)
}

// Uint64sOrZero 逐项解析字符串切片为 uint64；失败元素写入 0。
func (s strto) Uint64sOrZero(values []string) []uint64 {
	return convertSliceOrZero(values, s.Uint64OrZero)
}

// Float32sOrZero 逐项解析字符串切片为有限 float32；失败元素写入 0。
func (s strto) Float32sOrZero(values []string) []float32 {
	return convertSliceOrZero(values, s.Float32OrZero)
}

// Float64sOrZero 逐项解析字符串切片为有限 float64；失败元素写入 0。
func (s strto) Float64sOrZero(values []string) []float64 {
	return convertSliceOrZero(values, s.Float64OrZero)
}

// DurationsOrZero 逐项解析字符串切片为 time.Duration；失败元素写入 0。
func (s strto) DurationsOrZero(values []string) []stdtime.Duration {
	return convertSliceOrZero(values, s.DurationOrZero)
}

// ParseBools 将字符串切片解析为 bool 切片；任一元素失败时返回 ElementError。
func (s strto) ParseBools(values []string) ([]bool, error) {
	return parseSlice(values, s.ParseBool)
}

// ParseInts 将字符串切片解析为 int 切片；任一元素失败时返回 ElementError。
func (s strto) ParseInts(values []string) ([]int, error) {
	return parseSlice(values, s.ParseInt)
}

// ParseInt8s 将字符串切片解析为 int8 切片；任一元素失败时返回 ElementError。
func (s strto) ParseInt8s(values []string) ([]int8, error) {
	return parseSlice(values, s.ParseInt8)
}

// ParseInt16s 将字符串切片解析为 int16 切片；任一元素失败时返回 ElementError。
func (s strto) ParseInt16s(values []string) ([]int16, error) {
	return parseSlice(values, s.ParseInt16)
}

// ParseInt32s 将字符串切片解析为 int32 切片；任一元素失败时返回 ElementError。
func (s strto) ParseInt32s(values []string) ([]int32, error) {
	return parseSlice(values, s.ParseInt32)
}

// ParseInt64s 将字符串切片解析为 int64 切片；任一元素失败时返回 ElementError。
func (s strto) ParseInt64s(values []string) ([]int64, error) {
	return parseSlice(values, s.ParseInt64)
}

// ParseUints 将字符串切片解析为 uint 切片；任一元素失败时返回 ElementError。
func (s strto) ParseUints(values []string) ([]uint, error) {
	return parseSlice(values, s.ParseUint)
}

// ParseUint8s 将字符串切片解析为 uint8 切片；任一元素失败时返回 ElementError。
func (s strto) ParseUint8s(values []string) ([]uint8, error) {
	return parseSlice(values, s.ParseUint8)
}

// ParseUint16s 将字符串切片解析为 uint16 切片；任一元素失败时返回 ElementError。
func (s strto) ParseUint16s(values []string) ([]uint16, error) {
	return parseSlice(values, s.ParseUint16)
}

// ParseUint32s 将字符串切片解析为 uint32 切片；任一元素失败时返回 ElementError。
func (s strto) ParseUint32s(values []string) ([]uint32, error) {
	return parseSlice(values, s.ParseUint32)
}

// ParseUint64s 将字符串切片解析为 uint64 切片；任一元素失败时返回 ElementError。
func (s strto) ParseUint64s(values []string) ([]uint64, error) {
	return parseSlice(values, s.ParseUint64)
}

// ParseFiniteFloat32s 将字符串切片解析为有限 float32 切片；任一元素失败时返回 ElementError。
func (s strto) ParseFiniteFloat32s(values []string) ([]float32, error) {
	return parseSlice(values, s.ParseFiniteFloat32)
}

// ParseFiniteFloat64s 将字符串切片解析为有限 float64 切片；任一元素失败时返回 ElementError。
func (s strto) ParseFiniteFloat64s(values []string) ([]float64, error) {
	return parseSlice(values, s.ParseFiniteFloat64)
}

// ParseDurations 将字符串切片解析为 time.Duration 切片；任一元素失败时返回 ElementError。
func (s strto) ParseDurations(values []string) ([]stdtime.Duration, error) {
	return parseSlice(values, s.ParseDuration)
}
