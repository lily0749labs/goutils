package dbtypes

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestUint64ArrayPreservesFullRange(t *testing.T) {
	want := Array[uint64]{math.MaxInt32 + 1, math.MaxInt64 + 1, math.MaxUint64}
	raw, err := want.Value()
	if err != nil || raw != "{2147483648,9223372036854775808,18446744073709551615}" {
		t.Fatal("无符号编号写出时被截断", raw, err)
	}
	var got Array[uint64]
	if err := got.Scan(raw); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatal("无符号编号读取时被截断", got, err)
	}
	for _, invalid := range []string{"{-1}", "{18446744073709551616}", "{1.5}", "{NULL}"} {
		if err := got.Scan(invalid); err == nil {
			t.Fatal("非法数组元素未被拒绝", invalid)
		}
	}
	if got.GormDataType() != "numeric(20,0)[]" {
		t.Fatal("编号数组仍使用有符号 SQL 类型")
	}
}

func TestValue(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  func() (any, error)
		want any
	}{

		{"nil 写成 NULL", func() (any, error) { return Array[string](nil).Value() }, nil},

		{"空切片写成 {}", func() (any, error) { return Array[string]{}.Value() }, "{}"},

		{
			"字符串一律加引号",
			func() (
				any,
				error,
			) {
				return Array[string]{
					"a",
					"b c",
					"d,e",
				}.Value()
			},
			`{"a","b c","d,e"}`,
		},

		{
			"引号和反斜杠转义",
			func() (
				any,
				error,
			) {
				return Array[string]{
					`say "hi"`,
					`a\b`,
				}.Value()
			},
			`{"say \"hi\"","a\\b"}`,
		},

		{
			"字符串 NULL 加引号才是字符串",
			func() (
				any,
				error,
			) {
				return Array[string]{"NULL"}.Value()
			},
			`{"NULL"}`,
		},

		{"空字符串", func() (any, error) { return Array[string]{""}.Value() }, `{""}`},

		{
			"中文和花括号",
			func() (
				any,
				error,
			) {
				return Array[string]{
					"你好",
					"{x}",
				}.Value()
			},
			`{"你好","{x}"}`,
		},

		{"int32", func() (any, error) { return Array[int32]{1, -2, 3}.Value() }, `{1,-2,3}`},

		{
			"int64",
			func() (any, error) { return Array[int64]{math.MaxInt64, math.MinInt64}.Value() },
			`{9223372036854775807,-9223372036854775808}`,
		},

		{"bool", func() (any, error) { return Array[bool]{true, false}.Value() }, `{true,false}`},

		{"float64 含特殊值", func() (any, error) {
			return Array[float64]{1.5, math.NaN(), math.Inf(1), math.Inf(-1)}.Value()
		}, `{1.5,NaN,Infinity,-Infinity}`},

		{"float32 不带多余位数", func() (any, error) { return Array[float32]{0.1}.Value() }, `{0.1}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.got()
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Fatalf("got %#v want %#v", got, tc.want)
			}
		})
	}
}

func TestScan(t *testing.T) {
	var s Array[string]
	for _, tc := range []struct {
		name string
		src  any
		want Array[string]
	}{
		{"nil 得到 nil", nil, nil},
		{"空数组", `{}`, Array[string]{}},
		{"[]byte 来源", []byte(`{a,b}`), Array[string]{"a", "b"}},
		{"不带引号", `{a,b,c}`, Array[string]{"a", "b", "c"}},
		{"带引号含空格和逗号", `{"a b","c,d"}`, Array[string]{"a b", "c,d"}},
		{"转义", `{"say \"hi\"","a\\b"}`, Array[string]{`say "hi"`, `a\b`}},
		{"加引号的 NULL 是字符串", `{"NULL","null"}`, Array[string]{"NULL", "null"}},
		{"空字符串", `{""}`, Array[string]{""}},
		{"中文", `{"你好",世界}`, Array[string]{"你好", "世界"}},
		{"元素周围的空格", `{ a , "b" }`, Array[string]{"a", "b"}},
		{"下标前缀", `[0:1]={a,b}`, Array[string]{"a", "b"}},
		{"外层空白", "  {a}\n", Array[string]{"a"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s = Array[string]{"stale"}
			if err := s.Scan(tc.src); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(s, tc.want) {
				t.Fatalf("got %#v want %#v", s, tc.want)
			}
		})
	}
}

func TestScanTypes(t *testing.T) {
	var i32 Array[int32]
	if err := i32.Scan(`{1,-2,2147483647}`); err != nil || !reflect.DeepEqual(i32, Array[int32]{1, -2, math.MaxInt32}) {
		t.Fatalf("int32: %v %v", i32, err)
	}
	var i16 Array[int16]
	if err := i16.Scan(
		`{32767,-32768}`,
	); err != nil ||
		!reflect.DeepEqual(i16, Array[int16]{math.MaxInt16, math.MinInt16}) {
		t.Fatalf("int16: %v %v", i16, err)
	}
	var i64 Array[int64]
	if err := i64.Scan(`{9223372036854775807}`); err != nil || i64[0] != math.MaxInt64 {
		t.Fatalf("int64: %v %v", i64, err)
	}
	var b Array[bool]
	if err := b.Scan(`{t,f,true,false}`); err != nil || !reflect.DeepEqual(b, Array[bool]{true, false, true, false}) {
		t.Fatalf("bool: %v %v", b, err)
	}
	var f Array[float64]
	if err := f.Scan(
		`{1.5,NaN,Infinity,-Infinity,1e3}`,
	); err != nil || f[0] != 1.5 || !math.IsNaN(f[1]) || !math.IsInf(f[2], 1) || !math.IsInf(f[3], -1) ||
		f[4] != 1000 {
		t.Fatalf("float64: %v %v", f, err)
	}
	// 命名类型（~string）也能用。
	type Tag string
	var tags Array[Tag]
	if err := tags.Scan(`{x,y}`); err != nil || !reflect.DeepEqual(tags, Array[Tag]{"x", "y"}) {
		t.Fatalf("named: %v %v", tags, err)
	}
}

func TestScanErrors(t *testing.T) {
	var s Array[string]
	for name, src := range map[string]any{
		"元素是 NULL":   `{a,NULL}`,
		"NULL 不分大小写": `{null}`,
		"多维数组":       `{{a},{b}}`,
		"没有花括号":      `a,b`,
		"右括号后有多余内容":  `{a}b`,
		"引号没闭合":      `{"abc}`,
		"转义符后结束":     `{"abc\`,
		"逗号后结束":      `{a,}`,
		"空元素":        `{,a}`,
		"元素之间缺逗号":    `{"a" "b"}`,
		"下标前缀缺等号":    `[0:1]{a}`,
		"来源类型不支持":    12345,
		"空串":         ``,
	} {
		if err := s.Scan(src); err == nil {
			t.Errorf("%s: %v 应当报错，得到 %#v", name, src, s)
		}
	}
	var i16 Array[int16]
	if err := i16.Scan(`{40000}`); err == nil {
		t.Error("int16 溢出应当报错")
	}
	var i Array[int32]
	if err := i.Scan(`{abc}`); err == nil {
		t.Error("非数字应当报错")
	}
	var b Array[bool]
	if err := b.Scan(`{maybe}`); err == nil {
		t.Error("不是布尔值应当报错")
	}
	// NULL 元素的错误信息要说清楚是第几个。
	if err := s.Scan(`{a,NULL}`); err == nil || !strings.Contains(err.Error(), "第 2 个元素是 NULL") {
		t.Errorf("错误信息 = %v", err)
	}
}

// 写出去再读回来必须得到原值，包括各种需要转义的字符。
func TestRoundTrip(t *testing.T) {
	cases := Array[string]{
		"",
		" ",
		"a",
		"a b",
		"a,b",
		`a"b`,
		`a\b`,
		`\`,
		`"`,
		`\"`,
		"{",
		"}",
		"{a,b}",
		"NULL",
		"null",
		"你好 世界",
		"tab\there",
		"line\nbreak",
		"'single'",
		"a=b",
		"[0:1]",
	}
	v, err := cases.Value()
	if err != nil {
		t.Fatal(err)
	}
	var back Array[string]
	if err := back.Scan(v); err != nil {
		t.Fatalf("%v\n%v", v, err)
	}
	if !reflect.DeepEqual(back, cases) {
		t.Fatalf("round trip 不一致\n got  %#v\n want %#v", back, cases)
	}
	for _, f := range []Array[float64]{
		{0, 1.5, -2.25, 1e-9, 1e21, math.MaxFloat64, math.SmallestNonzeroFloat64},
		{math.Inf(1), math.Inf(-1)},
	} {
		v, _ := f.Value()
		var got Array[float64]
		if err := got.Scan(v); err != nil || !reflect.DeepEqual(got, f) {
			t.Fatalf("float round trip: %v %v %v", v, got, err)
		}
	}
}
