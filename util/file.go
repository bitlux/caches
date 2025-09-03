package util

import (
	"bufio"
	"os"
)

const maxBuf = 1024 * 1024 * 1024 // 1 GiB

func readLines(file string, size int) ([]string, error) {
	f, err := os.Open(file)
	Must(err)

	buf := make([]byte, size)
	scanner := bufio.NewScanner(f)
	scanner.Buffer(buf, size)

	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, scanner.Err()
}

// ReadLines opens the named file and returns a slice of the lines of the file.
func ReadLines(file string) []string {
	size := bufio.MaxScanTokenSize
	for size < maxBuf {
		ret, err := readLines(file, size)
		if err == nil {
			return ret
		}
		if err == bufio.ErrTooLong {
			size *= 2
			continue
		}
		Must(err)
	}
	Must(bufio.ErrTooLong)
	return nil
}
