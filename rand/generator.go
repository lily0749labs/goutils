package rand

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/lily0749labs/goutils/time"
)

var (
	// Rand 提供无状态的随机辅助方法。
	Rand = generator{}
	// RandPtr 指向 Rand，为兼容旧代码而保留。
	RandPtr = &Rand
)

type generator struct{}

// GetRand 返回一个使用新种子创建的伪随机数生成器。
func (generator) GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Time.NowUnixNano()))
}

// GenerateOrderNumber 返回一个基于时间戳生成的订单号。
func (r generator) GenerateOrderNumber() string {
	rr := r.GetRand()
	rd := rr.Intn(89999) + 10000
	return strconv.Itoa(int(time.Time.NowUnix()*100000) + rd)
}

// GenOrderNo 返回一个基于时间戳生成的订单号。
// Deprecated: 请使用 GenerateOrderNumber。
func (r generator) GenOrderNo() string {
	return r.GenerateOrderNumber()
}

// ***********************************************************
// 生成随机的字符串 包含大写字母/数字
// ***********************************************************

// 设计一个随机种子相加
var (
	randomParam int64
	randomMu    sync.Mutex
)

func nextSeed() int64 {
	randomMu.Lock()
	defer randomMu.Unlock()
	randomParam++
	if randomParam > 1000000000 {
		randomParam = 0
	}
	return randomParam + time.Time.TimestampMilliseconds()
}

// RandomString 返回由数字和大写字母组成的随机字符串。
func (generator) RandomString(size int) string {
	if size <= 0 {
		return ""
	}
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, size)
	r := rand.New(rand.NewSource(nextSeed()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// ***********************************************************
// 生成随机的字符串 只包含小写字母
// ***********************************************************
// RandomStringMix 返回由小写 ASCII 字母组成的随机字符串。
func (generator) RandomStringMix(size int) string {
	if size <= 0 {
		return ""
	}
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, size)
	r := rand.New(rand.NewSource(nextSeed()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// *****************************************************************************
// 产生int类型的随机数,在对应范围内
// *****************************************************************************

// RandomInt 返回区间 [0, max) 内的伪随机整数。
func (generator) RandomInt(max int) int {
	if max < 1 {
		return 0
	}
	randExecute := rand.New(rand.NewSource(nextSeed()))
	return randExecute.Intn(max)
}
