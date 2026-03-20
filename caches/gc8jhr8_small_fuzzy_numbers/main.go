package main

import (
	"fmt"

	"github.com/d4l3k/go-bfloat16"
)

func main() {
	for _, f := range []float32{7.112e-37, 5.031e-05, 5.049e-28, 3.964e-06} {
		fmt.Printf("%x\n", bfloat16.FromFloat32(f))
	}
}
