package strto

// ArrStrToInt []string转[]int
func ArrStrToInt(arrv []string) (r []int) {
	for _, v := range arrv {
		r = append(r, StrToInt(v))
	}
	return
}

// ArrStrToInt8 []string转[]int8
func ArrStrToInt8(arrv []string) (r []int8) {
	for _, v := range arrv {
		r = append(r, StrToInt8(v))
	}
	return
}

// ArrStrToInt16 []string转[]int16
func ArrStrToInt16(arrv []string) (r []int16) {
	for _, v := range arrv {
		r = append(r, StrToInt16(v))
	}
	return
}

// ArrStrToInt32 []string转[]int32
func ArrStrToInt32(arrv []string) (r []int32) {

	for _, v := range arrv {
		r = append(r, StrToInt32(v))
	}
	return
}

// ArrStrToInt64 []string转[]int64
func ArrStrToInt64(arrv []string) (r []int64) {
	for _, v := range arrv {
		r = append(r, StrToInt64(v))
	}
	return
}
