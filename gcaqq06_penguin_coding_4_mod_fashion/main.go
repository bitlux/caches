package main

import "fmt"

func main() {
	x := 1
	for {
		if x%3 == 1 && x%5 == 2 && x%11 == 10 && x%17 == 9 && x%23 == 19 && x%29 == 10 && x%37 == 5 {
			break
		}
		x++
	}

	y := 1
	for {
		if y%7 == 1 && y%13 == 1 && y%19 == 10 && y%31 == 23 && y%41 == 21 && y%43 == 20 {
			break
		}
		y++
	}

	fmt.Println(x, y)
}
