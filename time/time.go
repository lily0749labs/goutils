package time

import (
	"errors"
	"fmt"
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

	// Time 提供常用时间辅助方法。
	Time = Helper{timezone: defaultTimezone}
	// TimePtr 指向 Time，为兼容旧代码而保留。
	TimePtr = &Time
)

// Helper 是可配置、可注入时钟的时间工具门面。
// Helper 值可安全复制；WithNow 和 WithTimezone 均返回独立副本。
type Helper struct {
	now      func() stdtime.Time
	timezone string
}

// CurrentTimezone 返回当前配置的 IANA 时区名称。
func (t Helper) CurrentTimezone() string {
	if t.timezone == "" {
		return defaultTimezone
	}
	return t.timezone
}

// CurrentLayout 返回当前配置的时区。
// Deprecated: 请使用 CurrentTimezone。
func (t Helper) CurrentLayout() string {
	return t.CurrentTimezone()
}

// WithNow 返回使用指定当前时间函数的独立时间门面，便于测试和任务回放。
// now 为 nil 时恢复使用系统时间。
func (t Helper) WithNow(now func() stdtime.Time) Helper {
	t.now = now
	return t
}

// WithTimezone 返回使用指定时区的独立时间门面，不修改全局 Time。
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

// GetCurrLayout 返回当前配置的 Carbon 时区。
// Deprecated: 请使用 CurrentTimezone。
func (t Helper) GetCurrLayout() string {
	return t.CurrentTimezone()
}

// NowCarbon 返回 Carbon 类型的当前时间。
func (t Helper) NowCarbon() carbon.Carbon {
	return *carbon.CreateFromStdTime(t.currentTime(), t.CurrentTimezone())
}

// NowTime 返回时钟提供的当前时间。
// 系统时钟的返回值保留 Go 单调时钟，适合计算进程内持续时间。
// 如需将时间转换到配置时区，请使用 NowInTimezone。
func (t Helper) NowTime() stdtime.Time {
	return t.currentTime()
}

// NowInTimezone 返回转换到当前配置时区的标准库时间。
// 时区转换会按 Go 标准库规则去除单调时钟读数。
func (t Helper) NowInTimezone() stdtime.Time {
	return t.NowCarbonPtr().StdTime()
}

// NowCarbonPtr 返回 Carbon 类型的当前时间指针。
func (t Helper) NowCarbonPtr() *carbon.Carbon {
	now := t.NowCarbon()
	return &now
}

// StrToCarbon 使用当前配置的 Carbon 时区解析 strTime。
func (t Helper) StrToCarbon(strTime string) carbon.Carbon {
	return *t.parse(strTime)
}

func (t Helper) parse(strTime string) *carbon.Carbon {
	switch strTime {
	case "now":
		return t.NowCarbonPtr()
	case "yesterday":
		return t.NowCarbonPtr().SubDay()
	case "tomorrow":
		return t.NowCarbonPtr().AddDay()
	default:
		return carbon.Parse(strTime, t.CurrentTimezone())
	}
}

// ParseE 使用当前时区解析时间，失败时返回明确错误。
func (t Helper) ParseE(strTime string) (carbon.Carbon, error) {
	if strTime == "" {
		return carbon.Carbon{}, ErrInvalidTime
	}

	parsed := t.parse(strTime)
	if parsed.Error != nil {
		return carbon.Carbon{}, fmt.Errorf("%w: %w", ErrInvalidTime, parsed.Error)
	}
	return *parsed, nil
}

// ParseLayoutE 按指定 Go 时间布局严格解析时间。
func (t Helper) ParseLayoutE(strTime, layout string) (carbon.Carbon, error) {
	if strTime == "" {
		return carbon.Carbon{}, ErrInvalidTime
	}
	if layout == "" {
		return carbon.Carbon{}, ErrInvalidLayout
	}
	parsed := carbon.ParseByLayout(strTime, layout, t.CurrentTimezone())
	if parsed.Error != nil {
		return carbon.Carbon{}, fmt.Errorf("%w: %w", ErrInvalidTime, parsed.Error)
	}
	return *parsed, nil
}

// StrToCarbonPtr 解析 strTime，解析失败时返回 nil。
func (t Helper) StrToCarbonPtr(strTime string) *carbon.Carbon {
	dstTime, err := t.ParseE(strTime)
	if err != nil {
		return nil
	}
	return &dstTime
}

// StrPtrToCarbonPtr 解析字符串指针，输入为空或无效时返回 nil。
func (t Helper) StrPtrToCarbonPtr(strTime *string) *carbon.Carbon {
	if strTime == nil {
		return nil
	}
	return t.StrToCarbonPtr(*strTime)
}

// StartOfDay 返回指定日期的开始时间，输入无效时返回 nil。
func (t Helper) StartOfDay(strTime *string) *carbon.Carbon {
	if strTime == nil {
		return nil
	}
	start, err := t.StartOfDayE(*strTime)
	if err != nil {
		return nil
	}
	return &start
}

// EndOfDay 返回指定日期的结束时间，输入无效时返回 nil。
func (t Helper) EndOfDay(strTime *string) *carbon.Carbon {
	if strTime == nil {
		return nil
	}
	end, err := t.EndOfDayE(*strTime)
	if err != nil {
		return nil
	}
	return &end
}

// StartOfDayE 解析 strTime 并返回当天开始时间。
func (t Helper) StartOfDayE(strTime string) (carbon.Carbon, error) {
	parsed, err := t.ParseE(strTime)
	if err != nil {
		return carbon.Carbon{}, err
	}
	return *parsed.StartOfDay(), nil
}

// EndOfDayE 解析 strTime 并返回当天结束时间。
func (t Helper) EndOfDayE(strTime string) (carbon.Carbon, error) {
	parsed, err := t.ParseE(strTime)
	if err != nil {
		return carbon.Carbon{}, err
	}
	return *parsed.EndOfDay(), nil
}

// DayHalfOpenRangeE 解析 strTime，返回当天的半开区间 [start, end)。
// end 是次日零点，适合数据库时间范围查询。
func (t Helper) DayHalfOpenRangeE(strTime string) (start, end carbon.Carbon, err error) {
	start, err = t.StartOfDayE(strTime)
	if err != nil {
		return carbon.Carbon{}, carbon.Carbon{}, err
	}
	end = *start.AddDay()
	return start, end, nil
}

// NowFormatTime 使用 layout 格式化当前时间。
func (t Helper) NowFormatTime(layout string) string {
	return t.NowCarbonPtr().StdTime().Format(layout)
}

// CarbonToDateTime 将 at 格式化为日期时间字符串。
func (t Helper) CarbonToDateTime(at carbon.Carbon) string {
	return at.ToDateTimeString(t.CurrentTimezone())
}

// CarbonPtrToDateTimePtr 格式化 at，at 为 nil 时返回 nil。
func (t Helper) CarbonPtrToDateTimePtr(at *carbon.Carbon) *string {
	if at == nil {
		return nil
	}
	str := t.CarbonToDateTime(*at)
	return &str
}

// CarbonToDate 将 at 格式化为日期字符串。
func (t Helper) CarbonToDate(at carbon.Carbon) string {
	return at.ToDateString(t.CurrentTimezone())
}

// CarbonPtrToDate 格式化 at，at 为 nil 时返回空字符串。
func (t Helper) CarbonPtrToDate(at *carbon.Carbon) string {
	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at)
}

// AddDayToDate 为 at 增加指定天数并返回日期字符串。
func (t Helper) AddDayToDate(at *carbon.Carbon, days int) string {
	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at.AddDays(days))
}

// PStrToPCarbonDate 解析 str 并返回当天的开始时间。
// Deprecated: 请使用 StartOfDay 或 StartOfDayE。
func (t Helper) PStrToPCarbonDate(str *string) *carbon.Carbon {
	return t.StartOfDay(str)
}

// NowUnixNano 返回当前时间的纳秒级 Unix 时间戳。
func (t Helper) NowUnixNano() int64 {
	return t.currentTime().UnixNano()
}

// NowUnixMilli 返回当前时间的毫秒级 Unix 时间戳。
func (t Helper) NowUnixMilli() int64 {
	return t.currentTime().UnixMilli()
}

// NowUnix 返回当前时间的秒级 Unix 时间戳。
func (t Helper) NowUnix() int64 {
	return t.currentTime().Unix()
}

// NowNanosecond 返回当前秒内的纳秒偏移量。
func (t Helper) NowNanosecond() int {
	return t.currentTime().Nanosecond()
}

// NowAddSeconds 在当前时间基础上增加 d 秒。
func (t Helper) NowAddSeconds(d int) stdtime.Time {
	return t.currentTime().Add(stdtime.Duration(d) * stdtime.Second)
}

// NowAddMinutes 在当前时间基础上增加 d 分钟。
func (t Helper) NowAddMinutes(d int) stdtime.Time {
	return t.currentTime().Add(stdtime.Duration(d) * stdtime.Minute)
}

// NowAddHours 在当前时间基础上增加 d 小时。
func (t Helper) NowAddHours(d int) stdtime.Time {
	return t.currentTime().Add(stdtime.Duration(d) * stdtime.Hour)
}

// NowAddDays 在当前时间基础上增加 d 天。
func (t Helper) NowAddDays(d int) stdtime.Time {
	return t.NowCarbonPtr().AddDays(d).StdTime()
}

// NowAddMonths 在当前时间基础上增加 d 个月。
// 本方法保留 Go time.AddDate 的月末溢出语义；需要截断到目标月月末时请使用 NowAddMonthsNoOverflow。
func (t Helper) NowAddMonths(d int) stdtime.Time {
	return t.NowCarbonPtr().AddMonths(d).StdTime()
}

// NowAddMonthsNoOverflow 在当前时间基础上增加 d 个月，并将溢出日期截断到目标月月末。
func (t Helper) NowAddMonthsNoOverflow(d int) stdtime.Time {
	return t.NowCarbonPtr().AddMonthsNoOverflow(d).StdTime()
}

// TodayRange 返回当天的开始和结束时间。
func (t Helper) TodayRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := nowCarbon.EndOfDay()
	return *start, *end
}

// TodayHalfOpenRange 返回当天的半开区间 [start, end)，end 是次日零点。
func (t Helper) TodayHalfOpenRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := start.AddDay()
	return *start, *end
}

// TodyRange 返回当天的开始和结束时间。
// Deprecated: 使用 TodayRange。
func (t Helper) TodyRange() (carbon.Carbon, carbon.Carbon) {
	return t.TodayRange()
}

// CurrDayStartEnd 返回今天开始和结束时间的指针。
func (t Helper) CurrDayStartEnd() (startAt *carbon.Carbon, endAt *carbon.Carbon) {
	start, end := t.TodayRange()
	return &start, &end
}

// YesterdayRange 返回昨天的开始和结束时间。
func (t Helper) YesterdayRange() (carbon.Carbon, carbon.Carbon) {
	yesterday := t.NowCarbonPtr().SubDay()
	start := yesterday.StartOfDay()
	end := yesterday.EndOfDay()
	return *start, *end
}

// YdayRange 返回昨天的开始和结束时间。
// Deprecated: 请使用 YesterdayRange。
func (t Helper) YdayRange() (carbon.Carbon, carbon.Carbon) {
	return t.YesterdayRange()
}

// YesterdayHalfOpenRange 返回昨天的半开区间 [start, end)，end 是今天零点。
func (t Helper) YesterdayHalfOpenRange() (carbon.Carbon, carbon.Carbon) {
	todayStart := t.NowCarbonPtr().StartOfDay()
	start := todayStart.SubDay()
	return *start, *todayStart
}

// PreviousMonthRange 返回上个月的开始和结束时间。
func (t Helper) PreviousMonthRange() (carbon.Carbon, carbon.Carbon) {
	currentMonthStart := t.NowCarbonPtr().StartOfMonth()
	start := currentMonthStart.SubMonth()
	end := start.EndOfMonth()
	return *start, *end
}

// PreMonthRange 返回上个月的开始和结束时间。
// Deprecated: 请使用 PreviousMonthRange。
func (t Helper) PreMonthRange() (*carbon.Carbon, *carbon.Carbon) {
	start, end := t.PreviousMonthRange()
	return &start, &end
}

// PreviousMonthHalfOpenRange 返回上个月的半开区间 [start, end)，end 是本月月初。
func (t Helper) PreviousMonthHalfOpenRange() (carbon.Carbon, carbon.Carbon) {
	end := t.NowCarbonPtr().StartOfMonth()
	start := end.SubMonth()
	return *start, *end
}

// TimestampMilliseconds 返回当前时间的毫秒级 Unix 时间戳。
func (t Helper) TimestampMilliseconds() int64 {
	return t.NowUnixMilli()
}

// GetTimeStampMilsecd 返回当前时间的毫秒级 Unix 时间戳。
// Deprecated: 请使用 TimestampMilliseconds。
func (t Helper) GetTimeStampMilsecd() int64 {
	return t.TimestampMilliseconds()
}

// GetTimeStampMilliseconds 返回当前时间的毫秒级 Unix 时间戳。
// Deprecated: 请使用 TimestampMilliseconds。
func (t Helper) GetTimeStampMilliseconds() int64 {
	return t.TimestampMilliseconds()
}
