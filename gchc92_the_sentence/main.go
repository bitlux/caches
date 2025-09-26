package main

import (
	"fmt"
	"regexp"
	"strings"
)

var story = `Judge Roy Bean looked down upon Orfus Reed, the miserable wretch before him, and began to pass sentence:

"Neither the wails of your mother nor the pleas for mercy you mouth will avail you," he declared solemnly. With Reed quaking quite literally in his boots, Bean took a swig of whiskey and continued. "It is said that with enough penance it is possible to find redemption for your sins, perhaps even for a crime this heinous. Such sentimental speculation is not worth a plug nickel, in my opinion." His eyes ablaze, rotgut spilling from his glass, Bean became more agitated. "As I weigh this offense against the words of the statutes of Texas, I find that I cannot by law sentence you to be hanged. I realize rogues like you deserve to die on the gallows, but our legislature is not so enlightened. So that won't happen today, at least."

After yet another big gulp of liquor, the judge continued, "We stand in judgment over you today because you stole a man's horse, coat, and food and left him to die on the prairie on a cold winter's night. No one would find out, you thought. It would have been much worse for you today if he had perished - your neck would be stretched by sundown, I assure you, and no lawbook could have saved you. But he was a man of great stamina and resourcefulness; he managed to cut wood from the brush with his buck knife and start a fire. He nearly froze, roaming around the flatlands searching for scrub brush for hours, but he lived to tell us who it was that left him there. You thought that you could obtain mercy from us, your neighbors, but if our sympathies lie with any man, it is with Bob Carson, who almost died because of you." Reed began to shake and sweat, wondering what Bean was driving at. "If our custom this time of year is to give gifts, then I must follow it. It's Christmas Eve now, so I have a little present for you," Bean guffawed.

"I'm going to give you the same chance you gave Bob Carson. By the power vested in me by the state of Texas, I hereby sentence you to be taken out to the exact spot where you left your victim, where you are to be left with no horse, no food, and no coat. Of course, we can't let a felon have a deadly weapon, now can we, so no knife and no matches for you." Bean drew the gavel down with the force of a blacksmith's hammer. Reed screamed in terror as he was dragged away, but it was in vain. He was never seen again.`

func main() {
	fields := strings.Fields(strings.ToLower(story))
	story = strings.Join(fields, "")
	for _, punc := range []string{",", `"`, ".", "'"} {
		story = strings.ReplaceAll(story, punc, "")
	}

	re := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine)`)
	for _, m := range re.FindAllStringSubmatch(story, -1) {
			fmt.Println(m[0])
	}
}
