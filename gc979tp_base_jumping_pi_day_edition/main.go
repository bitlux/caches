package main

import (
	"fmt"
	"math"

	"github.com/bitlux/caches/util"
)

func toBasePi(n int, target float64) {
	digits := util.Digits(n)
	sum := 0.0
	power := 1.0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += float64(digits[i]) * power
		power *= math.Pi
	}

	if math.Abs(sum-target) < 0.01 {
		fmt.Println(util.ToCoord(digits))
	}
}

func main() {
	for i := 3737000; i < 3739999; i++ {
		toBasePi(i, 5609.01391540835720661562)
	}
	for i := 12146000; i < 12147999; i++ {
		toBasePi(i, 5914.65474769636512064608)
	}
}
