package strto

// Bools 将字符串切片转换为 bool 切片。
func (f facade) Bools(values []string) (result []bool) {
	for _, value := range values {
		result = append(result, f.Bool(value))
	}
	return result
}

// ArrStrToBool 将字符串切片转换为 bool 切片。
// Deprecated: 使用 StrTo.Bools。
func ArrStrToBool(values []string) []bool { return StrTo.Bools(values) }
