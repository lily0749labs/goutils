package floatutil

import (
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// Float 提供浮点数辅助函数的结构化入口。
var Float = floatutil{}

type floatutil struct{}

// Truncate 将 value 截断到 precision 位小数，不进行四舍五入。
// precision 小于等于零时截断到整数；NaN 和无穷值保持不变。
func (floatutil) Truncate(value float64, precision int) float64 {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return value
	}
	if precision <= 0 {
		return math.Trunc(value)
	}

	text := strconv.FormatFloat(value, 'f', -1, 64)
	dot := strings.IndexByte(text, '.')
	if dot < 0 || len(text)-dot-1 <= precision {
		return value
	}

	result, err := strconv.ParseFloat(text[:dot+1+precision], 64)
	if err != nil {
		return value
	}
	return result
}

// TruncateString 将 value 截断到 precision 位小数并返回字符串，不补零、不四舍五入。
// precision 小于零时按零处理。
func (floatutil) TruncateString(value float64, precision int) string {
	if precision < 0 {
		precision = 0
	}
	text := strconv.FormatFloat(value, 'f', -1, 64)
	dot := strings.IndexByte(text, '.')
	if dot < 0 || len(text)-dot-1 <= precision {
		return text
	}
	return text[:dot+1+precision]
}

// Add 使用十进制定点运算完成两个 float64 的加法，再转换回 float64。
func (floatutil) Add(first, second float64) float64 {
	result, _ := decimal.NewFromFloat(first).Add(decimal.NewFromFloat(second)).Float64()
	return result
}

// Mul 使用十进制定点运算完成两个 float64 的乘法，再转换回 float64。
func (floatutil) Mul(first, second float64) float64 {
	result, _ := decimal.NewFromFloat(first).Mul(decimal.NewFromFloat(second)).Float64()
	return result
}

// Truncate 是 Float.Truncate 的包级便捷入口。
func Truncate(value float64, precision int) float64 {
	return Float.Truncate(value, precision)
}

// TruncateString 是 Float.TruncateString 的包级便捷入口。
func TruncateString(value float64, precision int) string {
	return Float.TruncateString(value, precision)
}

// Add 是 Float.Add 的包级便捷入口。
func Add(first, second float64) float64 {
	return Float.Add(first, second)
}

// Mul 是 Float.Mul 的包级便捷入口。
func Mul(first, second float64) float64 {
	return Float.Mul(first, second)
}
