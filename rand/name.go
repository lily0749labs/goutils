package rand

import (
	"math/rand"

	engname "github.com/yelinaung/eng-name"

	"github.com/lily0749labs/goutils/time"
)

// EnglishName 返回一个随机英文名。
func (generator) EnglishName() string {
	seed := time.Time.NowUnixNano()
	randName := engname.New(seed)
	rr := rand.New(rand.NewSource(time.Time.NowUnixNano()))
	if rr.Intn(2) == 0 {
		return randName.GetMenName()
	}
	return randName.GetWomenName()
}

// EnglishMaleName 返回一个随机男性英文名。
func (generator) EnglishMaleName() string {
	seed := time.Time.NowUnixNano()
	randName := engname.New(seed)
	return randName.GetMenName()
}

// EnglishFemaleName 返回一个随机女性英文名。
func (generator) EnglishFemaleName() string {
	seed := time.Time.NowUnixNano()
	randName := engname.New(seed)
	return randName.GetWomenName()
}

// Deprecated: 使用 Rand.EnglishName。
func RandEnName() string { return Rand.EnglishName() }

// Deprecated: 使用 Rand.EnglishMaleName。
func RandEnMenName() string { return Rand.EnglishMaleName() }

// Deprecated: 使用 Rand.EnglishFemaleName。
func RandEnWomenName() string { return Rand.EnglishFemaleName() }
