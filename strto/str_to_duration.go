package strto

import (
	"fmt"
	stdtime "time"
)

// DurationE 按 Go 时长语法严格解析字符串，例如 1h30m、250ms。
func (strto) DurationE(value string) (stdtime.Duration, error) {
	parsed, err := stdtime.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidDuration, err)
	}
	return parsed, nil
}

// Duration 解析 Go 时长，失败时返回零值。
func (s strto) Duration(value string) stdtime.Duration {
	parsed, _ := s.DurationE(value)
	return parsed
}

// Deprecated: 使用 StrTo.Duration。
func StrToDuration(value string) stdtime.Duration { return StrTo.Duration(value) }
