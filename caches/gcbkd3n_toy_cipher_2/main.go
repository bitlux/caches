package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bitlux/caches/lib/data"
)

const (
	keyspace = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ,-. "
	N        = len(keyspace)
)

func decode(ct, key string) (string, error) {
	if len(ct) > len(key) {
		return "", errors.New("Key is shorter than ciphertext")
	}

	pt := ""
	for i := range ct {
		ki := strings.Index(keyspace, key[i:i+1])
		cti := strings.Index(keyspace, ct[i:i+1])
		pti := (cti + ki) % N
		pt += fmt.Sprintf("%s", keyspace[pti:pti+1])
	}
	return pt, nil
}

func main() {
	for _, pair := range []struct {
		ct, key string
	}{
		{"GY4.TVOE20BV4L2OKEHCV.B", "7931DGEZ24P8SK0HJVRQFAUNY"},
		{"UK7Y5M8UU 58QGXCC6WJEE", "BKL59H-JG43XD76WT.EUQZAMO"},
		{"09-05 V1-VY,YWV0.02.3", data.PiString()},
	} {
		pt, err := decode(pair.ct, pair.key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pt)
	}
}
