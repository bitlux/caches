package main

import (
	"errors"
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/bitlux/caches/util"
)

const (
	kN = "219753280453276360597069096621551388190092992509333925115459919"
	kP = "11860681821449926715169841912469"
	kQ = "18527879236744611830104086196051"
	e  = 65537
)

var ciphertexts = []string{
	"167920487915404547623873388967962871030201620708831816122323385",
	"37799817204385669244663809227945788377431419638490866485702822",
	"101824037531968900625851745425851023117556444756813166683348994",
	"33361206922153964532810760417793268547062886751178312579060387",
	"203211931590765493762089517885947364583042600365435117932488094",
	"174982532670024402744606370490906644395944184530097167882382567",
	"63966372043365776189142976956025713220201767300007355997480286",
}

func strToInt(s string) *big.Int {
	n, ok := new(big.Int).SetString(s, 10)
	util.MustBool(ok)
	return n
}

func g(x *big.Int, n *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(x, x), big.NewInt(1)), n)
}

func pollardRho(n *big.Int) (*big.Int, error) {
	x := big.NewInt(2)
	y := x
	d := big.NewInt(1)

	start := time.Now()

	for i := 0; d.Cmp(big.NewInt(1)) == 0; i++ {
		if i%1_000_000 == 0 {
			fmt.Printf("\r%s: %d", time.Now().Format(time.DateTime), i)
		}
		x = g(x, n)
		y = g(g(y, n), n)
		d = new(big.Int).GCD(nil, nil, new(big.Int).Abs(new(big.Int).Sub(x, y)), n)
	}

	fmt.Printf("d = %s, took %s\n", d, time.Since(start))

	if d.Cmp(n) == 0 {
		return nil, errors.New("no non-trivial factor found")
	}
	return d, nil
}

func decode(n *big.Int) string {
	m := big.NewInt(256)
	var runes []rune
	for n.Cmp(big.NewInt(1)) == 1 {
		c := new(big.Int).Mod(n, m)
		runes = append(runes, rune(c.Int64()))
		n.Div(n, m)
	}
	slices.Reverse(runes)
	return string(runes)
}

func main() {
	n := strToInt(kN)
	var p, q *big.Int
	if kP == "" {
		p, err := pollardRho(n)
		util.Must(err)
		q = new(big.Int).Div(n, p)
		fmt.Printf("%s = %s * %s\n", n, p, q)
	} else {
		p = strToInt(kP)
		q = strToInt(kQ)
	}

	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	start := time.Now()
	d := new(big.Int).ModInverse(big.NewInt(e), phi)
	fmt.Println("inverse took ", time.Since(start))

	for _, c := range ciphertexts {
		m := new(big.Int).Exp(strToInt(c), d, n)
		fmt.Println(decode(m))
	}
}
