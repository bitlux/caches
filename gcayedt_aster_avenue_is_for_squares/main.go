package main

import (
	"fmt"
	"os"

	"github.com/bitlux/caches/util"
)

const ct = "ECCPD YDSQW AWRWA VDTKK QQAZL VMLAZ ZTBDD DBXDD DXBXR RRMQV TBETN WBMWX ETASS ABXDD DXBXU WDTRR RMQVT BAYDD DUANB YTRNI VRSFD TCW"

func main() {
	fs := util.NewFourSquare(os.Args[1], os.Args[2], true)
	fmt.Println(fs)
	fmt.Println(fs.Encode(ct))
}
