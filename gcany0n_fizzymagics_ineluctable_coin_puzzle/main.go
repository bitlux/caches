package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitlux/caches/util"
)

const numbers = `19	23	30	 7	75	38	60	69	80
 6	74	28	37	19	 8	 5	26	 2
 5	32	 5	23	19	 2	 3	11	34
 8	68	64	29	29	23	13	22	 5
22	61	22	98	16	25	96	 8	29
17	23	29	48	 5	99	 4	16	32
17	36	99	22	36	28	23	30	28
 2	61	32	18	 6	 5	25	23	38
19	33	 8	22	69	26	32	 8	18`

func main() {
	for i, str := range strings.Fields(numbers) {
		n, err := strconv.Atoi(str)
		util.Must(err)
		fmt.Printf("%c", util.A1Decode(util.CollatzStoppingTime(n)))
		if (i+1)%9 == 0 {
			fmt.Println()
		}
	}
}
