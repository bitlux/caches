package main

import "fmt"

func main() {
	cheezburger := 1
	bucket := cheezburger
	icountz := 0
	var nomnom int

	for {
		if icountz < 28 {
			icountz += 2 // ???
			cheezburger -= bucket
			bucket = cheezburger - bucket
		} else {
			nomnom = bucket - 97
			fmt.Printf("W 121 43.%d\n", nomnom)
			break
		}

		if icountz == 12 {
			nomnom = cheezburger * 79
			fmt.Printf("N 37 41.%d", nomnom)
		}
	}
}
