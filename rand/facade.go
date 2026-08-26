package rand

// Rand6 generates a six-digit random string.
func (generator) Rand6() string { return Rand6() }

// Rand4 generates a four-digit random string.
func (generator) Rand4() string { return Rand4() }

func (generator) Intn(n int) int       { return Intn(n) }
func (generator) Int31n(n int32) int32 { return Int31n(n) }
func (generator) Int63n(n int64) int64 { return Int63n(n) }

func (generator) NewLenChars(length int) string { return NewLenChars(length) }
func (generator) RandInt(min, max int) int      { return RandInt(min, max) }
func (generator) RandIntM(min, max int) int     { return RandIntM(min, max) }

func (generator) RateToExec(rate int) bool { return RateToExec(rate) }
func (generator) RateToExecWan(rate int) bool {
	return RateToExecWan(rate)
}
func (generator) RateToExecWithIn(rate, max int) bool {
	return RateToExecWithIn(rate, max)
}

func (generator) RandEnName() string      { return RandEnName() }
func (generator) RandEnMenName() string   { return RandEnMenName() }
func (generator) RandEnWomenName() string { return RandEnWomenName() }
