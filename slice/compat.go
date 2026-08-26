package slice

import (
	"cmp"
	"slices"
)

// DeleteFirst 删除 target 首次出现的位置并返回新切片。
// target 不存在时原样返回 values。
func DeleteFirst[T comparable](values []T, target T) []T {
	index := Index(values, target)
	if index < 0 {
		return values
	}
	return DeleteAt(values, index)
}

// UniqueSorted 去重后按升序返回结果；nil 或空切片返回 nil。
func UniqueSorted[T cmp.Ordered](values []T) []T {
	if len(values) == 0 {
		return nil
	}
	result := Unique(values)
	slices.Sort(result)
	return result
}
