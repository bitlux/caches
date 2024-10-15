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
