package time

import stdtime "time"

type calendarDate struct {
	year  int
	month stdtime.Month
	day   int
}

func dateInLocation(value stdtime.Time, location *stdtime.Location) calendarDate {
	year, month, day := value.In(location).Date()
	return calendarDate{year: year, month: month, day: day}
}

// ordinal 返回公历日期相对 1970-01-01 的日序号，不受夏令时影响。
func (date calendarDate) ordinal() int64 {
	year := int64(date.year)
	month := int64(date.month)
	if month <= 2 {
		year--
	}

	era := year / 400
	if year < 0 {
		era = (year - 399) / 400
	}
	yearOfEra := year - era*400
	adjustedMonth := month - 3
	if month <= 2 {
		adjustedMonth = month + 9
	}
	dayOfYear := (153*adjustedMonth+2)/5 + int64(date.day) - 1
	dayOfEra := yearOfEra*365 + yearOfEra/4 - yearOfEra/100 + dayOfYear
	return era*146097 + dayOfEra - 719468
}

func (date calendarDate) weekday() stdtime.Weekday {
	// 1970-01-01 是星期四。
	weekday := (date.ordinal() + int64(stdtime.Thursday)) % 7
	if weekday < 0 {
		weekday += 7
	}
	return stdtime.Weekday(weekday)
}

func isWeekday(weekday stdtime.Weekday) bool {
	return weekday != stdtime.Saturday && weekday != stdtime.Sunday
}

func (t Helper) location() *stdtime.Location {
	location, err := stdtime.LoadLocation(t.CurrentTimezone())
	if err == nil {
		return location
	}

	// Helper 的时区字段不对包外开放，正常只能通过已校验的 WithTimezone 设置。
	// 这里仍为包内构造的异常值提供稳定兜底，避免意外使用主机的 Local 时区。
	location, err = stdtime.LoadLocation(defaultTimezone)
	if err == nil {
		return location
	}
	return stdtime.FixedZone(defaultTimezone, 8*60*60)
}

func holidaySet(location *stdtime.Location, holidays []stdtime.Time) map[calendarDate]struct{} {
	result := make(map[calendarDate]struct{}, len(holidays))
	for _, holiday := range holidays {
		result[dateInLocation(holiday, location)] = struct{}{}
	}
	return result
}

func isBusinessDate(date calendarDate, holidays map[calendarDate]struct{}) bool {
	if !isWeekday(date.weekday()) {
		return false
	}
	_, isHoliday := holidays[date]
	return !isHoliday
}

// IsBusinessDay 报告 value 是否为工作日。
// 周六、周日以及 holidays 指定的日期均为非工作日；所有时间先转换到 Helper 配置时区再取日期。
func (t Helper) IsBusinessDay(value stdtime.Time, holidays ...stdtime.Time) bool {
	location := t.location()
	return isBusinessDate(dateInLocation(value, location), holidaySet(location, holidays))
}

// AddBusinessDays 在 value 上增加 days 个工作日，并保留本地时分秒。
// days 可以为负数；起始日期不计数，返回值使用 Helper 配置时区。
// 周六、周日以及 holidays 指定的日期会被跳过。
func (t Helper) AddBusinessDays(value stdtime.Time, days int, holidays ...stdtime.Time) stdtime.Time {
	location := t.location()
	current := value.In(location)
	if days == 0 {
		return current
	}

	direction := 1
	remaining := uint(days)
	if days < 0 {
		direction = -1
		// -(days+1)+1 可避免 days 为最小 int 时取反溢出。
		remaining = uint(-(days + 1)) + 1
	}
	holidayValues := holidaySet(location, holidays)

	for remaining > 0 {
		// 没有节假日且当前是工作日时，可一次跳过完整工作周。
		if len(holidayValues) == 0 && isWeekday(current.Weekday()) && remaining >= 5 {
			weeks := remaining / 5
			maxWeeks := uint((int(^uint(0)>>1) - 31) / 7)
			if weeks > maxWeeks {
				weeks = maxWeeks
			}
			current = current.AddDate(0, 0, direction*int(weeks)*7)
			remaining -= weeks * 5
			continue
		}

		current = current.AddDate(0, 0, direction)
		if isBusinessDate(dateInLocation(current, location), holidayValues) {
			remaining--
		}
	}
	return current
}

// CalendarDaysBetween 返回 start 到 end 的有符号自然日差，忽略时分秒。
// 两个时间先转换到 Helper 配置时区，再按公历日期计算，因此不受夏令时影响。
// end 日期早于 start 日期时返回负数；日期相同时返回零。
func (t Helper) CalendarDaysBetween(start, end stdtime.Time) int {
	location := t.location()
	startDate := dateInLocation(start, location)
	endDate := dateInLocation(end, location)
	return int(endDate.ordinal() - startDate.ordinal())
}

// BusinessDaysBetween 返回 start 到 end 的有符号工作日数，计数区间为 (start, end]。
// start、end 和 holidays 都会先转换到 Helper 配置时区再取日期。
// end 日期早于 start 日期时交换边界计算并返回负数；重复节假日只扣除一次。
func (t Helper) BusinessDaysBetween(start, end stdtime.Time, holidays ...stdtime.Time) int {
	location := t.location()
	startDate := dateInLocation(start, location)
	endDate := dateInLocation(end, location)
	if startDate == endDate {
		return 0
	}

	sign := 1
	if endDate.ordinal() < startDate.ordinal() {
		startDate, endDate = endDate, startDate
		sign = -1
	}

	totalDays := endDate.ordinal() - startDate.ordinal()
	businessDays := (totalDays / 7) * 5
	weekday := startDate.weekday()
	for day := int64(0); day < totalDays%7; day++ {
		weekday = (weekday + 1) % 7
		if isWeekday(weekday) {
			businessDays++
		}
	}

	startOrdinal, endOrdinal := startDate.ordinal(), endDate.ordinal()
	for holiday := range holidaySet(location, holidays) {
		holidayOrdinal := holiday.ordinal()
		if holidayOrdinal > startOrdinal && holidayOrdinal <= endOrdinal && isWeekday(holiday.weekday()) {
			businessDays--
		}
	}
	return sign * int(businessDays)
}
