package anyto

import (
	"fmt"
)

// ElementError 描述严格切片中首个转换失败的元素。
type ElementError struct {
	// Index 是转换失败元素在输入切片中的下标。
	Index int
	// Value 是转换失败的原始元素。
	Value any
	// Err 是该元素转换失败的底层原因。
	Err error
}

// Error 返回包含失败元素下标、原始值和底层原因的错误文本。
func (e *ElementError) Error() string {
	return fmt.Sprintf("convert element at index %d with value %v: %v", e.Index, e.Value, e.Err)
}

// Unwrap 返回底层转换错误，使调用方可以继续使用 errors.Is 或 errors.As。
func (e *ElementError) Unwrap() error { return e.Err }

// convertSlice 按顺序严格转换 []any；任一元素失败时放弃部分结果并返回元素上下文。
func convertSlice[T any](values []any, convert func(any) (T, error)) ([]T, error) {
	if values == nil {
		return nil, nil
	}
	result := make([]T, len(values))
	for index, value := range values {
		if _, err := strictValue(value); err != nil {
			return nil, &ElementError{Index: index, Value: value, Err: err}
		}
		converted, err := convert(value)
		if err != nil {
			return nil, &ElementError{Index: index, Value: value, Err: err}
		}
		result[index] = converted
	}
	return result, nil
}

// Ints 将 []any 中的每个元素严格转换为 int。
func (f anyto) Ints(values []any) ([]int, error) { return convertSlice(values, f.Int) }

// Int8s 将 []any 中的每个元素严格转换为 int8。
func (f anyto) Int8s(values []any) ([]int8, error) { return convertSlice(values, f.Int8) }

// Int16s 将 []any 中的每个元素严格转换为 int16。
func (f anyto) Int16s(values []any) ([]int16, error) { return convertSlice(values, f.Int16) }

// Int32s 将 []any 中的每个元素严格转换为 int32。
func (f anyto) Int32s(values []any) ([]int32, error) { return convertSlice(values, f.Int32) }

// Int64s 将 []any 中的每个元素严格转换为 int64。
func (f anyto) Int64s(values []any) ([]int64, error) { return convertSlice(values, f.Int64) }

// Uints 将 []any 中的每个元素严格转换为 uint。
func (f anyto) Uints(values []any) ([]uint, error) { return convertSlice(values, f.Uint) }

// Uint8s 将 []any 中的每个元素严格转换为 uint8。
func (f anyto) Uint8s(values []any) ([]uint8, error) { return convertSlice(values, f.Uint8) }

// Uint16s 将 []any 中的每个元素严格转换为 uint16。
func (f anyto) Uint16s(values []any) ([]uint16, error) { return convertSlice(values, f.Uint16) }

// Uint32s 将 []any 中的每个元素严格转换为 uint32。
func (f anyto) Uint32s(values []any) ([]uint32, error) { return convertSlice(values, f.Uint32) }

// Uint64s 将 []any 中的每个元素严格转换为 uint64。
func (f anyto) Uint64s(values []any) ([]uint64, error) { return convertSlice(values, f.Uint64) }

// Float32s 将 []any 中的每个元素严格转换为有限的 float32。
func (f anyto) Float32s(values []any) ([]float32, error) {
	return convertSlice(values, f.Float32)
}

// Float64s 将 []any 中的每个元素严格转换为有限的 float64。
func (f anyto) Float64s(values []any) ([]float64, error) {
	return convertSlice(values, f.Float64)
}

// StrictBools 将 []any 中的每个元素严格转换为 bool。
func (f anyto) StrictBools(values []any) ([]bool, error) {
	return convertSlice(values, f.StrictBool)
}

// StrictStrings 将 []any 中的每个元素严格转换为 string。
func (f anyto) StrictStrings(values []any) ([]string, error) {
	return convertSlice(values, f.StrictString)
}
