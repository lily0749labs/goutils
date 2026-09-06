package strto

import stdtime "time"

// Float32s 将字符串切片宽松转换为 float32 切片。
func (s strto) Float32s(values []string) (result []float32) {
	for _, value := range values {
		result = append(result, s.Float32(value))
	}
	return result
}

// Float64s 将字符串切片宽松转换为 float64 切片。
func (s strto) Float64s(values []string) (result []float64) {
	for _, value := range values {
		result = append(result, s.Float64(value))
	}
	return result
}

// Durations 将字符串切片宽松转换为 time.Duration 切片。
func (s strto) Durations(values []string) (result []stdtime.Duration) {
	for _, value := range values {
		result = append(result, s.Duration(value))
	}
	return result
}
