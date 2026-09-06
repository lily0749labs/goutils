package strto

import (
	"fmt"
	stdtime "time"
)

// ElementError 描述严格切片转换失败的元素位置、原始值和底层错误。
type ElementError struct {
	Index int
	Value string
	Err   error
}

// Error 返回包含失败下标和值的错误说明。
func (e *ElementError) Error() string {
	return fmt.Sprintf("convert element at index %d with value %q: %v", e.Index, e.Value, e.Err)
}

// Unwrap 返回底层转换错误，支持 errors.Is 和 errors.As。
func (e *ElementError) Unwrap() error { return e.Err }

func convertSliceE[T any](values []string, convert func(string) (T, error)) ([]T, error) {
	if values == nil {
		return nil, nil
	}
	result := make([]T, len(values))
	for index, value := range values {
		converted, err := convert(value)
		if err != nil {
			return nil, &ElementError{Index: index, Value: value, Err: err}
		}
		result[index] = converted
	}
	return result, nil
}

// BoolsE 严格转换 bool 切片，任一元素失败时返回 ElementError。
func (s strto) BoolsE(values []string) ([]bool, error) {
	return convertSliceE(values, s.BoolE)
}

// IntsE 严格转换 int 切片，任一元素失败时返回 ElementError。
func (s strto) IntsE(values []string) ([]int, error) {
	return convertSliceE(values, s.IntE)
}

// Int8sE 严格转换 int8 切片，任一元素失败时返回 ElementError。
func (s strto) Int8sE(values []string) ([]int8, error) {
	return convertSliceE(values, s.Int8E)
}

// Int16sE 严格转换 int16 切片，任一元素失败时返回 ElementError。
func (s strto) Int16sE(values []string) ([]int16, error) {
	return convertSliceE(values, s.Int16E)
}

// Int32sE 严格转换 int32 切片，任一元素失败时返回 ElementError。
func (s strto) Int32sE(values []string) ([]int32, error) {
	return convertSliceE(values, s.Int32E)
}

// Int64sE 严格转换 int64 切片，任一元素失败时返回 ElementError。
func (s strto) Int64sE(values []string) ([]int64, error) {
	return convertSliceE(values, s.Int64E)
}

// UintsE 严格转换 uint 切片，任一元素失败时返回 ElementError。
func (s strto) UintsE(values []string) ([]uint, error) {
	return convertSliceE(values, s.UintE)
}

// Uint8sE 严格转换 uint8 切片，任一元素失败时返回 ElementError。
func (s strto) Uint8sE(values []string) ([]uint8, error) {
	return convertSliceE(values, s.Uint8E)
}

// Uint16sE 严格转换 uint16 切片，任一元素失败时返回 ElementError。
func (s strto) Uint16sE(values []string) ([]uint16, error) {
	return convertSliceE(values, s.Uint16E)
}

// Uint32sE 严格转换 uint32 切片，任一元素失败时返回 ElementError。
func (s strto) Uint32sE(values []string) ([]uint32, error) {
	return convertSliceE(values, s.Uint32E)
}

// Uint64sE 严格转换 uint64 切片，任一元素失败时返回 ElementError。
func (s strto) Uint64sE(values []string) ([]uint64, error) {
	return convertSliceE(values, s.Uint64E)
}

// Float32sE 严格转换 float32 切片，任一元素失败时返回 ElementError。
func (s strto) Float32sE(values []string) ([]float32, error) {
	return convertSliceE(values, s.Float32E)
}

// Float64sE 严格转换 float64 切片，任一元素失败时返回 ElementError。
func (s strto) Float64sE(values []string) ([]float64, error) {
	return convertSliceE(values, s.Float64E)
}

// DurationsE 严格转换 time.Duration 切片，任一元素失败时返回 ElementError。
func (s strto) DurationsE(values []string) ([]stdtime.Duration, error) {
	return convertSliceE(values, s.DurationE)
}
