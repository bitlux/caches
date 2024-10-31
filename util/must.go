package util

import (
	"fmt"
	"os"
)

func Must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func MustBool(b bool) {
	if !b {
		fmt.Println("unexpected false value")
		os.Exit(1)
	}
}
