package time

import "github.com/dromara/carbon/v2"

// CurrentLayout 返回 Helper 配置时区。
// Deprecated: 请使用 CurrentTimezone。
func (t Helper) CurrentLayout() string { return t.CurrentTimezone() }

// GetCurrLayout 返回 Helper 配置时区。
// Deprecated: 请使用 CurrentTimezone。
func (t Helper) GetCurrLayout() string { return t.CurrentTimezone() }

// ParseE 使用 Helper 配置时区解析时间。
// Deprecated: 请使用 Parse。
func (t Helper) ParseE(strTime string) (carbon.Carbon, error) { return t.Parse(strTime) }

// ParseLayoutE 按指定 Go 时间布局解析时间。
// Deprecated: 请使用 ParseLayout。
func (t Helper) ParseLayoutE(strTime, layout string) (carbon.Carbon, error) {
	return t.ParseLayout(strTime, layout)
}

// StrToCarbon 使用 Helper 配置时区解析时间，错误保存在返回值的 Error 字段中。
// Deprecated: 请使用 Parse 并处理其返回错误。
func (t Helper) StrToCarbon(strTime string) carbon.Carbon { return *t.parseCarbon(strTime) }

// StrToCarbonPtr 解析时间，失败时返回 nil。
// Deprecated: 请使用 Parse 并处理其返回错误。
func (t Helper) StrToCarbonPtr(strTime string) *carbon.Carbon {
	parsed, err := t.Parse(strTime)
	if err != nil {
		return nil
	}
	return &parsed
}

// StrPtrToCarbonPtr 解析字符串指针，输入为空或无效时返回 nil。
// Deprecated: 请使用 ParseOptionalCarbon。
func (t Helper) StrPtrToCarbonPtr(strTime *string) *carbon.Carbon {
	return t.ParseOptionalCarbon(strTime)
}

// StartOfDay 返回指定日期的开始时间，输入无效时返回 nil。
// Deprecated: 请使用 ParseOptionalStartOfDay；需要错误信息时使用 ParseStartOfDay。
func (t Helper) StartOfDay(strTime *string) *carbon.Carbon {
	return t.ParseOptionalStartOfDay(strTime)
}

// EndOfDay 返回指定日期的结束时间，输入无效时返回 nil。
// Deprecated: 请使用 ParseOptionalEndOfDay；需要错误信息时使用 ParseEndOfDay。
func (t Helper) EndOfDay(strTime *string) *carbon.Carbon {
	return t.ParseOptionalEndOfDay(strTime)
}

// StartOfDayE 解析时间并返回当天开始时间。
// Deprecated: 请使用 ParseStartOfDay。
func (t Helper) StartOfDayE(strTime string) (carbon.Carbon, error) {
	return t.ParseStartOfDay(strTime)
}

// EndOfDayE 解析时间并返回当天结束时间。
// Deprecated: 请使用 ParseEndOfDay。
func (t Helper) EndOfDayE(strTime string) (carbon.Carbon, error) {
	return t.ParseEndOfDay(strTime)
}

// DayHalfOpenRangeE 解析时间并返回当天的半开区间 [start, end)。
// Deprecated: 请使用 ParseDayHalfOpenRange。
func (t Helper) DayHalfOpenRangeE(strTime string) (start, end carbon.Carbon, err error) {
	return t.ParseDayHalfOpenRange(strTime)
}

// NowFormatTime 使用 layout 格式化当前时间。
// Deprecated: 请使用 FormatNow。
func (t Helper) NowFormatTime(layout string) string { return t.FormatNow(layout) }

// CarbonToDateTime 将 at 格式化为日期时间字符串。
// Deprecated: 请使用 FormatDateTime。
func (t Helper) CarbonToDateTime(at carbon.Carbon) string { return t.FormatDateTime(at) }

// CarbonPtrToDateTimePtr 格式化 at，at 为 nil 时返回 nil。
// Deprecated: 请使用 FormatOptionalDateTime。
func (t Helper) CarbonPtrToDateTimePtr(at *carbon.Carbon) *string {
	return t.FormatOptionalDateTime(at)
}

// CarbonToDate 将 at 格式化为日期字符串。
// Deprecated: 请使用 FormatDate。
func (t Helper) CarbonToDate(at carbon.Carbon) string { return t.FormatDate(at) }

// CarbonPtrToDate 格式化 at，at 为 nil 时返回空字符串。
// Deprecated: 请检查 nil 后使用 FormatDate。
func (t Helper) CarbonPtrToDate(at *carbon.Carbon) string {
	if at == nil {
		return ""
	}
	return t.FormatDate(*at)
}

// AddDayToDate 为 at 增加指定天数并返回日期字符串。
// Deprecated: 请检查 nil 后使用 AddDaysAndFormatDate。
func (t Helper) AddDayToDate(at *carbon.Carbon, days int) string {
	if at == nil {
		return ""
	}
	return t.AddDaysAndFormatDate(*at, days)
}

// PStrToPCarbonDate 解析 str 并返回当天的开始时间。
// Deprecated: 请使用 ParseOptionalStartOfDay；需要错误信息时使用 ParseStartOfDay。
func (t Helper) PStrToPCarbonDate(str *string) *carbon.Carbon {
	return t.ParseOptionalStartOfDay(str)
}

// TodyRange 返回当天的开始和结束时间。
// Deprecated: 请使用 TodayRange。
func (t Helper) TodyRange() (carbon.Carbon, carbon.Carbon) { return t.TodayRange() }

// CurrDayStartEnd 返回今天开始和结束时间的指针。
// Deprecated: 请使用 TodayRange。
func (t Helper) CurrDayStartEnd() (startAt *carbon.Carbon, endAt *carbon.Carbon) {
	start, end := t.TodayRange()
	return &start, &end
}

// YdayRange 返回昨天的开始和结束时间。
// Deprecated: 请使用 YesterdayRange。
func (t Helper) YdayRange() (carbon.Carbon, carbon.Carbon) { return t.YesterdayRange() }

// PreMonthRange 返回上个月的开始和结束时间。
// Deprecated: 请使用 PreviousMonthRange。
func (t Helper) PreMonthRange() (*carbon.Carbon, *carbon.Carbon) {
	start, end := t.PreviousMonthRange()
	return &start, &end
}

// TimestampMilliseconds 返回当前时刻的毫秒级 Unix 时间戳。
// Deprecated: 请使用 NowUnixMilli。
func (t Helper) TimestampMilliseconds() int64 { return t.NowUnixMilli() }

// GetTimeStampMilsecd 返回当前时刻的毫秒级 Unix 时间戳。
// Deprecated: 请使用 NowUnixMilli。
func (t Helper) GetTimeStampMilsecd() int64 { return t.NowUnixMilli() }

// GetTimeStampMilliseconds 返回当前时刻的毫秒级 Unix 时间戳。
// Deprecated: 请使用 NowUnixMilli。
func (t Helper) GetTimeStampMilliseconds() int64 { return t.NowUnixMilli() }
