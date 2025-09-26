//	 len | runtime | longest 0 prefix
//	 --- | ------- | ----------------
//		 4 |     10s |                6
//		 5 |     20m |                8
//     6 |  29h35m |                9
package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bitlux/caches/gc7pkyn_mining_for_basso_cryptocoins/work"
	"github.com/bitlux/caches/util"
)

const (
	prefix    = "bitlux_"
	tokenBase = "tokens"
)

var (
	target     = flag.Int("target", 6, "number of leading 0's to look for")
	length     = flag.Int("length", 4, "number of characters to append")
	workers    = flag.Int("workers", 10, "number of workers")
	checkpoint = flag.String("checkpoint", "", "checkpoint file to load from")

	ch chan token
)

type token struct {
	value string
	hash  string
	count int
}

func zeroCount(b []byte) token {
	tok := token{value: prefix + string(b)}
	hash := md5.Sum([]byte(tok.value))
	tok.hash = hex.EncodeToString(hash[:])

	for i := 0; tok.hash[i] == '0'; i++ {
		tok.count++
	}
	return tok
}

func findHashes(b []byte) {
	if len(b) == *length {
		if t := zeroCount(b); t.count >= *target {
			ch <- t
		}
		return
	}

	for x := 33; x < 127; x++ {
		findHashes(append(b, byte(x)))
	}
}

func makeWork(yield func(string) bool, prefix string, itemSize int) {
	if len(prefix) == itemSize {
		yield(prefix)
		return
	}
	for a := 33; a < 127; a++ {
		makeWork(yield, fmt.Sprintf("%s%c", prefix, a), itemSize)
	}
}

func main() {
	flag.Parse()

	itemSize := *length / 2
	fmt.Println("target =", *target, "len =", *length, "itemSize =", itemSize)
	tokenFilename := fmt.Sprintf("%s_%d_%d", tokenBase, *target, *length)

	var wq *work.Queue
	var tokenFP *os.File
	var err error
	if len(*checkpoint) != 0 {
		wq = work.NewFromCheckpoint(*checkpoint)
		tokenFP, err = os.OpenFile(tokenFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	} else {
		wq = work.NewFromSeq(func(yield func(string) bool) {
			makeWork(yield, "", itemSize)
		})
		tokenFP, err = os.Create(tokenFilename)
	}
	util.Must(err)

	// Reading token
	ch = make(chan token)
	done := make(chan struct{})
	go func() {
		for t := range ch {
			_, err := fmt.Fprintf(tokenFP, "%s %s %d %s\n", t.value, t.hash, t.count, time.Now().Format(time.DateTime))
			util.Must(err)
		}
		util.Must(tokenFP.Close())
		done <- struct{}{}
	}()

	sem := make(chan struct{}, *workers)
	for value, ok := wq.Next(); ok; value, ok = wq.Next() {
		sem <- struct{}{}
		go func() {
			findHashes([]byte(value))
			wq.MarkFinished(value)
			<-sem
		}()
	}

	// Wait for completion
	for n := *workers; n > 0; n-- {
		sem <- struct{}{}
	}
	close(ch)
	<-done
	fmt.Println()
}
