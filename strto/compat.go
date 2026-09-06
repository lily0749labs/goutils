package strto

import stdtime "time"

// IntE 是旧版兼容入口。
// Deprecated: 请使用 ParseInt。
func (s strto) IntE(value string) (int, error) { return s.ParseInt(value) }

// Int8E 是旧版兼容入口。
// Deprecated: 请使用 ParseInt8。
func (s strto) Int8E(value string) (int8, error) { return s.ParseInt8(value) }

// Int16E 是旧版兼容入口。
// Deprecated: 请使用 ParseInt16。
func (s strto) Int16E(value string) (int16, error) { return s.ParseInt16(value) }

// Int32E 是旧版兼容入口。
// Deprecated: 请使用 ParseInt32。
func (s strto) Int32E(value string) (int32, error) { return s.ParseInt32(value) }

// Int64E 是旧版兼容入口。
// Deprecated: 请使用 ParseInt64。
func (s strto) Int64E(value string) (int64, error) { return s.ParseInt64(value) }

// UintE 是旧版兼容入口。
// Deprecated: 请使用 ParseUint。
func (s strto) UintE(value string) (uint, error) { return s.ParseUint(value) }

// Uint8E 是旧版兼容入口。
// Deprecated: 请使用 ParseUint8。
func (s strto) Uint8E(value string) (uint8, error) { return s.ParseUint8(value) }

// Uint16E 是旧版兼容入口。
// Deprecated: 请使用 ParseUint16。
func (s strto) Uint16E(value string) (uint16, error) { return s.ParseUint16(value) }

// Uint32E 是旧版兼容入口。
// Deprecated: 请使用 ParseUint32。
func (s strto) Uint32E(value string) (uint32, error) { return s.ParseUint32(value) }

// Uint64E 是旧版兼容入口。
// Deprecated: 请使用 ParseUint64。
func (s strto) Uint64E(value string) (uint64, error) { return s.ParseUint64(value) }

// BoolE 是旧版兼容入口。
// Deprecated: 请使用 ParseBool。
func (s strto) BoolE(value string) (bool, error) { return s.ParseBool(value) }

// Float32E 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat32。
func (s strto) Float32E(value string) (float32, error) { return s.ParseFiniteFloat32(value) }

// Float64E 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat64。
func (s strto) Float64E(value string) (float64, error) { return s.ParseFiniteFloat64(value) }

// DurationE 是旧版兼容入口。
// Deprecated: 请使用 ParseDuration。
func (s strto) DurationE(value string) (stdtime.Duration, error) { return s.ParseDuration(value) }

// BoolsE 是旧版兼容入口。
// Deprecated: 请使用 ParseBools。
func (s strto) BoolsE(values []string) ([]bool, error) { return s.ParseBools(values) }

// IntsE 是旧版兼容入口。
// Deprecated: 请使用 ParseInts。
func (s strto) IntsE(values []string) ([]int, error) { return s.ParseInts(values) }

// Int8sE 是旧版兼容入口。
// Deprecated: 请使用 ParseInt8s。
func (s strto) Int8sE(values []string) ([]int8, error) { return s.ParseInt8s(values) }

// Int16sE 是旧版兼容入口。
// Deprecated: 请使用 ParseInt16s。
func (s strto) Int16sE(values []string) ([]int16, error) { return s.ParseInt16s(values) }

// Int32sE 是旧版兼容入口。
// Deprecated: 请使用 ParseInt32s。
func (s strto) Int32sE(values []string) ([]int32, error) { return s.ParseInt32s(values) }

// Int64sE 是旧版兼容入口。
// Deprecated: 请使用 ParseInt64s。
func (s strto) Int64sE(values []string) ([]int64, error) { return s.ParseInt64s(values) }

// UintsE 是旧版兼容入口。
// Deprecated: 请使用 ParseUints。
func (s strto) UintsE(values []string) ([]uint, error) { return s.ParseUints(values) }

// Uint8sE 是旧版兼容入口。
// Deprecated: 请使用 ParseUint8s。
func (s strto) Uint8sE(values []string) ([]uint8, error) { return s.ParseUint8s(values) }

// Uint16sE 是旧版兼容入口。
// Deprecated: 请使用 ParseUint16s。
func (s strto) Uint16sE(values []string) ([]uint16, error) { return s.ParseUint16s(values) }

// Uint32sE 是旧版兼容入口。
// Deprecated: 请使用 ParseUint32s。
func (s strto) Uint32sE(values []string) ([]uint32, error) { return s.ParseUint32s(values) }

// Uint64sE 是旧版兼容入口。
// Deprecated: 请使用 ParseUint64s。
func (s strto) Uint64sE(values []string) ([]uint64, error) { return s.ParseUint64s(values) }

// Float32sE 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat32s。
func (s strto) Float32sE(values []string) ([]float32, error) {
	return s.ParseFiniteFloat32s(values)
}

// Float64sE 是旧版兼容入口。
// Deprecated: 请使用 ParseFiniteFloat64s。
func (s strto) Float64sE(values []string) ([]float64, error) {
	return s.ParseFiniteFloat64s(values)
}

// DurationsE 是旧版兼容入口。
// Deprecated: 请使用 ParseDurations。
func (s strto) DurationsE(values []string) ([]stdtime.Duration, error) {
	return s.ParseDurations(values)
}

// StrictCommaInts 是旧版兼容入口。
// Deprecated: 请使用 ParseCommaInts。
func (s strto) StrictCommaInts(value string) ([]int, error) { return s.ParseCommaInts(value) }

// StrToInt 是旧版兼容入口。
// Deprecated: 请使用 StrTo.IntOrZero。
func StrToInt(value string) int { return StrTo.IntOrZero(value) }

// StrToInt8 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int8OrZero。
func StrToInt8(value string) int8 { return StrTo.Int8OrZero(value) }

// StrToInt16 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int16OrZero。
func StrToInt16(value string) int16 { return StrTo.Int16OrZero(value) }

// StrToInt32 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int32OrZero。
func StrToInt32(value string) int32 { return StrTo.Int32OrZero(value) }

// StrToInt64 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int64OrZero。
func StrToInt64(value string) int64 { return StrTo.Int64OrZero(value) }

// StrToUint 是旧版兼容入口。
// Deprecated: 请使用 StrTo.UintOrZero。
func StrToUint(value string) uint { return StrTo.UintOrZero(value) }

// StrToUint8 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint8OrZero。
func StrToUint8(value string) uint8 { return StrTo.Uint8OrZero(value) }

// StrToUint16 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint16OrZero。
func StrToUint16(value string) uint16 { return StrTo.Uint16OrZero(value) }

// StrToUint32 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint32OrZero。
func StrToUint32(value string) uint32 { return StrTo.Uint32OrZero(value) }

// StrToUint64 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint64OrZero。
func StrToUint64(value string) uint64 { return StrTo.Uint64OrZero(value) }

// StrToBool 是旧版兼容入口。
// Deprecated: 请使用 StrTo.BoolOrFalse。
func StrToBool(value string) bool { return StrTo.BoolOrFalse(value) }

// StrToFloat32 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Float32OrZero。
func StrToFloat32(value string) float32 { return StrTo.Float32OrZero(value) }

// StrToFloat64 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Float64OrZero。
func StrToFloat64(value string) float64 { return StrTo.Float64OrZero(value) }

// StrToDuration 是旧版兼容入口。
// Deprecated: 请使用 StrTo.DurationOrZero。
func StrToDuration(value string) stdtime.Duration { return StrTo.DurationOrZero(value) }

// StrToBytes 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Bytes。
func StrToBytes(value string) []byte { return StrTo.Bytes(value) }

// ArrStrToBool 是旧版兼容入口。
// Deprecated: 请使用 StrTo.BoolsOrFalse。
func ArrStrToBool(values []string) []bool { return StrTo.BoolsOrFalse(values) }

// ArrStrToInt 是旧版兼容入口。
// Deprecated: 请使用 StrTo.IntsOrZero。
func ArrStrToInt(values []string) []int { return StrTo.IntsOrZero(values) }

// ArrStrToInt8 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int8sOrZero。
func ArrStrToInt8(values []string) []int8 { return StrTo.Int8sOrZero(values) }

// ArrStrToInt16 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int16sOrZero。
func ArrStrToInt16(values []string) []int16 { return StrTo.Int16sOrZero(values) }

// ArrStrToInt32 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int32sOrZero。
func ArrStrToInt32(values []string) []int32 { return StrTo.Int32sOrZero(values) }

// ArrStrToInt64 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Int64sOrZero。
func ArrStrToInt64(values []string) []int64 { return StrTo.Int64sOrZero(values) }

// ArrStrToUint 是旧版兼容入口。
// Deprecated: 请使用 StrTo.UintsOrZero。
func ArrStrToUint(values []string) []uint { return StrTo.UintsOrZero(values) }

// ArrStrToUint8 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint8sOrZero。
func ArrStrToUint8(values []string) []uint8 { return StrTo.Uint8sOrZero(values) }

// ArrStrToUint16 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint16sOrZero。
func ArrStrToUint16(values []string) []uint16 { return StrTo.Uint16sOrZero(values) }

// ArrStrToUint32 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint32sOrZero。
func ArrStrToUint32(values []string) []uint32 { return StrTo.Uint32sOrZero(values) }

// ArrStrToUint64 是旧版兼容入口。
// Deprecated: 请使用 StrTo.Uint64sOrZero。
func ArrStrToUint64(values []string) []uint64 { return StrTo.Uint64sOrZero(values) }

// CommaInts 是旧版兼容入口。
// Deprecated: 请使用 StrTo.ParseCommaIntsLoose。
func CommaInts(value string) []int { return StrTo.ParseCommaIntsLoose(value) }

// ParseCommaInts 是旧版兼容入口。
// Deprecated: 请使用 StrTo.ParseCommaInts。
func ParseCommaInts(value string) ([]int, error) { return StrTo.ParseCommaInts(value) }
