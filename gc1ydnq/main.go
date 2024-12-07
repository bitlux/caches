package main

import (
	"fmt"
	"time"
)

func main() {
	for _, ts := range []int64{1231112162,
		1199185193,
		1167743901,
		1136108594,
		1136413534,
		1167611505,
		1199243660,
		1230771093,
	} {
		t := time.Unix(ts, 0).UTC()
		fmt.Println(t.Format(time.DateTime))
	}
}
