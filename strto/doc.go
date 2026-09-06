// Package strto 提供字符串到常见 Go 基础类型的解析与零值回退方法。
//
// Parse 前缀的方法返回可分类的错误，适合需要校验输入的场景；OrZero 和 OrFalse
// 后缀的方法忽略解析错误并返回目标类型的零值。复数方法处理 []string；任一元素
// 解析失败时，通过 ElementError 报告元素下标、原始值和底层错误。
//
// 推荐通过根包门面 goutils.StrTo 调用。完整的功能分类和边界行为参见包内 README.md。
package strto
