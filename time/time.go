package time

import (
	"time"

	"github.com/dromara/carbon/v2"
)

var (
	currLayout = carbon.Shanghai

	// Time 提供常用时间辅助方法。
	Time = helper{}
	// TimePtr 指向 Time，为兼容旧代码而保留。
	TimePtr = &Time
)

type helper struct{}

// CurrentLayout 返回当前配置的 Carbon 时区。
func (helper) CurrentLayout() string {
	return currLayout
}

// GetCurrLayout 返回当前配置的 Carbon 时区。
// Deprecated: 请使用 CurrentLayout。
func (t helper) GetCurrLayout() string {
	return t.CurrentLayout()
}

// NowCarbon 返回 Carbon 类型的当前时间。
func (t helper) NowCarbon() carbon.Carbon {
	return *carbon.Now(currLayout)
}

// NowTime 返回标准库 time.Time 类型的当前时间。
func (t helper) NowTime() time.Time {
	return t.NowCarbonPtr().StdTime()
}

// NowCarbonPtr 返回 Carbon 类型的当前时间指针。
func (t helper) NowCarbonPtr() *carbon.Carbon {
	now := t.NowCarbon()
	return &now
}

// StrToCarbon 使用当前配置的 Carbon 时区解析 strTime。
func (t helper) StrToCarbon(strTime string) carbon.Carbon {
	return *carbon.Parse(strTime, currLayout)
}

// StrToCarbonPtr 解析 strTime，解析失败时返回 nil。
func (t helper) StrToCarbonPtr(strTime string) *carbon.Carbon {
	dstTime := t.StrToCarbon(strTime)
	if dstTime.Error != nil {
		return nil
	}
	return &dstTime
}

// StrPtrToCarbonPtr 解析字符串指针，输入为空或无效时返回 nil。
func (t helper) StrPtrToCarbonPtr(strTime *string) *carbon.Carbon {
	if strTime == nil {
		return nil
	}
	return t.StrToCarbonPtr(*strTime)
}

// StartOfDay 返回指定日期的开始时间，输入无效时返回 nil。
func (t helper) StartOfDay(strTime *string) *carbon.Carbon {
	dstTime := t.StrPtrToCarbonPtr(strTime)
	if dstTime == nil {
		return nil
	}

	return dstTime.StartOfDay()
}

// EndOfDay 返回指定日期的结束时间，输入无效时返回 nil。
func (t helper) EndOfDay(strTime *string) *carbon.Carbon {
	dstTime := t.StrPtrToCarbonPtr(strTime)
	if dstTime == nil {
		return nil
	}

	return dstTime.EndOfDay()
}

// NowFormatTime 使用 layout 格式化当前时间。
func (t helper) NowFormatTime(layout string) string {
	return t.NowCarbonPtr().StdTime().Format(layout)
}

// func FormatTime(time carbon.Carbon) string {
// 	return time.StdTime().Format(currLayout)
// }

// CarbonToDateTime 将 at 格式化为日期时间字符串。
func (t helper) CarbonToDateTime(at carbon.Carbon) string {
	return at.ToDateTimeString(currLayout)
}

// CarbonPtrToDateTimePtr 格式化 at，at 为 nil 时返回 nil。
func (t helper) CarbonPtrToDateTimePtr(at *carbon.Carbon) *string {
	if at == nil {
		return nil
	}
	str := t.CarbonToDateTime(*at)
	return &str
}

// CarbonToDate 将 at 格式化为日期字符串。
func (t helper) CarbonToDate(at carbon.Carbon) string {
	return at.ToDateString(currLayout)
}

// CarbonPtrToDate 格式化 at，at 为 nil 时返回空字符串。
func (t helper) CarbonPtrToDate(at *carbon.Carbon) string {
	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at)
}

// AddDayToDate 为 at 增加指定天数并返回日期字符串。
func (t helper) AddDayToDate(at *carbon.Carbon, days int) string {
	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at.AddDays(days))
}

// func (t helper) FormatPointerTime(time *carbon.Carbon) *string {
// 	if time == nil {
// 		return nil
// 	}
// 	str := time.ToDateTimeString(currLayout) //.ToDateString(currLayout)
// 	return &str
// }

// PStrToPCarbonDate 解析 str 并返回当天的开始时间。
func (t helper) PStrToPCarbonDate(str *string) *carbon.Carbon {
	at := t.StrPtrToCarbonPtr(str)
	if at == nil {
		return nil
	}
	return at.StartOfDay()
}

// NowUnixNano 返回当前时间的纳秒级 Unix 时间戳。
func (t helper) NowUnixNano() int64 {
	return t.NowCarbonPtr().StdTime().UnixNano()
}

// NowUnixMilli 返回当前时间的毫秒级 Unix 时间戳。
func (t helper) NowUnixMilli() int64 {
	return t.NowCarbonPtr().StdTime().UnixMilli()
}

// NowUnix 返回当前时间的秒级 Unix 时间戳。
func (t helper) NowUnix() int64 {
	return t.NowCarbonPtr().StdTime().Unix()
}

// NowNanosecond 返回当前秒内的纳秒偏移量。
func (t helper) NowNanosecond() int {
	return t.NowCarbonPtr().StdTime().Nanosecond()
}

// NowAddSeconds 在当前时间基础上增加 d 秒。
func (t helper) NowAddSeconds(d int) time.Time {
	return t.NowCarbonPtr().AddSeconds(d).StdTime()
}

// NowAddMinutes 在当前时间基础上增加 d 分钟。
func (t helper) NowAddMinutes(d int) time.Time {
	return t.NowCarbonPtr().AddMinutes(d).StdTime()
}

// NowAddHours 在当前时间基础上增加 d 小时。
func (t helper) NowAddHours(d int) time.Time {
	return t.NowCarbonPtr().AddHours(d).StdTime()
}

// NowAddDays 在当前时间基础上增加 d 天。
func (t helper) NowAddDays(d int) time.Time {
	return t.NowCarbonPtr().AddDays(d).StdTime()
}

// NowAddMonths 在当前时间基础上增加 d 个月。
func (t helper) NowAddMonths(d int) time.Time {
	return t.NowCarbonPtr().AddMonths(d).StdTime()
}

// TodayRange 返回当天的开始和结束时间。
func (t helper) TodayRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := nowCarbon.EndOfDay()
	return *start, *end
}

// TodyRange 返回当天的开始和结束时间。
// Deprecated: 使用 TodayRange。
func (t helper) TodyRange() (carbon.Carbon, carbon.Carbon) {
	return t.TodayRange()
}

// CurrDayStartEnd 返回今天开始和结束时间的指针。
func (t helper) CurrDayStartEnd() (startAt *carbon.Carbon, endAt *carbon.Carbon) {
	start, end := t.TodayRange()
	return &start, &end
}

// YdayRange 返回昨天的开始和结束时间。
func (t helper) YdayRange() (carbon.Carbon, carbon.Carbon) {
	yesterday := t.NowCarbonPtr().SubDay()
	start := yesterday.StartOfDay()
	end := yesterday.EndOfDay()
	return *start, *end
}

// PreMonthRange 返回上个月的开始和结束时间。
func (t helper) PreMonthRange() (*carbon.Carbon, *carbon.Carbon) {
	previousMonth := t.NowCarbonPtr().SubMonth()
	start := previousMonth.StartOfMonth()
	end := previousMonth.EndOfMonth()
	return start, end
}

// TimestampMilliseconds 返回当前时间的毫秒级 Unix 时间戳。
func (t helper) TimestampMilliseconds() int64 {
	return t.NowUnixMilli()
}

// GetTimeStampMilsecd 返回当前时间的毫秒级 Unix 时间戳。
// Deprecated: 请使用 TimestampMilliseconds。
func (t helper) GetTimeStampMilsecd() int64 {
	return t.TimestampMilliseconds()
}

// GetTimeStampMilliseconds 返回当前时间的毫秒级 Unix 时间戳。
// Deprecated: 请使用 TimestampMilliseconds。
func (t helper) GetTimeStampMilliseconds() int64 {
	return t.TimestampMilliseconds()
}
