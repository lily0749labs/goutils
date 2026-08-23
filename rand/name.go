package rand

import (
	"math/rand"

	engname "github.com/yelinaung/eng-name"

	"github.com/lily0749labs/goutils/time"
)

func RandEnName() string {
	seed := time.Time.NowUnixNano()
	randName := engname.New(seed)
	rr := rand.New(rand.NewSource(time.Time.NowUnixNano()))
	if rr.Intn(2) == 0 {
		return randName.GetMenName()
	}
	return randName.GetWomenName()
}

func RandEnMenName() string {
	seed := time.Time.NowUnixNano()
	randName := engname.New(seed)
	return randName.GetMenName()
}

func RandEnWomenName() string {
	seed := time.Time.NowUnixNano()
	randName := engname.New(seed)
	return randName.GetWomenName()
}
