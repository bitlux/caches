package main

import "fmt"

func negativeOneToTheN(n int) int {
	if n%2 == 0 {
		return 1
	}
	return -1
}

func v(n int) int {
	return (n*n + 3*n*(negativeOneToTheN(n)+1) + 7*(negativeOneToTheN(n+1)+1)/2) / 8
}

func main() {
	for i := range 20 {
		fmt.Println(i, v(i))
	}
	fmt.Println()
	for _, i := range []int{17, 41, 80, 1, 13, 28, 179} {
		fmt.Println(i, v(i))
	}
}
