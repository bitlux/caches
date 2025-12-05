package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitlux/caches/util"
)

var data = `19	23	30	 7	75	38	60	69	80
 6	74	28	37	19	 8	 5	26	 2
 5	32	 5	23	19	 2	 3	11	34
 8	68	64	29	29	23	13	22	 5
22	61	22	98	16	25	96	 8	29
17	23	29	48	 5	99	 4	16	32
17	36	99	22	36	28	23	30	28
 2	61	32	18	 6	 5	25	23	38
19	33	 8	22	69	26	32	 8	18`

func main() {
	fields := strings.Fields(data)
	fmt.Println(len(fields))

	var ints [][]int
	for i := range 9 {
		var row []int
		for j := range 9 {
			n, err := strconv.Atoi(fields[9*i+j])
			util.Must(err)
			row = append(row, n)
		}
		ints = append(ints, row)
	}

	for i := range 9 {
		sum := 0
		for j := range 9 {
			sum += ints[i][j]
		}
		fmt.Println(sum)
	}

}
