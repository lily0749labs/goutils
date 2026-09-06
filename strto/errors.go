package strto

import (
	"fmt"
	"strconv"
)

var (
	// ErrSyntax 表示输入不符合目标类型语法，等同于 strconv.ErrSyntax。
	ErrSyntax = strconv.ErrSyntax

	// ErrOutOfRange 表示数值超出目标类型范围，或浮点数不是有限值；等同于 strconv.ErrRange。
	ErrOutOfRange = strconv.ErrRange

	// ErrInvalidBool 表示输入不是支持的布尔值。
	// 该错误同时包装 ErrSyntax，可分别按具体错误或错误分类判断。
	ErrInvalidBool = fmt.Errorf("value is not a supported boolean: %w", ErrSyntax)

	// ErrNonFiniteFloat 表示浮点数是 NaN 或无穷大。
	// 该错误同时包装 ErrOutOfRange。
	ErrNonFiniteFloat = fmt.Errorf("float value must be finite: %w", ErrOutOfRange)

	// ErrInvalidDuration 表示输入不是有效的 Go 时长。
	// 该错误同时包装 ErrSyntax。
	ErrInvalidDuration = fmt.Errorf("invalid duration: %w", ErrSyntax)
)

// ElementError 描述切片或列表中解析失败的元素位置、原始值和底层错误。
type ElementError struct {
	// Index 是失败元素在输入中的下标，从 0 开始。
	Index int
	// Value 是失败元素的原始字符串。
	Value string
	// Err 是该元素解析失败的底层原因。
	Err error
}

// Error 返回包含失败下标和值的错误说明。
func (e *ElementError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("convert element at index %d with value %q: %v", e.Index, e.Value, e.Err)
}

// Unwrap 返回底层解析错误，支持 errors.Is 和 errors.As。
func (e *ElementError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}
