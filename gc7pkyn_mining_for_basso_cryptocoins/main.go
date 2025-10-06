// Runtimes:
//
//	    | slowboi | ziltoid | longest
//	len | runtime | runtime | prefix
//	--- | ------- | --------| -------
//	  4 |     10s |      2s |       6
//	  5 |     20m |      4m |       8
//	  6 |  29h35m |         |       9
package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/bitlux/caches/util"
)

const (
	prefix    = "bitlux_"
)

var (
	target     = flag.Int("target", 6, "number of leading 0's to look for")
	length     = flag.Int("length", 4, "number of characters to append")
	checkpoint = flag.String("checkpoint", "", "checkpoint file to load from")
)

type token struct {
	value string
	hash  string // Hex-encoded
	count int
}

func (t *token) String() string {
	return fmt.Sprintf("%d %s %s", t.count, t.value, t.hash)
}

func zeroCount(s string) token {
	tok := token{value: prefix + s}
	hash := md5.Sum([]byte(tok.value))
	tok.hash = hex.EncodeToString(hash[:])

	for i := 0; tok.hash[i] == '0'; i++ {
		tok.count++
	}
	return tok
}

func findHashes(s string, ch chan<- string) {
	if len(s) == *length {
		if t := zeroCount(s); t.count >= *target {
			ch <- t.String()
		}
		return
	}

	for x := 33; x < 127; x++ {
		findHashes(fmt.Sprintf("%s%c", s, x), ch)
	}
}

func main() {
	flag.Parse()
	fmt.Println("target =", *target, "len =", *length)

	var wq *util.WorkQueue
	if len(*checkpoint) != 0 {
		wq = util.NewFromCheckpoint(*checkpoint)
	} else {
		var items []string
		var makeWork func(string)
		makeWork = func(path string) {
			if len(path) == *length/2 {
				items = append(items, path)
				return
			}
			for a := 33; a < 127; a++ {
				makeWork(fmt.Sprintf("%s%c", path, a))
			}
		}
		makeWork("")
		wq = util.NewWorkQueue(items, findHashes)
	}
	wq.Run()
}
