package stringutil

import (
	"strings"
	"unicode"
)

// IsBlank 判断字符串是否为空或只包含 Unicode 空白字符。
func (stringutil) IsBlank(value string) bool { return strings.TrimSpace(value) == "" }

// NormalizeSpace 将连续 Unicode 空白折叠为一个半角空格，并移除首尾空白。
func (stringutil) NormalizeSpace(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// Reverse 按 Unicode 字符反转字符串，不会拆开多字节字符。
func (stringutil) Reverse(value string) string {
	runes := []rune(value)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

// Truncate 最多保留 maxRunes 个 Unicode 字符。
func (stringutil) Truncate(value string, maxRunes int) string {
	return truncateWithSuffix(value, maxRunes, "")
}

// TruncateWithSuffix 将字符串限制在 maxRunes 个 Unicode 字符内，并在截断时追加 suffix。
// suffix 计入最大长度；suffix 本身过长时也会被截断。
func (stringutil) TruncateWithSuffix(value string, maxRunes int, suffix string) string {
	return truncateWithSuffix(value, maxRunes, suffix)
}

func truncateWithSuffix(value string, maxRunes int, suffix string) string {
	if maxRunes <= 0 {
		return ""
	}
	valueRunes := []rune(value)
	if len(valueRunes) <= maxRunes {
		return value
	}
	suffixRunes := []rune(suffix)
	if len(suffixRunes) >= maxRunes {
		return string(suffixRunes[:maxRunes])
	}
	return string(valueRunes[:maxRunes-len(suffixRunes)]) + suffix
}

// Mask 保留左右指定数量的 Unicode 字符，中间字符使用星号遮盖。
func (s stringutil) Mask(value string, visibleLeft, visibleRight int) string {
	return s.MaskWith(value, visibleLeft, visibleRight, '*')
}

// MaskWith 使用 mask 遮盖中间字符；mask 为零值时使用星号。
func (stringutil) MaskWith(value string, visibleLeft, visibleRight int, mask rune) string {
	if visibleLeft < 0 {
		visibleLeft = 0
	}
	if visibleRight < 0 {
		visibleRight = 0
	}
	if mask == 0 {
		mask = '*'
	}
	runes := []rune(value)
	if visibleLeft+visibleRight >= len(runes) {
		return value
	}
	for index := visibleLeft; index < len(runes)-visibleRight; index++ {
		runes[index] = mask
	}
	return string(runes)
}

// Snake 将标识符转换为 snake_case。
func (stringutil) Snake(value string) string { return strings.Join(splitWords(value), "_") }

// Kebab 将标识符转换为 kebab-case。
func (stringutil) Kebab(value string) string { return strings.Join(splitWords(value), "-") }

// Camel 将标识符转换为 lowerCamelCase。
func (stringutil) Camel(value string) string {
	words := splitWords(value)
	if len(words) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.WriteString(words[0])
	for _, word := range words[1:] {
		builder.WriteString(upperFirst(word))
	}
	return builder.String()
}

// Pascal 将标识符转换为 PascalCase。
func (stringutil) Pascal(value string) string {
	words := splitWords(value)
	var builder strings.Builder
	for _, word := range words {
		builder.WriteString(upperFirst(word))
	}
	return builder.String()
}

func upperFirst(value string) string {
	runes := []rune(value)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func splitWords(value string) []string {
	runes := []rune(strings.TrimSpace(value))
	words := make([]string, 0)
	start := -1
	flush := func(end int) {
		if start >= 0 && start < end {
			words = append(words, strings.ToLower(string(runes[start:end])))
		}
		start = -1
	}

	for index, current := range runes {
		if !unicode.IsLetter(current) && !unicode.IsDigit(current) {
			flush(index)
			continue
		}
		if start < 0 {
			start = index
			continue
		}

		previous := runes[index-1]
		isBoundary := unicode.IsUpper(current) && (unicode.IsLower(previous) || unicode.IsDigit(previous))
		if !isBoundary && unicode.IsUpper(current) && unicode.IsUpper(previous) && index+1 < len(runes) {
			isBoundary = unicode.IsLower(runes[index+1])
		}
		if !isBoundary {
			isBoundary = unicode.IsDigit(current) != unicode.IsDigit(previous)
		}
		if isBoundary {
			flush(index)
			start = index
		}
	}
	flush(len(runes))
	return words
}
