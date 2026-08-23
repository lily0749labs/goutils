package strto

// ArrStrToBool []string转[]bool
func ArrStrToBool(arrv []string) (r []bool) {
	for _, v := range arrv {
		r = append(r, StrToBool(v))
	}
	return
}
