package anyto

/* any */

func ArrAnyToStr(arrv []any) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}

/* int */
func ArrIntToStr(arrv []int) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}
func ArrInt8ToStr(arrv []int8) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}

func ArrInt16ToStr(arrv []int16) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}
func ArrInt32ToStr(arrv []int32) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}

func ArrInt64ToStr(arrv []int64) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}

/* uint */

func ArrUintToStr(arrv []uint) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}
func ArrUint8ToStr(arrv []uint8) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}

func ArrUint16ToStr(arrv []uint16) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}
func ArrUint32ToStr(arrv []uint32) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}

func ArrUint64ToStr(arrv []uint64) (r []string) {
	for _, v := range arrv {
		r = append(r, AnyToStr(v))
	}
	return
}
