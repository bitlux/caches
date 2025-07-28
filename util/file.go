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

	scanner := bufio.NewScanner(f)
	size := 100 * 1024
	buf := make([]byte, size)
	scanner.Buffer(buf, size)

	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	Must(scanner.Err())
	return ret
}
