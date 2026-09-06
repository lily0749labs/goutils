package strto

import (
	"fmt"
	stdtime "time"
)

// ParseDuration 按 Go time.Duration 语法解析字符串，例如 1h30m、250ms；裸数字不表示秒。
func (strto) ParseDuration(value string) (stdtime.Duration, error) {
	parsed, err := stdtime.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidDuration, err)
	}
	return parsed, nil
}

// DurationOrZero 按 Go time.Duration 语法解析字符串；解析失败时返回 0。
func (s strto) DurationOrZero(value string) stdtime.Duration {
	parsed, _ := s.ParseDuration(value)
	return parsed
}
