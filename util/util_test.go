package util_test

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func ExampleFactor() {
	fmt.Println(util.Factor(60))
	// Output: [2 2 3 5]
}

func ExampleIsPrime() {
	fmt.Println(util.IsPrime(101))
	// Output: true
}

func ExampleDigits() {
	fmt.Println(util.Digits(1234))
	// Output: [1 2 3 4]
}
