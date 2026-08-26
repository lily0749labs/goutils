package maps

import (
	"cmp"
	"slices"
)

// Keys 返回 Map 的全部键。Go Map 无固定迭代顺序，因此结果顺序不稳定。
func Keys[K comparable, V any](values map[K]V) []K {
	if values == nil {
		return nil
	}
	result := make([]K, 0, len(values))
	for key := range values {
		result = append(result, key)
	}
	return result
}

// SortedKeys 返回 Map 的全部键并按升序排列。
func SortedKeys[K cmp.Ordered, V any](values map[K]V) []K {
	result := Keys(values)
	slices.Sort(result)
	return result
}

// Values 返回 Map 的全部值。Go Map 无固定迭代顺序，因此结果顺序不稳定。
func Values[K comparable, V any](values map[K]V) []V {
	if values == nil {
		return nil
	}
	result := make([]V, 0, len(values))
	for _, value := range values {
		result = append(result, value)
	}
	return result
}

// Clone 返回 Map 的浅拷贝。
func Clone[K comparable, V any](values map[K]V) map[K]V {
	if values == nil {
		return nil
	}
	result := make(map[K]V, len(values))
	for key, value := range values {
		result[key] = value
	}
	return result
}

// Merge 合并多个 Map，后传入的 Map 覆盖相同键的旧值。
func Merge[K comparable, V any](values ...map[K]V) map[K]V {
	if values == nil {
		return nil
	}
	capacity := 0
	for _, value := range values {
		capacity += len(value)
	}
	result := make(map[K]V, capacity)
	for _, value := range values {
		for key, item := range value {
			result[key] = item
		}
	}
	return result
}

// Filter 返回满足 predicate 的键值对，nil predicate 返回 nil。
func Filter[K comparable, V any](values map[K]V, predicate func(K, V) bool) map[K]V {
	if values == nil || predicate == nil {
		return nil
	}
	result := make(map[K]V)
	for key, value := range values {
		if predicate(key, value) {
			result[key] = value
		}
	}
	return result
}

// MapValues 转换 Map 的值并保留原始键，nil mapper 返回 nil。
func MapValues[K comparable, V, R any](values map[K]V, mapper func(K, V) R) map[K]R {
	if values == nil || mapper == nil {
		return nil
	}
	result := make(map[K]R, len(values))
	for key, value := range values {
		result[key] = mapper(key, value)
	}
	return result
}

// Pick 返回 keys 指定且实际存在的键值对。
func Pick[K comparable, V any](values map[K]V, keys ...K) map[K]V {
	if values == nil {
		return nil
	}
	result := make(map[K]V, len(keys))
	for _, key := range keys {
		if value, exists := values[key]; exists {
			result[key] = value
		}
	}
	return result
}

// Omit 返回移除 keys 后的 Map 副本。
func Omit[K comparable, V any](values map[K]V, keys ...K) map[K]V {
	if values == nil {
		return nil
	}
	omitted := make(map[K]struct{}, len(keys))
	for _, key := range keys {
		omitted[key] = struct{}{}
	}
	result := make(map[K]V, len(values))
	for key, value := range values {
		if _, exists := omitted[key]; !exists {
			result[key] = value
		}
	}
	return result
}

// ContainsKeys 判断 values 是否包含全部 keys。
// keys 为空时返回 true；Map 中的零值或 nil 值仍视为键已存在。
func ContainsKeys[K comparable, V any](values map[K]V, keys ...K) bool {
	for _, key := range keys {
		if _, exists := values[key]; !exists {
			return false
		}
	}
	return true
}

// MissingKeys 返回 values 中缺少的 keys，保留请求顺序并去除重复键。
// keys 为空时返回 nil。
func MissingKeys[K comparable, V any](values map[K]V, keys ...K) []K {
	if len(keys) == 0 {
		return nil
	}

	missing := make([]K, 0)
	seen := make(map[K]struct{}, len(keys))
	for _, key := range keys {
		if _, checked := seen[key]; checked {
			continue
		}
		seen[key] = struct{}{}
		if _, exists := values[key]; !exists {
			missing = append(missing, key)
		}
	}
	return missing
}

// GetOr 返回 key 对应的值，不存在时返回 fallback。
func GetOr[K comparable, V any](values map[K]V, key K, fallback V) V {
	if value, exists := values[key]; exists {
		return value
	}
	return fallback
}
