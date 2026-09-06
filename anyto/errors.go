package anyto

import (
	"errors"
	"strconv"
)

var (
	// ErrSyntax 表示字符串不符合目标类型的解析语法。
	ErrSyntax = strconv.ErrSyntax

	// ErrUnsupportedType 表示输入类型不受支持；严格转换收到 nil 时也返回此错误。
	ErrUnsupportedType = errors.New("unsupported type")

	// ErrOutOfRange 表示输入超出目标类型的取值范围，或输入为 NaN、正负无穷大。
	ErrOutOfRange = errors.New("value out of range")

	// ErrNegativeToUnsigned 表示负数无法转换为无符号整数。
	ErrNegativeToUnsigned = errors.New("cannot convert negative value to unsigned integer")

	// ErrType 是 ErrUnsupportedType 的旧名称。
	// Deprecated: 请使用 ErrUnsupportedType。
	ErrType = ErrUnsupportedType

	// ErrValOut 是 ErrOutOfRange 的旧名称。
	// Deprecated: 请使用 ErrOutOfRange。
	ErrValOut = ErrOutOfRange

	// ErrUnsignedInt 是 ErrNegativeToUnsigned 的旧名称。
	// Deprecated: 请使用 ErrNegativeToUnsigned。
	ErrUnsignedInt = ErrNegativeToUnsigned
)

// classifyNumberError 将 strconv 的错误归一为 anyto 对外暴露的错误类型。
// 数值格式正确但超出目标范围时返回 ErrOutOfRange，其余情况返回 ErrSyntax。
func classifyNumberError(err error) error {
	if errors.Is(err, strconv.ErrRange) {
		return ErrOutOfRange
	}
	return ErrSyntax
}
