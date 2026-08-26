package maps

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// TestGenericMapFunctions 验证键值提取、复制、合并、筛选和转换。
func TestGenericMapFunctions(t *testing.T) {
	t.Parallel()
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	keys := Keys(values)
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"a", "b", "c"}) {
		t.Fatalf("Keys() = %v", keys)
	}
	items := Values(values)
	sort.Ints(items)
	if !reflect.DeepEqual(items, []int{1, 2, 3}) {
		t.Fatalf("Values() = %v", items)
	}
	if got := Merge(values, map[string]int{"b": 20, "d": 4}); !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 20, "c": 3, "d": 4}) {
		t.Fatalf("Merge() = %v", got)
	}
	if got := Filter(values, func(_ string, value int) bool { return value%2 == 1 }); !reflect.DeepEqual(got, map[string]int{"a": 1, "c": 3}) {
		t.Fatalf("Filter() = %v", got)
	}
	if got := MapValues(values, func(key string, value int) string { return key + strconv.Itoa(value) }); !reflect.DeepEqual(got, map[string]string{"a": "a1", "b": "b2", "c": "c3"}) {
		t.Fatalf("MapValues() = %v", got)
	}
	if got := Pick(values, "a", "missing"); !reflect.DeepEqual(got, map[string]int{"a": 1}) {
		t.Fatalf("Pick() = %v", got)
	}
	if got := Omit(values, "b"); !reflect.DeepEqual(got, map[string]int{"a": 1, "c": 3}) {
		t.Fatalf("Omit() = %v", got)
	}
	if GetOr(values, "a", 9) != 1 || GetOr(values, "missing", 9) != 9 {
		t.Fatal("GetOr() 结果不正确")
	}
	if !ContainsKeys(values, "a", "c") || ContainsKeys(values, "a", "missing") {
		t.Fatal("ContainsKeys() 结果不正确")
	}
	if got := MissingKeys(values, "missing-2", "a", "missing-1", "missing-2"); !reflect.DeepEqual(got, []string{"missing-2", "missing-1"}) {
		t.Fatalf("MissingKeys() = %v", got)
	}
}

// TestMapFunctionsDoNotMutateInput 验证返回值与输入 Map 相互独立。
func TestMapFunctionsDoNotMutateInput(t *testing.T) {
	t.Parallel()
	input := map[string]int{"a": 1}
	cloned := Clone(input)
	cloned["a"] = 2
	merged := Merge(input, map[string]int{"b": 2})
	merged["a"] = 3
	if input["a"] != 1 || len(input) != 1 {
		t.Fatalf("输入 Map 被修改：%v", input)
	}
}

// TestMapNilBehavior 验证 nil 输入保持 nil 语义。
func TestMapNilBehavior(t *testing.T) {
	t.Parallel()
	if Keys[string, int](nil) != nil || Values[string, int](nil) != nil || Clone[string, int](nil) != nil {
		t.Fatal("nil Map 应返回 nil")
	}
	if Merge[string, int]() != nil || Filter[string, int](nil, nil) != nil || MapValues[string, int, string](nil, nil) != nil {
		t.Fatal("nil 输入应返回 nil")
	}
	if !ContainsKeys[string, int](nil) || ContainsKeys[string, int](nil, "missing") {
		t.Fatal("nil Map 的 ContainsKeys() 结果不正确")
	}
	if MissingKeys[string, int](nil) != nil {
		t.Fatal("无待检查键时 MissingKeys() 应返回 nil")
	}
	if got := MissingKeys[string, int](nil, "a", "a", "b"); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("nil Map 的 MissingKeys() = %v", got)
	}
	valuesWithNil := map[string]any{"present": nil}
	if !ContainsKeys(valuesWithNil, "present") || len(MissingKeys(valuesWithNil, "present")) != 0 {
		t.Fatal("nil 值不应被视为键缺失")
	}
}
