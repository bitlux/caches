package util

import (
	"fmt"
	"os"
)

// Must prints err and exits if err is not nil.
func Must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// MustBool exits if b is false.
func MustBool(b bool) {
	if !b {
		fmt.Println("unexpected false value")
		os.Exit(1)
	}
}
