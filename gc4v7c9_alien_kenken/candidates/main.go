package main

import (
	"fmt"

	c "github.com/bitlux/caches/gc4v7c9_alien_kenken/common"
)

func main() {
	for _, tc := range []struct {
		op     func([]int) int
		size   int
		target int
	}{
		//	{c.Clubs, 5, 237},   // {1, 5, 7, 9, 9}
		//	{c.Diamonds, 3, 29}, // {2, 3, 6}
		//	{c.Diamonds, 5, 72}, // {v2, 4, v6, 7, 8} or {v3, v4, 4, 7, 8}

		{c.Hearts, 5, 36},
		//	{c.Diamonds, 2, 13}, // {1, 5}
		{c.Hearts, 5, 34},

		//	{c.Hearts, 4, 39}, // almost certainly contains 9
		//	{c.Hearts, 4, 20},

		//	{c.Hearts, 3, 9}, // {1, 3, 4} {2, 3, 4}
		//	{c.Spades, 5, 104}, // {2, 6, 7, 8, 8}
		// {c.Hearts, 3, 18}, // {1, 2, 7} {1, 5, 6}

		//	{c.Clubs, 2, 34}, // {3, 5}
		{c.Spades, 2, 15}, // {9, 2}

		//	{c.Clubs, 5, 130},
		// {c.Spades, 4, 21},
		// {c.Diamonds, 4, 79},
		//	{c.Clubs, 5, 91},

		//	{c.Diamonds, 5, 84}, // {2, 6, 7v, 8, 9v} {3v, 5, 6, 8, 8v} {4v, 5, 6, 6v, 8}
		//	{c.Hearts, 2, 28}, // {4, 9}
		{c.Hearts, 5, 35}, //
		// Valid for 18h = [2 7] -- 123478v 569^
		// [1v 5 6v 6 9]
		// [2v 5 6v 6 9]
		// [4v 5v 5 6 9]

		// Valid for 18h = [3 6] -- 13456v 29^
		// [1v 2v 2 8 9]
		// [1v 2 3v 8 9]
		// [2 2v 3v 8 9]
		// [2 3v 5v 7 9]
		// [2 4v 5v 7 9]

		// {c.Spades, 3, 24}, // {1, 7, 9} {2, 6, 7} {3, 4, 8}
	} {
		fmt.Println(tc.target)
		for _, cand := range c.Candidates(tc.op, tc.target, tc.size) {
			fmt.Println("  ", cand)
		}
	}
}
