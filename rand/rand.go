package rand

//随机数相关算法
import (
	"bytes"
	cryptorand "crypto/rand"
	"fmt"
	"math/rand"
	"sync"

	"github.com/dromara/carbon/v2"
)

var (
	rr   *rand.Rand
	rrMu sync.Mutex
)

func init() {
	rr = rand.New(rand.NewSource(carbon.Now(carbon.Shanghai).StdTime().UnixNano()))
}

// Rand6 产生 6 位随机数。
func (generator) Rand6() string {
	rrMu.Lock()
	defer rrMu.Unlock()
	code := fmt.Sprintf("%06v", rr.Int31n(1000000))
	return code
}

// Rand4 产生 4 位随机数。
func (generator) Rand4() string {
	rrMu.Lock()
	defer rrMu.Unlock()
	code := fmt.Sprintf("%04v", rr.Int31n(10000))
	return code
}

// Intn 返回区间 [0, n) 内的伪随机整数。
func (generator) Intn(n int) int {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Intn(n)
}

// Int31n 返回区间 [0, n) 内的伪随机 int32。
func (generator) Int31n(n int32) int32 {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Int31n(n)
}

// Int63n 返回区间 [0, n) 内的伪随机 int64。
func (generator) Int63n(n int64) int64 {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Int63n(n)
}

var Chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// var AsciiChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

// NewLenChars 产生指定长度的随机字符串。
func (generator) NewLenChars(length int) string {
	if length <= 0 {
		return ""
	}
	clen := len(Chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := cryptorand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = Chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// UpperString 产生指定长度的大写字母随机字符串。
func (g generator) UpperString(length int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < length; {
		temp = string(rune(g.RandInt(65, 90)))
		result.WriteString(temp)
		i++

	}
	return result.String()
}

// RandInt 返回区间 [min, max) 内的随机整数。
func (g generator) RandInt(min, max int) int {
	if min >= max {
		return max
	}
	return g.Intn(max-min) + min
}

// RandIntM 返回区间 [min, max] 内的随机整数。
func (g generator) RandIntM(min, max int) int {
	if min >= max {
		return max
	}
	max += 1
	return g.Intn(max-min) + min
}

// 传入指定概率，然后返回是否执行  比如 rate：90 表示有90%的概率要执行
func (g generator) RateToExec(rate int) bool {
	if rate <= 0 {
		return false
	}
	if rate >= 100 {
		return true
	}
	return g.RandInt(0, 100) < rate
}

// 传入指定概率，然后返回是否执行  比如 rate：90 表示有90%的概率要执行
func (g generator) RateToExecWan(rate int) bool {
	if rate <= 0 {
		return false
	}
	if rate >= 10000 {
		return true
	}
	return g.RandInt(0, 10000) < rate
}

// 从max中随机去一个数，看是否小于rate
func (g generator) RateToExecWithIn(rate, max int) bool {
	if rate <= 0 || max <= 0 {
		return false
	}
	if rate >= max {
		return true
	}
	return g.RandInt(0, max) < rate
}

// Deprecated: 使用 Rand.Rand6。
func Rand6() string { return Rand.Rand6() }

// Deprecated: 使用 Rand.Rand4。
func Rand4() string { return Rand.Rand4() }

// Deprecated: 使用 Rand.Intn。
func Intn(n int) int { return Rand.Intn(n) }

// Deprecated: 使用 Rand.Int31n。
func Int31n(n int32) int32 { return Rand.Int31n(n) }

// Deprecated: 使用 Rand.Int63n。
func Int63n(n int64) int64 { return Rand.Int63n(n) }

// Deprecated: 使用 Rand.NewLenChars。
func NewLenChars(length int) string { return Rand.NewLenChars(length) }

// RandomString 产生指定长度的大写字母随机字符串。
// Deprecated: 使用 Rand.UpperString。
func RandomString(length int) string { return Rand.UpperString(length) }

// Deprecated: 使用 Rand.RandInt。
func RandInt(min, max int) int { return Rand.RandInt(min, max) }

// Deprecated: 使用 Rand.RandIntM。
func RandIntM(min, max int) int { return Rand.RandIntM(min, max) }

// Deprecated: 使用 Rand.RateToExec。
func RateToExec(rate int) bool { return Rand.RateToExec(rate) }

// Deprecated: 使用 Rand.RateToExecWan。
func RateToExecWan(rate int) bool { return Rand.RateToExecWan(rate) }

// Deprecated: 使用 Rand.RateToExecWithIn。
func RateToExecWithIn(rate, max int) bool { return Rand.RateToExecWithIn(rate, max) }
