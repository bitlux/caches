package main

import (
	"fmt"
	"os"

	"github.com/bitlux/caches/lib/cipher"
	"github.com/bitlux/caches/lib/util"
)

const ct = "ECCPD YDSQW AWRWA VDTKK QQAZL VMLAZ ZTBDD DBXDD DXBXR RRMQV TBETN WBMWX ETASS ABXDD DXBXU WDTRR RMQVT BAYDD DUANB YTRNI VRSFD TCW"

func main() {
	fs := cipher.NewFourSquare(os.Args[1], os.Args[2], true)
	fmt.Println(fs)
	pt, err := fs.Encode(ct)
	util.Must(err)
	fmt.Println(pt)
}
