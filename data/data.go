package data

import (
	_ "embed"
	"iter"
	"strconv"
	"strings"
	"sync"

	"github.com/bitlux/caches/util"
)

// ----------------------------------------------------------------------------
// Word lists

//go:embed wordlist.1000.txt
var oneK string

//go:embed wordlist.10000.txt
var tenK string

//go:embed wordlist-20210729.txt
var large string

func OneK() iter.Seq[string] {
	return dict(oneK)
}

func OneKSet() map[string]bool {
	return set(OneK())
}

func TenK() iter.Seq[string] {
	return dict(tenK)
}

func TenKSet() map[string]bool {
	return set(TenK())
}

func Large() iter.Seq[string] {
	return dict(large)
}

func LargeSet() map[string]bool {
	s := set(Large())
	return s
}

func dict(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for k := range strings.Lines(s) {
			if !yield(strings.TrimSpace(k)) {
				return
			}
		}
	}
}

func set(it iter.Seq[string]) map[string]bool {
	m := map[string]bool{}
	for s := range it {
		m[s] = true
	}
	return m
}

// ----------------------------------------------------------------------------
// Bigrams

//go:embed english_2grams.csv
var bigrams string

var (
	bigramMap  map[string]int
	bigramOnce sync.Once
)

func initBigramMap() {
	bigramMap = map[string]int{}
	for line := range strings.Lines(bigrams) {
		tokens := strings.Split(strings.TrimSpace(line), ",")
		n, err := strconv.Atoi(tokens[1])
		util.Must(err)
		bigramMap[tokens[0]] = n
	}
}

func BigramScore(s string) int {
	bigramOnce.Do(initBigramMap)

	ret := 0
	for i := range len(s) {
		if i == len(s)-1 {
			return ret
		}
		ret += bigramMap[strings.ToLower(s[i:i+2])]
	}
	return ret
}
