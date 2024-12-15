package util

import (
	"bufio"
	"os"
)

func ReadLines(file string) []string {
	f, err := os.Open(file)
	Must(err)

	var ret []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	Must(scanner.Err())
	return ret
}
