package main

import "fmt"

func v(n int) int {
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return v(n-1) + n - 1
	}
	n /= 2
	return n*(n+1)/2 + 1
}

func main() {
	for _, i := range []int{17, 41, 80, 1, 13, 28, 179} {
		fmt.Println(i, v(i))
	}
}
