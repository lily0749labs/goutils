package time

import (
	"errors"
	"fmt"
	"strings"
	stdtime "time"

	"github.com/dromara/carbon/v2"
)

const defaultTimezone = carbon.Shanghai

var (
	// ErrInvalidTime 表示时间字符串为空或无法解析。
	ErrInvalidTime = errors.New("invalid time value")
	// ErrInvalidLayout 表示 Go 时间布局为空。
	ErrInvalidLayout = errors.New("invalid time layout")
	// ErrInvalidTimezone 表示时区名称无效。
	ErrInvalidTimezone = errors.New("invalid timezone")

	// Time 是使用 Asia/Shanghai 时区和系统时钟的默认 Helper。
	Time = Helper{timezone: defaultTimezone}
	// TimePtr 指向默认 Helper，仅用于兼容旧代码。
	// Deprecated: 请直接使用 Time。
	TimePtr = &Time
)

// Helper 提供带时区和可注入时钟的时间处理方法。
// Helper 可安全复制；WithNow 和 WithTimezone 都返回新副本，不修改原值。
type Helper struct {
	now      func() stdtime.Time
	timezone string
}

// CurrentTimezone 返回 Helper 使用的 IANA 时区名称。
// 零值 Helper 使用默认时区 Asia/Shanghai。
func (t Helper) CurrentTimezone() string {
	if t.timezone == "" {
		return defaultTimezone
	}
	return t.timezone
}

// WithNow 返回使用指定时钟函数的新 Helper，适合测试和任务回放。
// now 为 nil 时，新 Helper 使用系统时钟。
func (t Helper) WithNow(now func() stdtime.Time) Helper {
	t.now = now
	return t
}

// WithTimezone 返回使用指定 IANA 时区的新 Helper，不修改原 Helper 或全局 Time。
// timezone 无法由 time.LoadLocation 加载时返回 ErrInvalidTimezone。
func (t Helper) WithTimezone(timezone string) (Helper, error) {
	if _, err := stdtime.LoadLocation(timezone); err != nil {
		return t, fmt.Errorf("%w %q: %w", ErrInvalidTimezone, timezone, err)
	}
	t.timezone = timezone
	return t, nil
}

func (t Helper) currentTime() stdtime.Time {
	if t.now != nil {
		return t.now()
	}
	return stdtime.Now()
}

// NowCarbon 返回当前时刻的 carbon.Carbon，并转换到 Helper 配置时区。
func (t Helper) NowCarbon() carbon.Carbon {
	return *carbon.CreateFromStdTime(t.currentTime(), t.CurrentTimezone())
}

// NowTime 原样返回时钟函数提供的当前时间，不转换到 Helper 配置时区。
// 使用系统时钟时会保留单调时钟读数，适合计算进程内经过时间。
// 需要配置时区中的时间表示时，请使用 NowInTimezone。
func (t Helper) NowTime() stdtime.Time {
	return t.currentTime()
}

// NowInTimezone 返回转换到 Helper 配置时区的 time.Time。
// 转换会按 Go 标准库规则去除单调时钟读数。
func (t Helper) NowInTimezone() stdtime.Time {
	return t.NowCarbonPtr().StdTime()
}

// NowCarbonPtr 返回当前时刻的 carbon.Carbon 指针，并转换到 Helper 配置时区。
func (t Helper) NowCarbonPtr() *carbon.Carbon {
	now := t.NowCarbon()
	return &now
}

func (t Helper) parseCarbon(value string) *carbon.Carbon {
	switch value {
	case "now":
		return t.NowCarbonPtr()
	case "yesterday":
		return t.NowCarbonPtr().SubDay()
	case "tomorrow":
		return t.NowCarbonPtr().AddDay()
	default:
		return carbon.Parse(value, t.CurrentTimezone())
	}
}

// Parse 使用 Helper 配置时区解析字符串并返回 carbon.Carbon。
// 支持 Carbon 默认格式以及 now、yesterday、tomorrow；空字符串或格式无效时返回 ErrInvalidTime。
func (t Helper) Parse(value string) (carbon.Carbon, error) {
	if value == "" {
		return carbon.Carbon{}, ErrInvalidTime
	}

	parsed := t.parseCarbon(value)
	if parsed.Error != nil {
		return carbon.Carbon{}, fmt.Errorf("%w: %w", ErrInvalidTime, parsed.Error)
	}
	return *parsed, nil
}

// ParseOptionalCarbon 使用 Helper 配置时区解析可选字符串并返回 Carbon 指针。
// 输入为 nil 或格式无效时返回 nil。该方法不区分“未传值”和“解析失败”；
// 需要校验业务输入时，请使用 Parse 并处理其返回错误。
func (t Helper) ParseOptionalCarbon(value *string) *carbon.Carbon {
	if value == nil {
		return nil
	}

	parsed, err := t.Parse(*value)
	if err != nil {
		return nil
	}
	return &parsed
}

// ParseLayout 按 Go layout 和 Helper 配置时区严格解析字符串。
// 空 layout 返回 ErrInvalidLayout；空字符串或格式不匹配返回 ErrInvalidTime。
func (t Helper) ParseLayout(value, layout string) (carbon.Carbon, error) {
	if value == "" {
		return carbon.Carbon{}, ErrInvalidTime
	}
	if layout == "" {
		return carbon.Carbon{}, ErrInvalidLayout
	}
	parsed := carbon.ParseByLayout(value, layout, t.CurrentTimezone())
	if parsed.Error != nil {
		return carbon.Carbon{}, fmt.Errorf("%w: %w", ErrInvalidTime, parsed.Error)
	}
	return *parsed, nil
}

// ParseTime 使用 Helper 配置时区解析字符串并返回 time.Time。
func (t Helper) ParseTime(value string) (stdtime.Time, error) {
	parsed, err := t.Parse(value)
	if err != nil {
		return stdtime.Time{}, err
	}
	return parsed.StdTime(), nil
}

// ParseTimeOrZero 使用 Helper 配置时区解析字符串并返回 time.Time。
// 输入为空或格式无效时返回 time.Time 零值。该方法会忽略解析错误，
// 适合允许无效时间回退为零值的兼容场景；需要保留错误时，请使用 ParseTime。
func (t Helper) ParseTimeOrZero(value string) stdtime.Time {
	parsed, err := t.ParseTime(value)
	if err != nil {
		return stdtime.Time{}
	}
	return parsed
}

// ParseOptionalTime 使用 Helper 配置时区解析可选字符串并返回 time.Time 指针。
// 输入为 nil 或空白字符串时返回 (nil, nil)；格式无效时返回 ErrInvalidTime。
func (t Helper) ParseOptionalTime(value *string) (*stdtime.Time, error) {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil, nil
	}
	parsed, err := t.ParseTime(*value)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// ParseStartOfDay 解析字符串并返回对应日期的零点。
func (t Helper) ParseStartOfDay(value string) (carbon.Carbon, error) {
	parsed, err := t.Parse(value)
	if err != nil {
		return carbon.Carbon{}, err
	}
	return *parsed.StartOfDay(), nil
}

// ParseOptionalStartOfDay 解析可选字符串并返回对应日期零点的 Carbon 指针。
// 输入为 nil 或格式无效时返回 nil；需要保留解析错误时，请使用 ParseStartOfDay。
func (t Helper) ParseOptionalStartOfDay(value *string) *carbon.Carbon {
	if value == nil {
		return nil
	}

	start, err := t.ParseStartOfDay(*value)
	if err != nil {
		return nil
	}
	return &start
}

// ParseEndOfDay 解析字符串并返回对应日期的最后一纳秒。
func (t Helper) ParseEndOfDay(value string) (carbon.Carbon, error) {
	parsed, err := t.Parse(value)
	if err != nil {
		return carbon.Carbon{}, err
	}
	return *parsed.EndOfDay(), nil
}

// ParseOptionalEndOfDay 解析可选字符串并返回对应日期最后一纳秒的 Carbon 指针。
// 输入为 nil 或格式无效时返回 nil；需要保留解析错误时，请使用 ParseEndOfDay。
func (t Helper) ParseOptionalEndOfDay(value *string) *carbon.Carbon {
	if value == nil {
		return nil
	}

	end, err := t.ParseEndOfDay(*value)
	if err != nil {
		return nil
	}
	return &end
}

// ParseDayHalfOpenRange 解析字符串并返回对应日期的半开区间 [start, end)。
// start 是当天零点，end 是次日零点，适合数据库条件 value >= start AND value < end。
func (t Helper) ParseDayHalfOpenRange(value string) (start, end carbon.Carbon, err error) {
	start, err = t.ParseStartOfDay(value)
	if err != nil {
		return carbon.Carbon{}, carbon.Carbon{}, err
	}
	end = *start.AddDay()
	return start, end, nil
}

// FormatNow 使用 Go layout 和 Helper 配置时区格式化当前时间。
func (t Helper) FormatNow(layout string) string {
	return t.NowCarbonPtr().StdTime().Format(layout)
}

// FormatDateTime 将 value 转换到 Helper 配置时区，并格式化为 YYYY-MM-DD HH:mm:ss。
func (t Helper) FormatDateTime(value carbon.Carbon) string {
	return value.ToDateTimeString(t.CurrentTimezone())
}

// FormatOptionalDateTime 将可选 Carbon 转换到 Helper 配置时区，
// 并格式化为 YYYY-MM-DD HH:mm:ss；输入为 nil 时返回 nil。
func (t Helper) FormatOptionalDateTime(value *carbon.Carbon) *string {
	if value == nil {
		return nil
	}

	formatted := t.FormatDateTime(*value)
	return &formatted
}

// FormatDate 将 value 转换到 Helper 配置时区，并格式化为 YYYY-MM-DD。
func (t Helper) FormatDate(value carbon.Carbon) string {
	return value.ToDateString(t.CurrentTimezone())
}

// AddDaysAndFormatDate 为 value 增加指定自然日后，按 Helper 配置时区格式化为 YYYY-MM-DD。
// days 可以为负数。
func (t Helper) AddDaysAndFormatDate(value carbon.Carbon, days int) string {
	return t.FormatDate(*value.AddDays(days))
}

// NowUnixNano 返回当前时刻的纳秒级 Unix 时间戳。
func (t Helper) NowUnixNano() int64 {
	return t.currentTime().UnixNano()
}

// NowUnixMilli 返回当前时刻的毫秒级 Unix 时间戳。
func (t Helper) NowUnixMilli() int64 {
	return t.currentTime().UnixMilli()
}

// NowUnix 返回当前时刻的秒级 Unix 时间戳。
func (t Helper) NowUnix() int64 {
	return t.currentTime().Unix()
}

// NowNanosecond 返回当前秒内的纳秒偏移量，取值范围是 0 到 999999999。
// 它不是 Unix 纳秒时间戳；需要时间戳时请使用 NowUnixNano。
func (t Helper) NowNanosecond() int {
	return t.currentTime().Nanosecond()
}

// NowAddSeconds 在当前时刻增加 seconds 秒，按固定时长计算；负数表示向前推移。
func (t Helper) NowAddSeconds(seconds int) stdtime.Time {
	return t.currentTime().Add(stdtime.Duration(seconds) * stdtime.Second)
}

// NowAddMinutes 在当前时刻增加 minutes 分钟，按固定时长计算；负数表示向前推移。
func (t Helper) NowAddMinutes(minutes int) stdtime.Time {
	return t.currentTime().Add(stdtime.Duration(minutes) * stdtime.Minute)
}

// NowAddHours 在当前时刻增加 hours 小时，按固定时长计算；负数表示向前推移。
func (t Helper) NowAddHours(hours int) stdtime.Time {
	return t.currentTime().Add(stdtime.Duration(hours) * stdtime.Hour)
}

// NowAddDays 在 Helper 配置时区的当前时间上增加 days 个自然日；负数表示向前推移。
// 自然日按日历计算，因此跨越夏令时切换时不一定等于 24 小时。
func (t Helper) NowAddDays(days int) stdtime.Time {
	return t.NowCarbonPtr().AddDays(days).StdTime()
}

// NowAddMonths 在 Helper 配置时区的当前时间上增加 months 个月；负数表示向前推移。
// 本方法保留 time.AddDate 的月末溢出语义；需要截断到目标月末时请使用 NowAddMonthsNoOverflow。
func (t Helper) NowAddMonths(months int) stdtime.Time {
	return t.NowCarbonPtr().AddMonths(months).StdTime()
}

// NowAddMonthsNoOverflow 在 Helper 配置时区的当前时间上增加 months 个月。
// 当原日期超出目标月天数时截断到目标月末；负数表示向前推移。
func (t Helper) NowAddMonthsNoOverflow(months int) stdtime.Time {
	return t.NowCarbonPtr().AddMonthsNoOverflow(months).StdTime()
}

// TodayRange 返回 Helper 配置时区中今天的闭区间边界 [当天零点, 当天最后一纳秒]。
func (t Helper) TodayRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := nowCarbon.EndOfDay()
	return *start, *end
}

// TodayHalfOpenRange 返回 Helper 配置时区中今天的半开区间 [当天零点, 次日零点)。
func (t Helper) TodayHalfOpenRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := start.AddDay()
	return *start, *end
}

// YesterdayRange 返回 Helper 配置时区中昨天的闭区间边界 [昨天零点, 昨天最后一纳秒]。
func (t Helper) YesterdayRange() (carbon.Carbon, carbon.Carbon) {
	yesterday := t.NowCarbonPtr().SubDay()
	start := yesterday.StartOfDay()
	end := yesterday.EndOfDay()
	return *start, *end
}

// YesterdayHalfOpenRange 返回 Helper 配置时区中昨天的半开区间 [昨天零点, 今天零点)。
func (t Helper) YesterdayHalfOpenRange() (carbon.Carbon, carbon.Carbon) {
	todayStart := t.NowCarbonPtr().StartOfDay()
	start := todayStart.SubDay()
	return *start, *todayStart
}

// PreviousMonthRange 返回 Helper 配置时区中上个月的闭区间边界 [月初, 月末最后一纳秒]。
func (t Helper) PreviousMonthRange() (carbon.Carbon, carbon.Carbon) {
	currentMonthStart := t.NowCarbonPtr().StartOfMonth()
	start := currentMonthStart.SubMonth()
	end := start.EndOfMonth()
	return *start, *end
}

// PreviousMonthHalfOpenRange 返回 Helper 配置时区中上个月的半开区间 [上月月初, 本月月初)。
func (t Helper) PreviousMonthHalfOpenRange() (carbon.Carbon, carbon.Carbon) {
	end := t.NowCarbonPtr().StartOfMonth()
	start := end.SubMonth()
	return *start, *end
}
