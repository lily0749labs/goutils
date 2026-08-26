package syncutil

import "sync"

// Sync 提供同步容器辅助函数的结构化入口。
var Sync = syncutil{}

type syncutil struct{}

// MapLen 返回 sync.Map 当前遍历到的元素数量。
// sync.Map 不提供一致性快照；并发写入时结果代表遍历过程中的观测值。
func (syncutil) MapLen(values *sync.Map) int {
	if values == nil {
		return 0
	}
	length := 0
	values.Range(func(_, _ any) bool {
		length++
		return true
	})
	return length
}

// MapLen 是 Sync.MapLen 的包级便捷入口。
func MapLen(values *sync.Map) int {
	return Sync.MapLen(values)
}
