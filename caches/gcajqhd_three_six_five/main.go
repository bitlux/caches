package main

import "fmt"

func sq(i int) int {
	return i * i
}

func main() {
	for i := 10; i < 100; i++ {
		if sq(i)+sq(i+1)+sq(i+2) == 365 {
			fmt.Println(i)
		}
		if sq(i)+sq(i+1) == 365 {
			fmt.Println(i)
		}
	}
}
