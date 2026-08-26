package rand

// 随机数相关算法。
import (
	"bytes"
	"fmt"
	mathrand "math/rand"
	"sync"

	"github.com/dromara/carbon/v2"
)

var (
	rr   *mathrand.Rand
	rrMu sync.Mutex
)

func init() {
	rr = mathrand.New(mathrand.NewSource(carbon.Now(carbon.Shanghai).StdTime().UnixNano()))
}

// Rand6 产生 6 位随机数。
func (rand) Rand6() string {
	rrMu.Lock()
	defer rrMu.Unlock()
	code := fmt.Sprintf("%06v", rr.Int31n(1000000))
	return code
}

// Rand4 产生 4 位随机数。
func (rand) Rand4() string {
	rrMu.Lock()
	defer rrMu.Unlock()
	code := fmt.Sprintf("%04v", rr.Int31n(10000))
	return code
}

// Intn 返回区间 [0, n) 内的伪随机整数。
func (rand) Intn(n int) int {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Intn(n)
}

// Int31n 返回区间 [0, n) 内的伪随机 int32。
func (rand) Int31n(n int32) int32 {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Int31n(n)
}

// Int63n 返回区间 [0, n) 内的伪随机 int64。
func (rand) Int63n(n int64) int64 {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Int63n(n)
}

var Chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// NewLenCharsE 产生指定长度的安全随机字符串，并报告非法参数或随机源错误。
func (r rand) NewLenCharsE(length int) (string, error) {
	return r.SecureStringWithCharset(length, string(Chars))
}

// NewLenChars 产生指定长度的随机字符串。
// 失败时返回空字符串；需要错误信息时请使用 NewLenCharsE。
func (r rand) NewLenChars(length int) string {
	value, _ := r.NewLenCharsE(length)
	return value
}

// UpperString 产生指定长度的大写字母随机字符串。
func (g rand) UpperString(length int) string {
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
func (g rand) RandInt(min, max int) int {
	if min >= max {
		return max
	}
	return g.Intn(max-min) + min
}

// RandIntM 返回区间 [min, max] 内的随机整数。
func (g rand) RandIntM(min, max int) int {
	if min >= max {
		return max
	}
	max += 1
	return g.Intn(max-min) + min
}

// RateToExec 按百分比概率返回是否执行，rate 为 90 表示 90% 的概率。
func (g rand) RateToExec(rate int) bool {
	if rate <= 0 {
		return false
	}
	if rate >= 100 {
		return true
	}
	return g.RandInt(0, 100) < rate
}

// RateToExecWan 按万分比概率返回是否执行，rate 为 9000 表示 90% 的概率。
func (g rand) RateToExecWan(rate int) bool {
	if rate <= 0 {
		return false
	}
	if rate >= 10000 {
		return true
	}
	return g.RandInt(0, 10000) < rate
}

// RateToExecWithIn 在 [0, max) 中取随机数，并判断其是否小于 rate。
func (g rand) RateToExecWithIn(rate, max int) bool {
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
