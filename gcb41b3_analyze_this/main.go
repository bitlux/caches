package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/util"
)

const text = `The Zen of Python, by Tim Peters

Beautiful is better than ugly.
Explicit is better than implicit.
Simple is better than complex.
Complex is better than complicated.
Flat is better than nested.
Sparse is better than dense.
Readability counts.
Special cases aren't special enough to break the rules.
Although practicality beats purity.
Errors should never pass silently.
Unless explicitly silenced.
In the face of ambiguity, refuse the temptation to guess.
There should be one-- and preferably only one --obvious way to do it.
Although that way may not be obvious at first unless you're Dutch.
Now is better than never.
Although never is often better than *right* now.
If the implementation is hard to explain, it's a bad idea.
If the implementation is easy to explain, it may be a good idea.
Namespaces are one honking great idea -- let's do more of those!`

func main() {
	histo := util.RuneCount(strings.ToLower(text))
	fmt.Println(histo['m'] + histo['u'])
	fmt.Println(histo['x'] + histo['y'])
	fmt.Println(histo['x'] + histo['p'])
	fmt.Println(histo['x'] + histo['k'])
	fmt.Println(histo['g'] + histo['z'])
	fmt.Println(histo['g'] + histo['v'] + histo['w'])
	fmt.Println(histo['r'] + histo['z'])
	fmt.Println(histo['l'] + histo['r'] + histo['x'])
}
