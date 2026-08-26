package strto

// Bytes 将字符串转换为字节切片。
func (facade) Bytes(v string) []byte {
	return []byte(v)
}

// StrToBytes 将字符串转换为字节切片。
// Deprecated: 使用 StrTo.Bytes。
func StrToBytes(v string) []byte { return StrTo.Bytes(v) }
