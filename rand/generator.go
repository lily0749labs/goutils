package rand

import (
	mathrand "math/rand"
	"strconv"
	"sync/atomic"

	"github.com/lily-study-utils/goutils/time"
)

var (
	// Rand 提供无状态的随机辅助方法。
	Rand = rand{}
	// RandPtr 指向 Rand，为兼容旧代码而保留。
	RandPtr = &Rand
)

type rand struct{}

// GetRand 返回一个使用新种子创建的伪随机数生成器。
func (rand) GetRand() *mathrand.Rand {
	return mathrand.New(mathrand.NewSource(time.Time.NowUnixNano()))
}

var lastOrderNumber atomic.Int64

// GenerateOrderNumber 返回进程内单调递增且并发不重复的数字订单号。
// 订单号由毫秒时间戳和进程内序列组成，不承诺跨进程唯一；跨进程场景请使用 ID.SnowflakeID。
func (rand) GenerateOrderNumber() string {
	base := time.Time.NowUnixMilli() * 1000
	for {
		last := lastOrderNumber.Load()
		next := base
		if next <= last {
			next = last + 1
		}
		if lastOrderNumber.CompareAndSwap(last, next) {
			return strconv.FormatInt(next, 10)
		}
	}
}

// GenOrderNo 返回一个基于时间戳生成的订单号。
// Deprecated: 请使用 GenerateOrderNumber。
func (r rand) GenOrderNo() string {
	return r.GenerateOrderNumber()
}

// 以下状态用于为每次随机生成追加不同的种子偏移量。
var randomParam atomic.Int64

func nextSeed() int64 {
	return randomParam.Add(1) + time.Time.NowUnixNano()
}

// RandomString 返回由数字和大写字母组成的随机字符串。
func (rand) RandomString(size int) string {
	if size <= 0 {
		return ""
	}
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, size)
	r := mathrand.New(mathrand.NewSource(nextSeed()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// RandomStringMix 返回由小写 ASCII 字母组成的随机字符串。
func (rand) RandomStringMix(size int) string {
	if size <= 0 {
		return ""
	}
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, size)
	r := mathrand.New(mathrand.NewSource(nextSeed()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// RandomInt 返回区间 [0, max) 内的伪随机整数。
func (rand) RandomInt(max int) int {
	if max < 1 {
		return 0
	}
	randExecute := mathrand.New(mathrand.NewSource(nextSeed()))
	return randExecute.Intn(max)
}
