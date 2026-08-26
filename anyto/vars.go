package anyto

import (
	"errors"
	"strconv"
)

var (
	// AnyTo 提供任意值转换的结构化入口。
	AnyTo = facade{}

	// ErrSyntax 表示值不符合目标类型要求的语法。
	ErrSyntax = strconv.ErrSyntax

	// ErrType 表示不支持输入值的类型。
	ErrType = errors.New("unsupported type")

	// ErrValOut 表示值超出目标类型的取值范围。
	ErrValOut = errors.New("value out of range")

	// ErrUnsignedInt 表示负数无法转换为无符号整数。
	ErrUnsignedInt = errors.New("cannot convert negative value to unsigned integer")
)

type facade struct{}
