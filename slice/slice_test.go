package slice

import (
	"reflect"
	"strconv"
	"testing"
)

// TestGenericSliceFunctions 验证泛型映射、筛选、聚合、分组和集合运算。
func TestGenericSliceFunctions(t *testing.T) {
	t.Parallel()
	values := []int{1, 2, 2, 3, 4}

	if got := Map(values, strconv.Itoa); !reflect.DeepEqual(got, []string{"1", "2", "2", "3", "4"}) {
		t.Fatalf("Map() = %v", got)
	}
	if got := Filter(values, func(value int) bool { return value%2 == 0 }); !reflect.DeepEqual(got, []int{2, 2, 4}) {
		t.Fatalf("Filter() = %v", got)
	}
	if got := Reduce(values, 0, func(total, value int) int { return total + value }); got != 12 {
		t.Fatalf("Reduce() = %d，期望 12", got)
	}
	if got := Unique(values); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Fatalf("Unique() = %v", got)
	}
	if !Contains(values, 3) || Contains(values, 9) || Index(values, 2) != 1 {
		t.Fatal("Contains() 或 Index() 结果不正确")
	}
	if got := DeleteAt(values, 2); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Fatalf("DeleteAt() = %v", got)
	}
	if got := Chunk(values, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {2, 3}, {4}}) {
		t.Fatalf("Chunk() = %v", got)
	}
	if got := GroupBy(values, func(value int) string { return []string{"偶数", "奇数"}[value%2] }); !reflect.DeepEqual(got, map[string][]int{"奇数": {1, 3}, "偶数": {2, 2, 4}}) {
		t.Fatalf("GroupBy() = %v", got)
	}
	if got := Intersection(values, []int{2, 3, 5}); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Fatalf("Intersection() = %v", got)
	}
	if got := Difference(values, []int{2, 5}); !reflect.DeepEqual(got, []int{1, 3, 4}) {
		t.Fatalf("Difference() = %v", got)
	}
}

// TestCollectionChainingAndCopies 验证链式集合不会修改或泄露输入底层数组。
func TestCollectionChainingAndCopies(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 2, 3}
	collection := New(input).Map(func(value int) int { return value * 2 }).Filter(func(value int) bool { return value >= 4 }).Unique()
	if got := collection.Values(); !reflect.DeepEqual(got, []int{4, 6}) {
		t.Fatalf("collection.Values() = %v", got)
	}
	if !collection.Contains(4) || collection.Contains(2) {
		t.Fatal("collection.Contains() 结果不正确")
	}
	input[0] = 99
	values := collection.Values()
	values[0] = 88
	if got := collection.Values(); !reflect.DeepEqual(got, []int{4, 6}) {
		t.Fatalf("集合内容被外部修改：%v", got)
	}
	if got := collection.Intersection([]int{6}).Values(); !reflect.DeepEqual(got, []int{6}) {
		t.Fatalf("Collection.Intersection() = %v", got)
	}
	if got := collection.Difference([]int{6}).Values(); !reflect.DeepEqual(got, []int{4}) {
		t.Fatalf("Collection.Difference() = %v", got)
	}
	if got := New([]int{1, 2, 3}).DeleteAt(1).Values(); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Fatalf("Collection.DeleteAt() = %v", got)
	}
}

// TestSliceBoundaryBehavior 验证 nil、空回调和非法分块大小不会触发 panic。
func TestSliceBoundaryBehavior(t *testing.T) {
	t.Parallel()
	if Clone[int](nil) != nil || Map[int, int](nil, nil) != nil || Filter[int](nil, nil) != nil {
		t.Fatal("nil 输入应返回 nil")
	}
	if got := Reduce([]int{1}, 7, nil); got != 7 {
		t.Fatalf("Reduce(nil reducer) = %d，期望 7", got)
	}
	if Chunk([]int{1}, 0) != nil || GroupBy[int, int](nil, nil) != nil {
		t.Fatal("非法边界输入应返回 nil")
	}
	if got := Chunk([]int{1}, int(^uint(0)>>1)); !reflect.DeepEqual(got, [][]int{{1}}) {
		t.Fatalf("超大分块尺寸结果 = %v", got)
	}
}

// TestDeleteAtCopies 验证删除和越界操作都不会共享输入的底层数组。
func TestDeleteAtCopies(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3}
	deleted := DeleteAt(input, 1)
	deleted[0] = 99
	if !reflect.DeepEqual(input, []int{1, 2, 3}) {
		t.Fatalf("DeleteAt() 修改了输入：%v", input)
	}

	unchanged := DeleteAt(input, -1)
	unchanged[0] = 88
	if !reflect.DeepEqual(input, []int{1, 2, 3}) {
		t.Fatalf("越界 DeleteAt() 与输入共享底层数组：%v", input)
	}
	if DeleteAt[int](nil, 0) != nil {
		t.Fatal("nil 输入应返回 nil")
	}
}
