// Run with go test -bench=.
package test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bitlux/caches/util"
)

const targetLen = 28

var pi string

func init() {
	pi = util.ReadLines("../data/pi")[0]
}

func slidingWindowCount(blackbirds map[byte]bool) string {
	i := 0
	count := 0
	for ; i < targetLen; i++ {
		if blackbirds[pi[i]] {
			count++
		}
	}
	for ; i < len(pi)-targetLen; i++ {
		if blackbirds[pi[i]] {
			count++
		}
		if blackbirds[pi[i-targetLen]] {
			count--
		}
		if count == targetLen {
			return pi[i-targetLen+1 : i+1]
		}
	}
	return ""
}

func BenchmarkBlackbirds(b *testing.B) {
	blackbirds := map[byte]bool{}
	for _, d := range util.CBF("blackbirds") {
		blackbirds[byte('0'+d)] = true
	}

	b.Run("SlidingWindow", func(b *testing.B) {
		for b.Loop() {
			slidingWindowCount(blackbirds)
		}
	})

	digits := ""
	for d := range blackbirds {
		digits += fmt.Sprintf("%c", d)
	}
	re := regexp.MustCompile(fmt.Sprintf("[%s]{%d}", digits, targetLen))

	b.Run("RegExp", func(b *testing.B) {
		for b.Loop() {
			re.FindString(pi)
		}
	})
}
