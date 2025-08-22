package data

import (
	_ "embed"
	"iter"
	"strings"
)

//go:embed wordlist.1000.txt
var oneK string

//go:embed wordlist.10000.txt
var tenK string

//go:embed wordlist-20210729.txt
var large string

func OneK() iter.Seq[string] {
	return dict(oneK)
}

func OneKSet() map[string]bool{
	return set(OneK())
}

func TenK() iter.Seq[string] {
	return dict(tenK)
}

func TenKSet() map[string]bool{
	return set(TenK())
}

func Large() iter.Seq[string] {
	return dict(large)
}

func LargeSet() map[string]bool{
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