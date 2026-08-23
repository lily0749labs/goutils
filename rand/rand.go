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

// 产生6位随机数
func Rand6() string {
	rrMu.Lock()
	defer rrMu.Unlock()
	code := fmt.Sprintf("%06v", rr.Int31n(1000000))
	return code
}

// 产生4位随机数
func Rand4() string {
	rrMu.Lock()
	defer rrMu.Unlock()
	code := fmt.Sprintf("%04v", rr.Int31n(10000))
	return code
}

func Intn(n int) int {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Intn(n)
}

func Int31n(n int32) int32 {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Int31n(n)
}

func Int63n(n int64) int64 {
	rrMu.Lock()
	defer rrMu.Unlock()
	return rr.Int63n(n)
}

var Chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// var AsciiChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

// 产生指定长度随机字符串
func NewLenChars(length int) string {
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

func RandomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		temp = string(rune(RandInt(65, 90)))
		result.WriteString(temp)
		i++

	}
	return result.String()
}

func RandInt(min, max int) int {
	if min >= max {
		return max
	}
	return Intn(max-min) + min
}

func RandIntM(min, max int) int {
	if min >= max {
		return max
	}
	max += 1
	return Intn(max-min) + min
}

// 传入指定概率，然后返回是否执行  比如 rate：90 表示有90%的概率要执行
func RateToExec(rate int) bool {
	if rate <= 0 {
		return false
	}
	if rate >= 100 {
		return true
	}
	return RandInt(0, 100) < rate
}

// 传入指定概率，然后返回是否执行  比如 rate：90 表示有90%的概率要执行
func RateToExecWan(rate int) bool {
	if rate <= 0 {
		return false
	}
	if rate >= 10000 {
		return true
	}
	return RandInt(0, 10000) < rate
}

// 从max中随机去一个数，看是否小于rate
func RateToExecWithIn(rate, max int) bool {
	if rate <= 0 || max <= 0 {
		return false
	}
	if rate >= max {
		return true
	}
	return RandInt(0, max) < rate
}
