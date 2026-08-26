package money

import "fmt"

// Money 提供金额格式化的结构化入口。
var Money = money{}

type money struct{}

// FormatCents 将以分为单位的整数格式化为两位小数字符串。
// 该方法不添加货币符号或千位分隔符。
func (money) FormatCents(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
	}

	major := cents / 100
	minor := cents % 100
	if major < 0 {
		major = -major
	}
	if minor < 0 {
		minor = -minor
	}

	return fmt.Sprintf("%s%d.%02d", sign, major, minor)
}

// FormatCents 是 Money.FormatCents 的包级便捷入口。
func FormatCents(cents int64) string { return Money.FormatCents(cents) }
