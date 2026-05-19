package main

import "fmt"

func main() {
	for lat := 5500; lat <= 6600; lat++ {
		for long := 8000; long <= 9999; long++ {
			if float64(3710000+lat)/float64(12150000+long) == 0.3056483045117939 {
				fmt.Println(lat, long)
			}
		}
	}
}
