// Package anyto 提供运行时任意值到常见 Go 基础类型的转换工具。
//
// 数值方法支持整数、浮点数、复数实部、布尔值、字符串以及这些值的多级指针，
// 并通过哨兵错误区分语法错误、不支持类型、范围错误和负数转无符号整数。
// Bool 和 String 提供失败时返回零值的宽松转换；StrictBool 和 StrictString
// 用于需要保留失败原因的校验场景。不带 ToStrings 的复数方法接收 []any，并在
// 失败时通过 ElementError 返回元素下标、原始值和底层错误。
package anyto
