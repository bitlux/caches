package util

import (
	"bufio"
	"os"
)

// ReadLines opens the named file and returns a slice of the lines of the
// file.
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
