package main

import (
	"fmt"

	"github.com/bitlux/caches/data"
	"github.com/bitlux/caches/util"
)

func main() {
	m := map[string][]string{}
	for w := range data.Large() {
		sorted := util.SortLetters(w)
		m[sorted] = append(m[sorted], w)
	}

	for _, word := range []string{
		"ropiest", "encamps", "surfeit", "potions",
		"lashing", "regales", "gunshot", "amenity",
		"attends", "aptness", "agonies", "handcar",
		"generic", "platoon", "drowned", "needles",
		"dualities",
	} {
		fmt.Printf("\n%s\n", word)
		for l := range util.Alphabet() {
			meep := m[util.SortLetters(fmt.Sprintf("%s%c", word, l))]
			if len(meep) > 0 {
				fmt.Printf("%c %v\n", l, meep)
			}
		}
	}
}
