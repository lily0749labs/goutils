package time

import (
	"errors"
	stdtime "time"
)

var (
	// ErrInvalidRange 表示结束时间早于开始时间，无法构成有效区间。
	ErrInvalidRange = errors.New("time range end must not be before start")
	// ErrInvalidStep 表示拆分步长小于或等于零。
	ErrInvalidStep = errors.New("time range split step must be positive")
)

// Range 表示闭区间 [Start, End]，开始和结束边界均包含在内。
type Range struct {
	Start stdtime.Time
	End   stdtime.Time
}

// NewRange 创建闭区间 [start, end]。
// end 早于 start 时返回 ErrInvalidRange；两者相等时返回单点区间。
func (Helper) NewRange(start, end stdtime.Time) (Range, error) {
	if end.Before(start) {
		return Range{}, ErrInvalidRange
	}
	return Range{Start: start, End: end}, nil
}

// Duration 返回 End.Sub(Start) 的结果。
// 调用方需要先用 Valid 判断区间是否合法。
func (r Range) Duration() stdtime.Duration { return r.End.Sub(r.Start) }

// Valid 报告 End 是否不早于 Start。
func (r Range) Valid() bool { return !r.End.Before(r.Start) }

// Contains 报告 at 是否位于闭区间内；无效区间始终返回 false。
func (r Range) Contains(at stdtime.Time) bool {
	return r.Valid() && !at.Before(r.Start) && !at.After(r.End)
}

// Overlaps 报告两个闭区间是否有交集。
// 两个区间仅边界相接时也视为重叠；任一无效时返回 false。
func (r Range) Overlaps(other Range) bool {
	return r.Valid() && other.Valid() && !r.End.Before(other.Start) && !other.End.Before(r.Start)
}

// Intersection 返回两个闭区间的交集。
// 不相交或任一无效时，第二个返回值为 false。
func (r Range) Intersection(other Range) (Range, bool) {
	if !r.Overlaps(other) {
		return Range{}, false
	}
	start := r.Start
	if other.Start.After(start) {
		start = other.Start
	}
	end := r.End
	if other.End.Before(end) {
		end = other.End
	}
	return Range{Start: start, End: end}, true
}

// Split 按固定时长 step 拆分闭区间。
// 相邻结果共享边界，最后一个区间可能短于 step；单点区间返回只包含自身的切片。
// step 小于或等于零时返回 ErrInvalidStep，无效区间返回 ErrInvalidRange。
func (r Range) Split(step stdtime.Duration) ([]Range, error) {
	if step <= 0 {
		return nil, ErrInvalidStep
	}
	if !r.Valid() {
		return nil, ErrInvalidRange
	}
	if r.Start.Equal(r.End) {
		return []Range{r}, nil
	}

	result := make([]Range, 0)
	for start := r.Start; start.Before(r.End); {
		end := start.Add(step)
		if end.After(r.End) {
			end = r.End
		}
		result = append(result, Range{Start: start, End: end})
		start = end
	}
	return result, nil
}

// SplitRange 创建闭区间后按固定时长 step 拆分，是 NewRange 和 Range.Split 的便捷组合。
func (t Helper) SplitRange(start, end stdtime.Time, step stdtime.Duration) ([]Range, error) {
	rangeValue, err := t.NewRange(start, end)
	if err != nil {
		return nil, err
	}
	return rangeValue.Split(step)
}
