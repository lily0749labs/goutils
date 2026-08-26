package time

import (
	"errors"
	"fmt"
	stdtime "time"

	"github.com/dromara/carbon/v2"
)

var (
	currLayout = carbon.Shanghai

	// ErrInvalidTime 表示时间字符串为空或无法解析。
	ErrInvalidTime = errors.New("invalid time value")
	// ErrInvalidTimezone 表示时区名称无效。
	ErrInvalidTimezone = errors.New("invalid timezone")

	// Time 提供常用时间辅助方法。
	Time = time{layout: currLayout}
	// TimePtr 指向 Time，为兼容旧代码而保留。
	TimePtr = &Time
)

type time struct {
	now    func() stdtime.Time
	layout string
}

// CurrentLayout 返回当前配置的 Carbon 时区。
func (t time) CurrentLayout() string {
	if t.layout == "" {
		return currLayout
	}
	return t.layout
}

// WithNow 返回使用指定当前时间函数的独立时间门面，便于测试和任务回放。
// now 为 nil 时恢复使用系统时间。
func (t time) WithNow(now func() stdtime.Time) time {
	t.now = now
	return t
}

// WithTimezone 返回使用指定时区的独立时间门面，不修改全局 Time。
func (t time) WithTimezone(timezone string) (time, error) {
	if _, err := stdtime.LoadLocation(timezone); err != nil {
		return t, fmt.Errorf("%w: %s", ErrInvalidTimezone, timezone)
	}
	t.layout = timezone
	return t, nil
}

func (t time) currentTime() stdtime.Time {
	if t.now != nil {
		return t.now()
	}
	return stdtime.Now()
}

// GetCurrLayout 返回当前配置的 Carbon 时区。
// Deprecated: 请使用 CurrentLayout。
func (t time) GetCurrLayout() string {
	return t.CurrentLayout()
}

// NowCarbon 返回 Carbon 类型的当前时间。
func (t time) NowCarbon() carbon.Carbon {
	return *carbon.CreateFromStdTime(t.currentTime(), t.CurrentLayout())
}

// NowTime 返回标准库 time.Time 类型的当前时间。
func (t time) NowTime() stdtime.Time {
	return t.NowCarbonPtr().StdTime()
}

// NowCarbonPtr 返回 Carbon 类型的当前时间指针。
func (t time) NowCarbonPtr() *carbon.Carbon {
	now := t.NowCarbon()
	return &now
}

// StrToCarbon 使用当前配置的 Carbon 时区解析 strTime。
func (t time) StrToCarbon(strTime string) carbon.Carbon {
	return *carbon.Parse(strTime, t.CurrentLayout())
}

// ParseE 使用当前时区解析时间，失败时返回明确错误。
func (t time) ParseE(strTime string) (carbon.Carbon, error) {
	if strTime == "" {
		return carbon.Carbon{}, ErrInvalidTime
	}
	parsed := carbon.Parse(strTime, t.CurrentLayout())
	if parsed.Error != nil {
		return carbon.Carbon{}, fmt.Errorf("%w: %v", ErrInvalidTime, parsed.Error)
	}
	return *parsed, nil
}

// ParseLayoutE 按指定 Go 时间布局严格解析时间。
func (t time) ParseLayoutE(strTime, layout string) (carbon.Carbon, error) {
	if strTime == "" || layout == "" {
		return carbon.Carbon{}, ErrInvalidTime
	}
	parsed := carbon.ParseByLayout(strTime, layout, t.CurrentLayout())
	if parsed.Error != nil {
		return carbon.Carbon{}, fmt.Errorf("%w: %v", ErrInvalidTime, parsed.Error)
	}
	return *parsed, nil
}

// StrToCarbonPtr 解析 strTime，解析失败时返回 nil。
func (t time) StrToCarbonPtr(strTime string) *carbon.Carbon {
	dstTime, err := t.ParseE(strTime)
	if err != nil {
		return nil
	}
	return &dstTime
}

// StrPtrToCarbonPtr 解析字符串指针，输入为空或无效时返回 nil。
func (t time) StrPtrToCarbonPtr(strTime *string) *carbon.Carbon {
	if strTime == nil {
		return nil
	}
	return t.StrToCarbonPtr(*strTime)
}

// StartOfDay 返回指定日期的开始时间，输入无效时返回 nil。
func (t time) StartOfDay(strTime *string) *carbon.Carbon {
	dstTime := t.StrPtrToCarbonPtr(strTime)
	if dstTime == nil {
		return nil
	}

	return dstTime.StartOfDay()
}

// EndOfDay 返回指定日期的结束时间，输入无效时返回 nil。
func (t time) EndOfDay(strTime *string) *carbon.Carbon {
	dstTime := t.StrPtrToCarbonPtr(strTime)
	if dstTime == nil {
		return nil
	}

	return dstTime.EndOfDay()
}

// NowFormatTime 使用 layout 格式化当前时间。
func (t time) NowFormatTime(layout string) string {
	return t.NowCarbonPtr().StdTime().Format(layout)
}

// CarbonToDateTime 将 at 格式化为日期时间字符串。
func (t time) CarbonToDateTime(at carbon.Carbon) string {
	return at.ToDateTimeString(t.CurrentLayout())
}

// CarbonPtrToDateTimePtr 格式化 at，at 为 nil 时返回 nil。
func (t time) CarbonPtrToDateTimePtr(at *carbon.Carbon) *string {
	if at == nil {
		return nil
	}
	str := t.CarbonToDateTime(*at)
	return &str
}

// CarbonToDate 将 at 格式化为日期字符串。
func (t time) CarbonToDate(at carbon.Carbon) string {
	return at.ToDateString(t.CurrentLayout())
}

// CarbonPtrToDate 格式化 at，at 为 nil 时返回空字符串。
func (t time) CarbonPtrToDate(at *carbon.Carbon) string {
	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at)
}

// AddDayToDate 为 at 增加指定天数并返回日期字符串。
func (t time) AddDayToDate(at *carbon.Carbon, days int) string {
	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at.AddDays(days))
}

// PStrToPCarbonDate 解析 str 并返回当天的开始时间。
func (t time) PStrToPCarbonDate(str *string) *carbon.Carbon {
	at := t.StrPtrToCarbonPtr(str)
	if at == nil {
		return nil
	}
	return at.StartOfDay()
}

// NowUnixNano 返回当前时间的纳秒级 Unix 时间戳。
func (t time) NowUnixNano() int64 {
	return t.NowCarbonPtr().StdTime().UnixNano()
}

// NowUnixMilli 返回当前时间的毫秒级 Unix 时间戳。
func (t time) NowUnixMilli() int64 {
	return t.NowCarbonPtr().StdTime().UnixMilli()
}

// NowUnix 返回当前时间的秒级 Unix 时间戳。
func (t time) NowUnix() int64 {
	return t.NowCarbonPtr().StdTime().Unix()
}

// NowNanosecond 返回当前秒内的纳秒偏移量。
func (t time) NowNanosecond() int {
	return t.NowCarbonPtr().StdTime().Nanosecond()
}

// NowAddSeconds 在当前时间基础上增加 d 秒。
func (t time) NowAddSeconds(d int) stdtime.Time {
	return t.NowCarbonPtr().AddSeconds(d).StdTime()
}

// NowAddMinutes 在当前时间基础上增加 d 分钟。
func (t time) NowAddMinutes(d int) stdtime.Time {
	return t.NowCarbonPtr().AddMinutes(d).StdTime()
}

// NowAddHours 在当前时间基础上增加 d 小时。
func (t time) NowAddHours(d int) stdtime.Time {
	return t.NowCarbonPtr().AddHours(d).StdTime()
}

// NowAddDays 在当前时间基础上增加 d 天。
func (t time) NowAddDays(d int) stdtime.Time {
	return t.NowCarbonPtr().AddDays(d).StdTime()
}

// NowAddMonths 在当前时间基础上增加 d 个月。
func (t time) NowAddMonths(d int) stdtime.Time {
	return t.NowCarbonPtr().AddMonths(d).StdTime()
}

// TodayRange 返回当天的开始和结束时间。
func (t time) TodayRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := nowCarbon.EndOfDay()
	return *start, *end
}

// TodyRange 返回当天的开始和结束时间。
// Deprecated: 使用 TodayRange。
func (t time) TodyRange() (carbon.Carbon, carbon.Carbon) {
	return t.TodayRange()
}

// CurrDayStartEnd 返回今天开始和结束时间的指针。
func (t time) CurrDayStartEnd() (startAt *carbon.Carbon, endAt *carbon.Carbon) {
	start, end := t.TodayRange()
	return &start, &end
}

// YdayRange 返回昨天的开始和结束时间。
func (t time) YdayRange() (carbon.Carbon, carbon.Carbon) {
	yesterday := t.NowCarbonPtr().SubDay()
	start := yesterday.StartOfDay()
	end := yesterday.EndOfDay()
	return *start, *end
}

// PreMonthRange 返回上个月的开始和结束时间。
func (t time) PreMonthRange() (*carbon.Carbon, *carbon.Carbon) {
	previousMonth := t.NowCarbonPtr().SubMonth()
	start := previousMonth.StartOfMonth()
	end := previousMonth.EndOfMonth()
	return start, end
}

// TimestampMilliseconds 返回当前时间的毫秒级 Unix 时间戳。
func (t time) TimestampMilliseconds() int64 {
	return t.NowUnixMilli()
}

// GetTimeStampMilsecd 返回当前时间的毫秒级 Unix 时间戳。
// Deprecated: 请使用 TimestampMilliseconds。
func (t time) GetTimeStampMilsecd() int64 {
	return t.TimestampMilliseconds()
}

// GetTimeStampMilliseconds 返回当前时间的毫秒级 Unix 时间戳。
// Deprecated: 请使用 TimestampMilliseconds。
func (t time) GetTimeStampMilliseconds() int64 {
	return t.TimestampMilliseconds()
}
