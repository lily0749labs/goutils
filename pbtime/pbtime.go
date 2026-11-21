// Package pbtime 提供 protobuf 时间戳（timestamppb.Timestamp）的格式化。
//
// 单独成包，是为了只有用到 protobuf 的项目才会链接 google.golang.org/protobuf；
// 其他子包和根包的门面都不依赖它。
package pbtime

import (
	"time"

	"github.com/dromara/carbon/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Format 把 protobuf 时间戳转换到 loc 时区，并格式化为 YYYY-MM-DD HH:mm:ss。
// value 为 nil（字段没有设置）时返回空字符串，而不是 1970-01-01；loc 为 nil 时按 UTC。
func Format(value *timestamppb.Timestamp, loc *time.Location) string {
	// 没有设置的时间戳没有时间可显示。
	if value == nil {
		return ""
	}
	// 没有指定时区时按 UTC，避免 In(nil) 触发 panic。
	if loc == nil {
		loc = time.UTC
	}
	// 先换算到指定时区，再按固定格式输出。
	return value.AsTime().In(loc).Format(time.DateTime)
}

// FromCarbon 把 carbon 时间（常见于数据库生成模型里的可空时间）转成 protobuf 时间戳。
// t 为 nil 或零时间（0001-01-01）时返回 nil，表示这个时间字段没有设置。
func FromCarbon(t *carbon.Carbon) *timestamppb.Timestamp {
	// 没有时间或零时间：不设置，调用方显示成空。
	if t == nil || t.IsZero() {
		return nil
	}
	return timestamppb.New(t.StdTime())
}
