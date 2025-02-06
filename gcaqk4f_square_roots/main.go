package main

import (
	"fmt"

	"github.com/keep94/sqroot/v3"
)

func main() {
	n := sqroot.Sqrt(3)
	for _, x := range []int{2, 1, 3, 68, 500187, 325001, 263263} {
		fmt.Printf("%d", n.At(x))
	}
	fmt.Println()

	fmt.Println(sqroot.FindFirst(sqroot.Sqrt(12), []int{9, 9, 9, 9, 9, 9, 9, 9, 9}))

	const (
		length = 200_000
		target = 900_000
	)
	sum := 0

	n := sqroot.Sqrt(2)
	for index, value := range n.All() {
		sum += value
		if index >= length {
			sum -= n.At(index - length)
		}
		if sum == target {
			fmt.Println(index - length + 1)
			break
		}
	}
}
