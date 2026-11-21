// Package dbtypes 放数据库列的通用 Go 类型，供 gorm-rueidis 的 dataMap 映射和手写模型使用。
package dbtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Element 是数组元素允许的 Go 类型；uint64 用 numeric(20,0)[] 精确保存完整编号。
// 元素不能为 NULL（Go 的切片装不下 NULL，读到会报错）。
type Element interface {
	~string | ~bool | ~int16 | ~int32 | ~int64 | ~uint64 | ~float32 | ~float64
}

// Array 是 PostgreSQL 一维数组列（text[]、smallint[]、integer[]、bigint[]、boolean[]、real[]、double precision[]）
// 在 Go 里的通用类型，一个泛型替代 lib/pq 的 StringArray、Int32Array、Int64Array 等一组类型：
//
//	Tags   dbtypes.Array[string] // text[]
//	Scores dbtypes.Array[int32]  // integer[]
//
// 它实现 sql.Scanner 和 driver.Valuer，按 PostgreSQL 的数组文本格式（{a,"b c",d}）读写，
// 经 gorm 的 pgx 驱动（database/sql）直接可用。nil 切片写成 NULL，空切片写成 {}。
//
// 限制：只支持一维数组；元素不能为 NULL；不支持多维数组和下标带范围的写法之外的其他数组修饰。
type Array[T Element] []T

// GormDataType 声明 PostgreSQL 数组类型，让未携带列类型标签的更新 DTO 也能被 GORM 解析。
// 仅依赖 Value 时，nil 数组返回 nil，GORM 无法据此判断切片元素的数据库类型。
func (Array[T]) GormDataType() string {
	switch reflect.TypeFor[T]().Kind() {
	case reflect.String:
		return "text[]"
	case reflect.Bool:
		return "boolean[]"
	case reflect.Int16:
		return "smallint[]"
	case reflect.Int32:
		return "integer[]"
	case reflect.Int64:
		return "bigint[]"
	case reflect.Uint64:
		return "numeric(20,0)[]"
	case reflect.Float32:
		return "real[]"
	case reflect.Float64:
		return "double precision[]"
	default:
		panic("dbtypes.Array: 不支持的元素类型")
	}
}

// Value 实现 driver.Valuer：把切片写成 PostgreSQL 数组字面量。
func (a Array[T]) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	var b strings.Builder
	b.WriteByte('{')
	for i, e := range a {
		if i > 0 {
			b.WriteByte(',')
		}
		if err := writeElement(&b, reflect.ValueOf(e)); err != nil {
			return nil, err
		}
	}
	b.WriteByte('}')
	return b.String(), nil
}

// Scan 实现 sql.Scanner：读取 PostgreSQL 数组的文本格式。来源可以是 []byte、string 或 nil（NULL，得到 nil 切片）。
func (a *Array[T]) Scan(src any) error {
	var text string
	switch v := src.(type) {
	case nil:
		*a = nil
		return nil
	case []byte:
		text = string(v)
	case string:
		text = v
	default:
		return fmt.Errorf("dbtypes.Array: 不支持从 %T 读取", src)
	}
	items, err := parseArray(text)
	if err != nil {
		return err
	}
	out := make(Array[T], len(items))
	for i, item := range items {
		if item.null {
			return fmt.Errorf("dbtypes.Array: 第 %d 个元素是 NULL，Go 的切片装不下 NULL；请在数据库里禁止空元素", i+1)
		}
		if err := setElement(reflect.ValueOf(&out[i]).Elem(), item.text); err != nil {
			return fmt.Errorf("dbtypes.Array: 第 %d 个元素 %q: %w", i+1, item.text, err)
		}
	}
	*a = out
	return nil
}

// writeElement 把一个元素写成数组字面量里的一项：字符串一律加双引号并转义，其余原样。
func writeElement(b *strings.Builder, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		b.WriteByte('"')
		for _, r := range v.String() {
			if r == '"' || r == '\\' {
				b.WriteByte('\\')
			}
			b.WriteRune(r)
		}
		b.WriteByte('"')
	case reflect.Bool:
		b.WriteString(strconv.FormatBool(v.Bool()))
	case reflect.Int16, reflect.Int32, reflect.Int64:
		b.WriteString(strconv.FormatInt(v.Int(), 10))
	case reflect.Uint64:
		b.WriteString(strconv.FormatUint(v.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		switch {
		case math.IsNaN(f):
			b.WriteString("NaN")
		case math.IsInf(f, 1):
			b.WriteString("Infinity")
		case math.IsInf(f, -1):
			b.WriteString("-Infinity")
		default:
			bits := 64
			if v.Kind() == reflect.Float32 {
				bits = 32
			}
			b.WriteString(strconv.FormatFloat(f, 'g', -1, bits))
		}
	default:
		return fmt.Errorf("dbtypes.Array: 不支持的元素类型 %s", v.Type())
	}
	return nil
}

// setElement 把数组里取出的一项文本转成元素类型。
func setElement(dst reflect.Value, text string) error {
	switch dst.Kind() {
	case reflect.String:
		dst.SetString(text)
	case reflect.Bool:
		// PostgreSQL 输出 t、f；字面量里也可能是 true、false 等。
		switch strings.ToLower(text) {
		case "t", "true", "y", "yes", "on", "1":
			dst.SetBool(true)
		case "f", "false", "n", "no", "off", "0":
			dst.SetBool(false)
		default:
			return errors.New("不是布尔值")
		}
	case reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(text, 10, dst.Type().Bits())
		if err != nil {
			return err
		}
		dst.SetInt(n)
	case reflect.Uint64:
		n, err := strconv.ParseUint(text, 10, 64)
		if err != nil {
			return err
		}
		dst.SetUint(n)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(text, dst.Type().Bits())
		if err != nil {
			return err
		}
		dst.SetFloat(f)
	default:
		return fmt.Errorf("不支持的元素类型 %s", dst.Type())
	}
	return nil
}

// arrayItem 是解析出的一个数组元素。
type arrayItem struct {
	text string
	null bool // 没加引号的 NULL（不区分大小写）；加了引号的 "NULL" 是字符串
}

// parseArray 解析 PostgreSQL 一维数组的文本格式：{a,"b c","d\"e",NULL}，也接受带下标范围的前缀 [1:3]={...}。
func parseArray(s string) ([]arrayItem, error) {
	s = strings.TrimSpace(s)
	// 下标范围前缀只在数组下标不是从 1 开始时才会出现，忽略它。
	if strings.HasPrefix(s, "[") {
		eq := strings.Index(s, "=")
		if eq < 0 {
			return nil, errors.New("dbtypes.Array: 数组下标前缀缺少 =")
		}
		s = strings.TrimSpace(s[eq+1:])
	}
	if len(s) < 2 || s[0] != '{' || s[len(s)-1] != '}' {
		return nil, fmt.Errorf("dbtypes.Array: 不是数组格式: %q", s)
	}
	body := s[1 : len(s)-1]
	if strings.TrimSpace(body) == "" {
		return []arrayItem{}, nil
	}
	var items []arrayItem
	i := 0
	for {
		for i < len(body) && body[i] == ' ' {
			i++
		}
		if i >= len(body) {
			return nil, fmt.Errorf("dbtypes.Array: 数组在逗号后意外结束: %q", s)
		}
		switch body[i] {
		case '{':
			return nil, errors.New("dbtypes.Array: 不支持多维数组")
		case '"':
			i++
			var b strings.Builder
			closed := false
			for i < len(body) {
				c := body[i]
				if c == '\\' {
					i++
					if i >= len(body) {
						return nil, fmt.Errorf("dbtypes.Array: 转义符后意外结束: %q", s)
					}
					b.WriteByte(body[i])
				} else if c == '"' {
					closed = true
					i++
					break
				} else {
					b.WriteByte(c)
				}
				i++
			}
			if !closed {
				return nil, fmt.Errorf("dbtypes.Array: 引号没有闭合: %q", s)
			}
			items = append(items, arrayItem{text: b.String()})
		default:
			start := i
			for i < len(body) && body[i] != ',' {
				i++
			}
			raw := strings.TrimSpace(body[start:i])
			if raw == "" {
				return nil, fmt.Errorf("dbtypes.Array: 出现空元素: %q", s)
			}
			items = append(items, arrayItem{text: raw, null: strings.EqualFold(raw, "NULL")})
		}
		for i < len(body) && body[i] == ' ' {
			i++
		}
		if i >= len(body) {
			return items, nil
		}
		if body[i] != ',' {
			return nil, fmt.Errorf("dbtypes.Array: 元素之间缺少逗号: %q", s)
		}
		i++
	}
}
