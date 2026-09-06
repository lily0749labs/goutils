package time

import (
	stdtime "time"

	"github.com/dromara/carbon/v2"
)

// FromUnix 将秒级 Unix 时间戳转换为 Helper 配置时区中的 carbon.Carbon。
func (t Helper) FromUnix(timestamp int64) carbon.Carbon {
	return *carbon.CreateFromTimestamp(timestamp, t.CurrentTimezone())
}

// FromUnixMilli 将毫秒级 Unix 时间戳转换为 Helper 配置时区中的 carbon.Carbon。
func (t Helper) FromUnixMilli(timestamp int64) carbon.Carbon {
	return *carbon.CreateFromTimestampMilli(timestamp, t.CurrentTimezone())
}

// FromUnixNano 将纳秒级 Unix 时间戳转换为 Helper 配置时区中的 carbon.Carbon。
func (t Helper) FromUnixNano(timestamp int64) carbon.Carbon {
	return *carbon.CreateFromTimestampNano(timestamp, t.CurrentTimezone())
}

// ToUnix 返回 value 的秒级 Unix 时间戳；结果与 value 的显示时区无关。
func (Helper) ToUnix(value stdtime.Time) int64 { return value.Unix() }

// ToUnixMilli 返回 value 的毫秒级 Unix 时间戳；结果与 value 的显示时区无关。
func (Helper) ToUnixMilli(value stdtime.Time) int64 { return value.UnixMilli() }

// ToUnixNano 返回 value 的纳秒级 Unix 时间戳；结果与 value 的显示时区无关。
func (Helper) ToUnixNano(value stdtime.Time) int64 { return value.UnixNano() }
