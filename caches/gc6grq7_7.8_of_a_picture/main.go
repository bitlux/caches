package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/lib/cipher"
)

const words = `ability accordingly activation aneurysm annuity anteater ballplayer banishment benevolent blowtorch brainpower brawny brevity bureaucracy candleholder capitalist capture celebration certitude choreograph complement compostable conceit confront consecrate consort countess crispy crockpot crossbow dehumidifier dehydrate deodorize desolation dollhouse doorbell eavesdrop echocardiogram educator ellipsis encounter entirety evident excelsior exploit feisty formality garrison goddaughter gunpoint hallucinogen handsome hustling immaterial insect joyous lakefront livestock luxury mapmaker marketplace marsupial mastery maternal memorize moderation modesty mutilate neighbor neurotic offspring pathfinder powerful pressure proof quality realignment redundant`

func dollar(s string) int {
	sum := 0
	for _, c := range s {
		sum += cipher.A1Encode(c)
	}
	return sum
}

func main() {
	for word := range strings.FieldsSeq(words) {
		fmt.Printf("%c", dollar(word))
	}
	fmt.Println()
}
