package time_test

import (
	"errors"
	"strings"
	"testing"
	stdtime "time"

	timeutil "github.com/lily0749labs/goutils/time"
)

var _ timeutil.Helper = timeutil.Time

func frozenHelper(t *testing.T, now stdtime.Time, timezone string) timeutil.Helper {
	t.Helper()

	helper, err := timeutil.Time.WithTimezone(timezone)
	if err != nil {
		t.Fatalf("WithTimezone(%q) error = %v", timezone, err)
	}
	return helper.WithNow(func() stdtime.Time { return now })
}

func requireTime(t *testing.T, got, want stdtime.Time) {
	t.Helper()

	if !got.Equal(want) {
		t.Fatalf("time = %s, want %s", got.Format(stdtime.RFC3339Nano), want.Format(stdtime.RFC3339Nano))
	}
	if got.Location().String() != want.Location().String() {
		t.Fatalf("location = %q, want %q", got.Location(), want.Location())
	}
}

func TestHelperTimezoneAndNow(t *testing.T) {
	t.Parallel()

	if got := (timeutil.Helper{}).CurrentTimezone(); got != "Asia/Shanghai" {
		t.Fatalf("zero Helper timezone = %q, want Asia/Shanghai", got)
	}
	if got := timeutil.Time.CurrentLayout(); got != timeutil.Time.CurrentTimezone() {
		t.Fatalf("CurrentLayout() = %q, CurrentTimezone() = %q", got, timeutil.Time.CurrentTimezone())
	}

	if _, err := timeutil.Time.WithTimezone("Not/A_Timezone"); !errors.Is(err, timeutil.ErrInvalidTimezone) {
		t.Fatalf("WithTimezone() error = %v, want ErrInvalidTimezone", err)
	}

	fixed := stdtime.Date(2025, stdtime.January, 2, 3, 4, 5, 6, stdtime.UTC)
	helper := frozenHelper(t, fixed, "Asia/Shanghai")
	requireTime(t, helper.NowTime(), fixed)
	requireTime(t, helper.NowInTimezone(), stdtime.Date(2025, stdtime.January, 2, 11, 4, 5, 6, mustLocation(t, "Asia/Shanghai")))

	if got := helper.WithNow(nil).NowTime(); got.IsZero() {
		t.Fatal("WithNow(nil).NowTime() returned zero time")
	}
}

func TestNowTimePreservesMonotonicClock(t *testing.T) {
	t.Parallel()

	if got := timeutil.Time.NowTime().String(); !strings.Contains(got, "m=") {
		t.Fatalf("NowTime() = %q, want a monotonic clock reading", got)
	}
}

func TestParseEUsesInjectedClockForRelativeValues(t *testing.T) {
	t.Parallel()

	location := mustLocation(t, "Asia/Shanghai")
	fixed := stdtime.Date(2025, stdtime.March, 31, 23, 30, 45, 123, location)
	helper := frozenHelper(t, fixed, location.String())

	tests := []struct {
		value string
		want  stdtime.Time
	}{
		{value: "now", want: fixed},
		{value: "yesterday", want: stdtime.Date(2025, stdtime.March, 30, 23, 30, 45, 123, location)},
		{value: "tomorrow", want: stdtime.Date(2025, stdtime.April, 1, 23, 30, 45, 123, location)},
	}

	for _, tc := range tests {
		t.Run(tc.value, func(t *testing.T) {
			parsed, err := helper.ParseE(tc.value)
			if err != nil {
				t.Fatalf("ParseE(%q) error = %v", tc.value, err)
			}
			requireTime(t, parsed.StdTime(), tc.want)
			compatible := helper.StrToCarbon(tc.value)
			requireTime(t, compatible.StdTime(), tc.want)
		})
	}
}

func TestParseErrorsAndLayout(t *testing.T) {
	t.Parallel()

	if _, err := timeutil.Time.ParseE(""); !errors.Is(err, timeutil.ErrInvalidTime) {
		t.Fatalf("ParseE(empty) error = %v, want ErrInvalidTime", err)
	}
	if _, err := timeutil.Time.ParseE("not-a-time"); !errors.Is(err, timeutil.ErrInvalidTime) {
		t.Fatalf("ParseE(invalid) error = %v, want ErrInvalidTime", err)
	}
	if _, err := timeutil.Time.ParseLayoutE("2025-01-02", ""); !errors.Is(err, timeutil.ErrInvalidLayout) {
		t.Fatalf("ParseLayoutE(empty layout) error = %v, want ErrInvalidLayout", err)
	}
	if _, err := timeutil.Time.ParseLayoutE("", stdtime.DateOnly); !errors.Is(err, timeutil.ErrInvalidTime) {
		t.Fatalf("ParseLayoutE(empty time) error = %v, want ErrInvalidTime", err)
	}
	if _, err := timeutil.Time.ParseLayoutE("2025/01/02", stdtime.DateOnly); !errors.Is(err, timeutil.ErrInvalidTime) {
		t.Fatalf("ParseLayoutE(mismatch) error = %v, want ErrInvalidTime", err)
	}

	parsed, err := timeutil.Time.ParseLayoutE("2025-01-02", stdtime.DateOnly)
	if err != nil {
		t.Fatalf("ParseLayoutE(valid) error = %v", err)
	}
	if got := parsed.ToDateString(); got != "2025-01-02" {
		t.Fatalf("ParseLayoutE(valid) = %q, want 2025-01-02", got)
	}
}

func TestFormattingAndCompatibilityMethods(t *testing.T) {
	t.Parallel()

	utcHelper := frozenHelper(t, stdtime.Date(2025, stdtime.January, 2, 3, 4, 5, 6, stdtime.UTC), "UTC")
	shanghaiHelper, err := utcHelper.WithTimezone("Asia/Shanghai")
	if err != nil {
		t.Fatalf("WithTimezone(Asia/Shanghai) error = %v", err)
	}

	if got := shanghaiHelper.GetCurrLayout(); got != "Asia/Shanghai" {
		t.Fatalf("GetCurrLayout() = %q, want Asia/Shanghai", got)
	}
	if got := shanghaiHelper.NowFormatTime(stdtime.DateTime); got != "2025-01-02 11:04:05" {
		t.Fatalf("NowFormatTime() = %q, want 2025-01-02 11:04:05", got)
	}

	at, err := utcHelper.ParseLayoutE("2025-01-02 03:04:05", stdtime.DateTime)
	if err != nil {
		t.Fatalf("ParseLayoutE() error = %v", err)
	}
	if got := shanghaiHelper.CarbonToDateTime(at); got != "2025-01-02 11:04:05" {
		t.Fatalf("CarbonToDateTime() = %q, want 2025-01-02 11:04:05", got)
	}
	if got := shanghaiHelper.CarbonToDate(at); got != "2025-01-02" {
		t.Fatalf("CarbonToDate() = %q, want 2025-01-02", got)
	}
	if got := shanghaiHelper.CarbonPtrToDateTimePtr(&at); got == nil || *got != "2025-01-02 11:04:05" {
		t.Fatalf("CarbonPtrToDateTimePtr() = %v, want 2025-01-02 11:04:05", got)
	}
	if got := shanghaiHelper.CarbonPtrToDateTimePtr(nil); got != nil {
		t.Fatalf("CarbonPtrToDateTimePtr(nil) = %v, want nil", got)
	}
	if got := shanghaiHelper.CarbonPtrToDate(&at); got != "2025-01-02" {
		t.Fatalf("CarbonPtrToDate() = %q, want 2025-01-02", got)
	}
	if got := shanghaiHelper.CarbonPtrToDate(nil); got != "" {
		t.Fatalf("CarbonPtrToDate(nil) = %q, want empty", got)
	}
	if got := shanghaiHelper.AddDayToDate(&at, 1); got != "2025-01-03" {
		t.Fatalf("AddDayToDate() = %q, want 2025-01-03", got)
	}
	if got := shanghaiHelper.AddDayToDate(nil, 1); got != "" {
		t.Fatalf("AddDayToDate(nil) = %q, want empty", got)
	}

	input := "2025-01-02 03:04:05"
	if got := utcHelper.StrToCarbonPtr(input); got == nil || got.ToDateTimeString() != input {
		t.Fatalf("StrToCarbonPtr() = %v, want %s", got, input)
	}
	if got := utcHelper.StrToCarbonPtr("invalid"); got != nil {
		t.Fatalf("StrToCarbonPtr(invalid) = %v, want nil", got)
	}
	if got := utcHelper.StrPtrToCarbonPtr(&input); got == nil || got.ToDateTimeString() != input {
		t.Fatalf("StrPtrToCarbonPtr() = %v, want %s", got, input)
	}
	if got := utcHelper.StrPtrToCarbonPtr(nil); got != nil {
		t.Fatalf("StrPtrToCarbonPtr(nil) = %v, want nil", got)
	}
	if got := utcHelper.PStrToPCarbonDate(&input); got == nil || got.ToDateTimeString() != "2025-01-02 00:00:00" {
		t.Fatalf("PStrToPCarbonDate() = %v, want start of day", got)
	}

	start, end := utcHelper.TodayRange()
	legacyStart, legacyEnd := utcHelper.TodyRange()
	requireTime(t, legacyStart.StdTime(), start.StdTime())
	requireTime(t, legacyEnd.StdTime(), end.StdTime())
	pointerStart, pointerEnd := utcHelper.CurrDayStartEnd()
	requireTime(t, pointerStart.StdTime(), start.StdTime())
	requireTime(t, pointerEnd.StdTime(), end.StdTime())

	if got := utcHelper.TimestampMilliseconds(); got != utcHelper.NowUnixMilli() {
		t.Fatalf("TimestampMilliseconds() = %d, want %d", got, utcHelper.NowUnixMilli())
	}
	if got := utcHelper.GetTimeStampMilsecd(); got != utcHelper.NowUnixMilli() {
		t.Fatalf("GetTimeStampMilsecd() = %d, want %d", got, utcHelper.NowUnixMilli())
	}
	if got := utcHelper.GetTimeStampMilliseconds(); got != utcHelper.NowUnixMilli() {
		t.Fatalf("GetTimeStampMilliseconds() = %d, want %d", got, utcHelper.NowUnixMilli())
	}
}

func TestDayBoundariesAndHalfOpenRange(t *testing.T) {
	t.Parallel()

	location := mustLocation(t, "Asia/Shanghai")
	helper := frozenHelper(t, stdtime.Date(2025, stdtime.February, 10, 12, 0, 0, 0, location), location.String())
	input := "2025-02-10 12:34:56"

	start, err := helper.StartOfDayE(input)
	if err != nil {
		t.Fatalf("StartOfDayE() error = %v", err)
	}
	end, err := helper.EndOfDayE(input)
	if err != nil {
		t.Fatalf("EndOfDayE() error = %v", err)
	}
	requireTime(t, start.StdTime(), stdtime.Date(2025, stdtime.February, 10, 0, 0, 0, 0, location))
	requireTime(t, end.StdTime(), stdtime.Date(2025, stdtime.February, 10, 23, 59, 59, 999999999, location))

	rangeStart, rangeEnd, err := helper.DayHalfOpenRangeE(input)
	if err != nil {
		t.Fatalf("DayHalfOpenRangeE() error = %v", err)
	}
	requireTime(t, rangeStart.StdTime(), start.StdTime())
	requireTime(t, rangeEnd.StdTime(), stdtime.Date(2025, stdtime.February, 11, 0, 0, 0, 0, location))

	if got := helper.StartOfDay(nil); got != nil {
		t.Fatalf("StartOfDay(nil) = %v, want nil", got)
	}
	invalid := "not-a-time"
	if got := helper.EndOfDay(&invalid); got != nil {
		t.Fatalf("EndOfDay(invalid) = %v, want nil", got)
	}
	if _, _, err := helper.DayHalfOpenRangeE(invalid); !errors.Is(err, timeutil.ErrInvalidTime) {
		t.Fatalf("DayHalfOpenRangeE(invalid) error = %v, want ErrInvalidTime", err)
	}
}

func TestPreviousMonthRangeAtMonthEnd(t *testing.T) {
	t.Parallel()

	location := mustLocation(t, "Asia/Shanghai")
	tests := []struct {
		name      string
		now       stdtime.Time
		wantStart stdtime.Time
		wantEnd   stdtime.Time
	}{
		{
			name:      "march 31 after common february",
			now:       stdtime.Date(2025, stdtime.March, 31, 12, 0, 0, 0, location),
			wantStart: stdtime.Date(2025, stdtime.February, 1, 0, 0, 0, 0, location),
			wantEnd:   stdtime.Date(2025, stdtime.February, 28, 23, 59, 59, 999999999, location),
		},
		{
			name:      "march 31 after leap february",
			now:       stdtime.Date(2024, stdtime.March, 31, 12, 0, 0, 0, location),
			wantStart: stdtime.Date(2024, stdtime.February, 1, 0, 0, 0, 0, location),
			wantEnd:   stdtime.Date(2024, stdtime.February, 29, 23, 59, 59, 999999999, location),
		},
		{
			name:      "january crosses year",
			now:       stdtime.Date(2025, stdtime.January, 15, 12, 0, 0, 0, location),
			wantStart: stdtime.Date(2024, stdtime.December, 1, 0, 0, 0, 0, location),
			wantEnd:   stdtime.Date(2024, stdtime.December, 31, 23, 59, 59, 999999999, location),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			helper := frozenHelper(t, tc.now, location.String())
			start, end := helper.PreviousMonthRange()
			requireTime(t, start.StdTime(), tc.wantStart)
			requireTime(t, end.StdTime(), tc.wantEnd)

			legacyStart, legacyEnd := helper.PreMonthRange()
			requireTime(t, legacyStart.StdTime(), tc.wantStart)
			requireTime(t, legacyEnd.StdTime(), tc.wantEnd)

			halfOpenStart, halfOpenEnd := helper.PreviousMonthHalfOpenRange()
			requireTime(t, halfOpenStart.StdTime(), tc.wantStart)
			requireTime(t, halfOpenEnd.StdTime(), tc.wantEnd.Add(stdtime.Nanosecond))
		})
	}
}

func TestMonthAdditionOverflowOptions(t *testing.T) {
	t.Parallel()

	location := mustLocation(t, "Asia/Shanghai")
	fixed := stdtime.Date(2025, stdtime.January, 31, 8, 30, 0, 0, location)
	helper := frozenHelper(t, fixed, location.String())

	requireTime(t, helper.NowAddMonths(1), stdtime.Date(2025, stdtime.March, 3, 8, 30, 0, 0, location))
	requireTime(t, helper.NowAddMonthsNoOverflow(1), stdtime.Date(2025, stdtime.February, 28, 8, 30, 0, 0, location))

	marchEnd := stdtime.Date(2025, stdtime.March, 31, 8, 30, 0, 0, location)
	marchEndHelper := frozenHelper(t, marchEnd, location.String())
	requireTime(t, marchEndHelper.NowAddMonthsNoOverflow(-1), stdtime.Date(2025, stdtime.February, 28, 8, 30, 0, 0, location))
}

func TestTodayAndYesterdayRanges(t *testing.T) {
	t.Parallel()

	location := mustLocation(t, "Asia/Shanghai")
	fixed := stdtime.Date(2025, stdtime.July, 2, 12, 0, 0, 0, location)
	helper := frozenHelper(t, fixed, location.String())

	todayStart, todayEnd := helper.TodayRange()
	requireTime(t, todayStart.StdTime(), stdtime.Date(2025, stdtime.July, 2, 0, 0, 0, 0, location))
	requireTime(t, todayEnd.StdTime(), stdtime.Date(2025, stdtime.July, 2, 23, 59, 59, 999999999, location))

	yesterdayStart, yesterdayEnd := helper.YesterdayRange()
	requireTime(t, yesterdayStart.StdTime(), stdtime.Date(2025, stdtime.July, 1, 0, 0, 0, 0, location))
	requireTime(t, yesterdayEnd.StdTime(), stdtime.Date(2025, stdtime.July, 1, 23, 59, 59, 999999999, location))

	legacyStart, legacyEnd := helper.YdayRange()
	requireTime(t, legacyStart.StdTime(), yesterdayStart.StdTime())
	requireTime(t, legacyEnd.StdTime(), yesterdayEnd.StdTime())

	halfOpenStart, halfOpenEnd := helper.YesterdayHalfOpenRange()
	requireTime(t, halfOpenStart.StdTime(), yesterdayStart.StdTime())
	requireTime(t, halfOpenEnd.StdTime(), todayStart.StdTime())
}

func TestRangesAcrossDaylightSavingTransitions(t *testing.T) {
	t.Parallel()

	location := mustLocation(t, "America/New_York")
	tests := []struct {
		name         string
		day          stdtime.Time
		previousDay  stdtime.Time
		wantDuration stdtime.Duration
	}{
		{
			name:         "spring forward",
			day:          stdtime.Date(2025, stdtime.March, 9, 12, 0, 0, 0, location),
			previousDay:  stdtime.Date(2025, stdtime.March, 8, 12, 0, 0, 0, location),
			wantDuration: 23 * stdtime.Hour,
		},
		{
			name:         "fall back",
			day:          stdtime.Date(2025, stdtime.November, 2, 12, 0, 0, 0, location),
			previousDay:  stdtime.Date(2025, stdtime.November, 1, 12, 0, 0, 0, location),
			wantDuration: 25 * stdtime.Hour,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			helper := frozenHelper(t, tc.day, location.String())
			start, end := helper.TodayHalfOpenRange()
			if got := end.StdTime().Sub(start.StdTime()); got != tc.wantDuration {
				t.Fatalf("day duration = %s, want %s", got, tc.wantDuration)
			}

			previousDayHelper := frozenHelper(t, tc.previousDay, location.String())
			if got := previousDayHelper.NowAddDays(1).Sub(tc.previousDay); got != tc.wantDuration {
				t.Fatalf("NowAddDays(1) = %s, want %s", got, tc.wantDuration)
			}
		})
	}
}

func TestInjectedClockTimestampsAndElapsedAdditions(t *testing.T) {
	t.Parallel()

	fixed := stdtime.Date(2025, stdtime.June, 1, 2, 3, 4, 567890123, stdtime.UTC)
	helper := frozenHelper(t, fixed, "UTC")

	if got := helper.NowUnix(); got != fixed.Unix() {
		t.Fatalf("NowUnix() = %d, want %d", got, fixed.Unix())
	}
	if got := helper.NowUnixMilli(); got != fixed.UnixMilli() {
		t.Fatalf("NowUnixMilli() = %d, want %d", got, fixed.UnixMilli())
	}
	if got := helper.NowUnixNano(); got != fixed.UnixNano() {
		t.Fatalf("NowUnixNano() = %d, want %d", got, fixed.UnixNano())
	}
	if got := helper.NowNanosecond(); got != fixed.Nanosecond() {
		t.Fatalf("NowNanosecond() = %d, want %d", got, fixed.Nanosecond())
	}
	requireTime(t, helper.NowAddSeconds(2), fixed.Add(2*stdtime.Second))
	requireTime(t, helper.NowAddMinutes(2), fixed.Add(2*stdtime.Minute))
	requireTime(t, helper.NowAddHours(2), fixed.Add(2*stdtime.Hour))
}

func mustLocation(t *testing.T, name string) *stdtime.Location {
	t.Helper()

	location, err := stdtime.LoadLocation(name)
	if err != nil {
		t.Fatalf("LoadLocation(%q) error = %v", name, err)
	}
	return location
}
