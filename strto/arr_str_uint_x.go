package strto

// ArrStrToUint []string转[]uint
func ArrStrToUint(arrv []string) (r []uint) {
	for _, v := range arrv {
		r = append(r, StrToUint(v))
	}
	return
}

// ArrStrToUint8 []string转[]uint8
func ArrStrToUint8(arrv []string) (r []uint8) {
	for _, v := range arrv {
		r = append(r, StrToUint8(v))
	}
	return
}

// ArrStrToUint16 []string转[]uint16
func ArrStrToUint16(arrv []string) (r []uint16) {
	for _, v := range arrv {
		r = append(r, StrToUint16(v))
	}
	return
}

// ArrStrToUint32 []string转[]uint32
func ArrStrToUint32(arrv []string) (r []uint32) {

	for _, v := range arrv {
		r = append(r, StrToUint32(v))
	}
	return
}

// ArrStrToUint64 []string转[]uint64
func ArrStrToUint64(arrv []string) (r []uint64) {
	for _, v := range arrv {
		r = append(r, StrToUint64(v))
	}
	return
}
