package slice

// Collection 提供同类型切片的链式处理能力。
// 需要改变元素类型时请使用包级 Map 函数。
type Collection[T comparable] struct {
	values []T
}

// New 创建一个不共享输入底层数组的集合。
func New[T comparable](values []T) Collection[T] {
	return Collection[T]{values: Clone(values)}
}

// Values 返回集合内容的副本。
func (c Collection[T]) Values() []T { return Clone(c.values) }

// Map 对集合执行同类型映射。
func (c Collection[T]) Map(mapper func(T) T) Collection[T] {
	return New(Map(c.values, mapper))
}

// Filter 筛选集合元素。
func (c Collection[T]) Filter(predicate func(T) bool) Collection[T] {
	return New(Filter(c.values, predicate))
}

// Unique 去除集合中的重复元素。
func (c Collection[T]) Unique() Collection[T] {
	return New(Unique(c.values))
}

// Contains 判断集合是否包含 target。
func (c Collection[T]) Contains(target T) bool { return Contains(c.values, target) }

// DeleteAt 删除 index 位置的元素。
func (c Collection[T]) DeleteAt(index int) Collection[T] {
	return New(DeleteAt(c.values, index))
}

// Intersection 返回与 other 的交集。
func (c Collection[T]) Intersection(other []T) Collection[T] {
	return New(Intersection(c.values, other))
}

// Difference 返回与 other 的差集。
func (c Collection[T]) Difference(other []T) Collection[T] {
	return New(Difference(c.values, other))
}

// Chunk 将集合拆分为多个独立切片。
func (c Collection[T]) Chunk(size int) [][]T { return Chunk(c.values, size) }
