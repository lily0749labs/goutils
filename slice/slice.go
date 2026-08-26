package slice

// Clone 复制切片，返回与输入互不共享底层数组的结果。
func Clone[T any](values []T) []T {
	if values == nil {
		return nil
	}
	return append([]T(nil), values...)
}

// Map 将每个元素映射为另一种类型，nil mapper 返回 nil。
func Map[T, R any](values []T, mapper func(T) R) []R {
	if values == nil || mapper == nil {
		return nil
	}
	result := make([]R, len(values))
	for index, value := range values {
		result[index] = mapper(value)
	}
	return result
}

// Filter 返回所有满足 predicate 的元素，保持原始顺序且不修改输入。
func Filter[T any](values []T, predicate func(T) bool) []T {
	if values == nil || predicate == nil {
		return nil
	}
	result := make([]T, 0, len(values))
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce 从 initial 开始依次聚合元素，nil reducer 返回 initial。
func Reduce[T, R any](values []T, initial R, reducer func(R, T) R) R {
	if reducer == nil {
		return initial
	}
	result := initial
	for _, value := range values {
		result = reducer(result, value)
	}
	return result
}

// Contains 判断切片是否包含 target。
func Contains[T comparable](values []T, target T) bool {
	return Index(values, target) >= 0
}

// Index 返回 target 首次出现的下标，不存在时返回 -1。
func Index[T comparable](values []T, target T) int {
	for index, value := range values {
		if value == target {
			return index
		}
	}
	return -1
}

// DeleteAt 删除 index 位置的元素并返回新切片。
// 输入为 nil 时返回 nil；index 越界时返回输入的副本。
func DeleteAt[T any](values []T, index int) []T {
	if values == nil {
		return nil
	}
	if index < 0 || index >= len(values) {
		return Clone(values)
	}

	result := make([]T, len(values)-1)
	copy(result, values[:index])
	copy(result[index:], values[index+1:])
	return result
}

// Unique 去除重复元素并保留首次出现顺序。
func Unique[T comparable](values []T) []T {
	if values == nil {
		return nil
	}
	result := make([]T, 0, len(values))
	seen := make(map[T]struct{}, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

// Chunk 按 size 拆分切片，每个分块都拥有独立底层数组；size 非正数时返回 nil。
func Chunk[T any](values []T, size int) [][]T {
	if values == nil || size <= 0 {
		return nil
	}
	if len(values) == 0 {
		return [][]T{}
	}
	result := make([][]T, 0, (len(values)-1)/size+1)
	for start := 0; start < len(values); {
		end := min(start+size, len(values))
		result = append(result, Clone(values[start:end]))
		if end == len(values) {
			break
		}
		start = end
	}
	return result
}

// GroupBy 按 keySelector 返回的键分组，组内元素保持原始顺序。
func GroupBy[T any, K comparable](values []T, keySelector func(T) K) map[K][]T {
	if values == nil || keySelector == nil {
		return nil
	}
	result := make(map[K][]T)
	for _, value := range values {
		key := keySelector(value)
		result[key] = append(result[key], value)
	}
	return result
}

// Intersection 返回两个切片的去重交集，顺序取自 first。
func Intersection[T comparable](first, second []T) []T {
	if first == nil || second == nil {
		return nil
	}
	secondSet := make(map[T]struct{}, len(second))
	for _, value := range second {
		secondSet[value] = struct{}{}
	}
	result := make([]T, 0)
	seen := make(map[T]struct{})
	for _, value := range first {
		if _, exists := secondSet[value]; !exists {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

// Difference 返回只出现在 first 中的去重元素，顺序取自 first。
func Difference[T comparable](first, second []T) []T {
	if first == nil {
		return nil
	}
	secondSet := make(map[T]struct{}, len(second))
	for _, value := range second {
		secondSet[value] = struct{}{}
	}
	result := make([]T, 0)
	seen := make(map[T]struct{})
	for _, value := range first {
		if _, exists := secondSet[value]; exists {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
