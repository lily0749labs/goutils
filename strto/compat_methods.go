package strto

import stdtime "time"

// StrictInt 是旧版兼容入口。
// Deprecated: 请使用 ParseInt。
func (s strto) StrictInt(value string) (int, error) { return s.ParseInt(value) }

// StrictInt8 是旧版兼容入口。
// Deprecated: 请使用 ParseInt8。
func (s strto) StrictInt8(value string) (int8, error) { return s.ParseInt8(value) }

// StrictInt16 是旧版兼容入口。
// Deprecated: 请使用 ParseInt16。
func (s strto) StrictInt16(value string) (int16, error) { return s.ParseInt16(value) }

// StrictInt32 是旧版兼容入口。
// Deprecated: 请使用 ParseInt32。
func (s strto) StrictInt32(value string) (int32, error) { return s.ParseInt32(value) }

// StrictInt64 是旧版兼容入口。
// Deprecated: 请使用 ParseInt64。
func (s strto) StrictInt64(value string) (int64, error) { return s.ParseInt64(value) }

// StrictUint 是旧版兼容入口。
// Deprecated: 请使用 ParseUint。
func (s strto) StrictUint(value string) (uint, error) { return s.ParseUint(value) }

// StrictUint8 是旧版兼容入口。
// Deprecated: 请使用 ParseUint8。
func (s strto) StrictUint8(value string) (uint8, error) { return s.ParseUint8(value) }

// StrictUint16 是旧版兼容入口。
// Deprecated: 请使用 ParseUint16。
func (s strto) StrictUint16(value string) (uint16, error) { return s.ParseUint16(value) }

// StrictUint32 是旧版兼容入口。
// Deprecated: 请使用 ParseUint32。
func (s strto) StrictUint32(value string) (uint32, error) { return s.ParseUint32(value) }

// StrictUint64 是旧版兼容入口。
// Deprecated: 请使用 ParseUint64。
func (s strto) StrictUint64(value string) (uint64, error) { return s.ParseUint64(value) }

// StrictBool 是旧版兼容入口。
// Deprecated: 请使用 ParseBool。
func (s strto) StrictBool(value string) (bool, error) { return s.ParseBool(value) }

// StrictFloat32 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat32。
func (s strto) StrictFloat32(value string) (float32, error) {
	return s.ParseFiniteFloat32(value)
}

// StrictFloat64 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat64。
func (s strto) StrictFloat64(value string) (float64, error) {
	return s.ParseFiniteFloat64(value)
}

// StrictDuration 是旧版兼容入口。
// Deprecated: 请使用 ParseDuration。
func (s strto) StrictDuration(value string) (stdtime.Duration, error) {
	return s.ParseDuration(value)
}

// Int 是旧版兼容入口。
// Deprecated: 请使用 IntOrZero。
func (s strto) Int(value string) int { return s.IntOrZero(value) }

// Int8 是旧版兼容入口。
// Deprecated: 请使用 Int8OrZero。
func (s strto) Int8(value string) int8 { return s.Int8OrZero(value) }

// Int16 是旧版兼容入口。
// Deprecated: 请使用 Int16OrZero。
func (s strto) Int16(value string) int16 { return s.Int16OrZero(value) }

// Int32 是旧版兼容入口。
// Deprecated: 请使用 Int32OrZero。
func (s strto) Int32(value string) int32 { return s.Int32OrZero(value) }

// Int64 是旧版兼容入口。
// Deprecated: 请使用 Int64OrZero。
func (s strto) Int64(value string) int64 { return s.Int64OrZero(value) }

// Uint 是旧版兼容入口。
// Deprecated: 请使用 UintOrZero。
func (s strto) Uint(value string) uint { return s.UintOrZero(value) }

// Uint8 是旧版兼容入口。
// Deprecated: 请使用 Uint8OrZero。
func (s strto) Uint8(value string) uint8 { return s.Uint8OrZero(value) }

// Uint16 是旧版兼容入口。
// Deprecated: 请使用 Uint16OrZero。
func (s strto) Uint16(value string) uint16 { return s.Uint16OrZero(value) }

// Uint32 是旧版兼容入口。
// Deprecated: 请使用 Uint32OrZero。
func (s strto) Uint32(value string) uint32 { return s.Uint32OrZero(value) }

// Uint64 是旧版兼容入口。
// Deprecated: 请使用 Uint64OrZero。
func (s strto) Uint64(value string) uint64 { return s.Uint64OrZero(value) }

// Bool 是旧版兼容入口。
// Deprecated: 请使用 BoolOrFalse。
func (s strto) Bool(value string) bool { return s.BoolOrFalse(value) }

// Float32 是旧版兼容入口。
// Deprecated: 请使用 Float32OrZero。
func (s strto) Float32(value string) float32 { return s.Float32OrZero(value) }

// Float64 是旧版兼容入口。
// Deprecated: 请使用 Float64OrZero。
func (s strto) Float64(value string) float64 { return s.Float64OrZero(value) }

// Duration 是旧版兼容入口。
// Deprecated: 请使用 DurationOrZero。
func (s strto) Duration(value string) stdtime.Duration { return s.DurationOrZero(value) }

// StrictBools 是旧版兼容入口。
// Deprecated: 请使用 ParseBools。
func (s strto) StrictBools(values []string) ([]bool, error) { return s.ParseBools(values) }

// StrictInts 是旧版兼容入口。
// Deprecated: 请使用 ParseInts。
func (s strto) StrictInts(values []string) ([]int, error) { return s.ParseInts(values) }

// StrictInt8s 是旧版兼容入口。
// Deprecated: 请使用 ParseInt8s。
func (s strto) StrictInt8s(values []string) ([]int8, error) { return s.ParseInt8s(values) }

// StrictInt16s 是旧版兼容入口。
// Deprecated: 请使用 ParseInt16s。
func (s strto) StrictInt16s(values []string) ([]int16, error) { return s.ParseInt16s(values) }

// StrictInt32s 是旧版兼容入口。
// Deprecated: 请使用 ParseInt32s。
func (s strto) StrictInt32s(values []string) ([]int32, error) { return s.ParseInt32s(values) }

// StrictInt64s 是旧版兼容入口。
// Deprecated: 请使用 ParseInt64s。
func (s strto) StrictInt64s(values []string) ([]int64, error) { return s.ParseInt64s(values) }

// StrictUints 是旧版兼容入口。
// Deprecated: 请使用 ParseUints。
func (s strto) StrictUints(values []string) ([]uint, error) { return s.ParseUints(values) }

// StrictUint8s 是旧版兼容入口。
// Deprecated: 请使用 ParseUint8s。
func (s strto) StrictUint8s(values []string) ([]uint8, error) { return s.ParseUint8s(values) }

// StrictUint16s 是旧版兼容入口。
// Deprecated: 请使用 ParseUint16s。
func (s strto) StrictUint16s(values []string) ([]uint16, error) { return s.ParseUint16s(values) }

// StrictUint32s 是旧版兼容入口。
// Deprecated: 请使用 ParseUint32s。
func (s strto) StrictUint32s(values []string) ([]uint32, error) { return s.ParseUint32s(values) }

// StrictUint64s 是旧版兼容入口。
// Deprecated: 请使用 ParseUint64s。
func (s strto) StrictUint64s(values []string) ([]uint64, error) { return s.ParseUint64s(values) }

// StrictFloat32s 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat32s。
func (s strto) StrictFloat32s(values []string) ([]float32, error) {
	return s.ParseFiniteFloat32s(values)
}

// StrictFloat64s 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat64s。
func (s strto) StrictFloat64s(values []string) ([]float64, error) {
	return s.ParseFiniteFloat64s(values)
}

// StrictDurations 是旧版兼容入口。
// Deprecated: 请使用 ParseDurations。
func (s strto) StrictDurations(values []string) ([]stdtime.Duration, error) {
	return s.ParseDurations(values)
}

// Bools 是旧版兼容入口。
// Deprecated: 请使用 BoolsOrFalse。
func (s strto) Bools(values []string) []bool { return s.BoolsOrFalse(values) }

// Ints 是旧版兼容入口。
// Deprecated: 请使用 IntsOrZero。
func (s strto) Ints(values []string) []int { return s.IntsOrZero(values) }

// Int8s 是旧版兼容入口。
// Deprecated: 请使用 Int8sOrZero。
func (s strto) Int8s(values []string) []int8 { return s.Int8sOrZero(values) }

// Int16s 是旧版兼容入口。
// Deprecated: 请使用 Int16sOrZero。
func (s strto) Int16s(values []string) []int16 { return s.Int16sOrZero(values) }

// Int32s 是旧版兼容入口。
// Deprecated: 请使用 Int32sOrZero。
func (s strto) Int32s(values []string) []int32 { return s.Int32sOrZero(values) }

// Int64s 是旧版兼容入口。
// Deprecated: 请使用 Int64sOrZero。
func (s strto) Int64s(values []string) []int64 { return s.Int64sOrZero(values) }

// Uints 是旧版兼容入口。
// Deprecated: 请使用 UintsOrZero。
func (s strto) Uints(values []string) []uint { return s.UintsOrZero(values) }

// Uint8s 是旧版兼容入口。
// Deprecated: 请使用 Uint8sOrZero。
func (s strto) Uint8s(values []string) []uint8 { return s.Uint8sOrZero(values) }

// Uint16s 是旧版兼容入口。
// Deprecated: 请使用 Uint16sOrZero。
func (s strto) Uint16s(values []string) []uint16 { return s.Uint16sOrZero(values) }

// Uint32s 是旧版兼容入口。
// Deprecated: 请使用 Uint32sOrZero。
func (s strto) Uint32s(values []string) []uint32 { return s.Uint32sOrZero(values) }

// Uint64s 是旧版兼容入口。
// Deprecated: 请使用 Uint64sOrZero。
func (s strto) Uint64s(values []string) []uint64 { return s.Uint64sOrZero(values) }

// Float32s 是旧版兼容入口。
// Deprecated: 请使用 Float32sOrZero。
func (s strto) Float32s(values []string) []float32 { return s.Float32sOrZero(values) }

// Float64s 是旧版兼容入口。
// Deprecated: 请使用 Float64sOrZero。
func (s strto) Float64s(values []string) []float64 { return s.Float64sOrZero(values) }

// Durations 是旧版兼容入口。
// Deprecated: 请使用 DurationsOrZero。
func (s strto) Durations(values []string) []stdtime.Duration { return s.DurationsOrZero(values) }

// CommaInts 是旧版兼容入口。
// Deprecated: 请使用 ParseCommaIntsLoose。
func (s strto) CommaInts(value string) []int { return s.ParseCommaIntsLoose(value) }
