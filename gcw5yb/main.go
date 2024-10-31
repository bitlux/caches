package main

import (
	"fmt"
	"math/big"

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

func decode(n *big.Int) string {
	m := big.NewInt(256)
	s := ""
	for n.Cmp(big.NewInt(1)) == 1 {
		c := new(big.Int).Mod(n, m)
		s = string(c.Int64()) + s
		n.Div(n, m)
	}
	return s
}

func main() {
	n := strToInt(kN)
	p := strToInt(kP)
	q := strToInt(kQ)

	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	d := new(big.Int).ModInverse(big.NewInt(e), phi)

	for _, c := range ciphertexts {
		m := new(big.Int).Exp(strToInt(c), d, n)
		fmt.Println(decode(m))
	}
}
