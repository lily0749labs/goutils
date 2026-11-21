package pbtime

import (
	"testing"
	"time"

	"github.com/dromara/carbon/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Format 按传入的时区输出；没有设置（nil）的时间戳返回空字符串；时区为 nil 时按 UTC。
func TestFormat(t *testing.T) {
	// 同一时刻 UTC 08:30:15，换算成各时区的墙上时间。
	instant := timestamppb.New(time.Date(2026, 10, 24, 8, 30, 15, 0, time.UTC))
	for _, zone := range []struct{ name, want string }{
		{"UTC", "2026-10-24 08:30:15"},
		{"Asia/Singapore", "2026-10-24 16:30:15"},
		{"America/New_York", "2026-10-24 04:30:15"},
	} {
		loc, err := time.LoadLocation(zone.name)
		if err != nil {
			t.Fatal(err)
		}
		if got := Format(instant, loc); got != zone.want {
			t.Errorf("%s：Format = %q，应为 %q", zone.name, got, zone.want)
		}
		// 没有设置的时间戳返回空字符串，不能变成 1970-01-01。
		if got := Format(nil, loc); got != "" {
			t.Errorf("%s：Format(nil) = %q，应为空字符串", zone.name, got)
		}
	}
	// 没有指定时区：按 UTC，不能 panic。
	if got := Format(instant, nil); got != "2026-10-24 08:30:15" {
		t.Errorf("Format(ts, nil) = %q，应按 UTC 输出", got)
	}
}

// FromCarbon：有效时间转成同一时刻的时间戳；nil 和零时间返回 nil（没有设置）。
func TestFromCarbon(t *testing.T) {
	instant := time.Date(2026, 10, 24, 8, 30, 15, 123, time.UTC)
	// 有效时间：时刻不变，与 carbon 自带的时区无关。
	for _, c := range []*carbon.Carbon{
		carbon.CreateFromStdTime(instant),
		carbon.CreateFromStdTime(instant, "Asia/Singapore"),
	} {
		got := FromCarbon(c)
		if got == nil || !got.AsTime().Equal(instant) {
			t.Errorf("FromCarbon(%v) = %v，应为同一时刻 %v", c, got, instant)
		}
	}
	// nil 和零时间：不设置。
	for name, c := range map[string]*carbon.Carbon{
		"nil":       nil,
		"零值":        {},
		"NewCarbon": carbon.NewCarbon(),
	} {
		if got := FromCarbon(c); got != nil {
			t.Errorf("%s：FromCarbon 应返回 nil，实际 %v", name, got)
		}
	}
}
