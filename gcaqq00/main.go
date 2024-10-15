package main

import "fmt"

const Max = 30

func main() {
	seq := []int{-222, 563}

	for i := 2; i <= Max; i++ {
		seq = append(seq, seq[i-1]+seq[i-2])
	}
	fmt.Println(seq)
}
