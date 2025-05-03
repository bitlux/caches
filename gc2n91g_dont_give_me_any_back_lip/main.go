package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/util"
)

var text = []string{
	`Mrs. Zelda Baxter used to be the town's pedagogue, and she lived on a quaint little street just a little ways from me as I grew up with my family and our dog Jake.`,
	`As an instructor she was very strict, but even when she seemed mean she wasn't. Quite frankly she was just so nice, well except on quiz day and she sees you cheat.`,
	`Our class quiz was kinda of hard for most of us. You'd pray hours would go by fast. A tough quiz could just mix up your thoughts and you wish you could avoid it.`,
	`We all respected Mrs. Baxter because in her own way she was still a very smart teacher. And our parents would actually request her to come to our house for dinner. Jake who was a lazy hound, was always excited to see her.`,
	`One thing I will always remember about Mrs. Zelda Baxter was what she would always say to us when we are in a disagreement with her of her teachings “Don't give me any back lip” That would just about keep us quiet till the bell rang.`,
	`From time we start class til time we get out is just torture for us. Every day we do tons of extra work so we become quite tired. Some of us would kinda fall asleep and Mrs. Baxter feels zero tolerance for it.`,
	`We didn't have schools like big fancy towns, ours was just a small school with the grades squished together making them extra crowded. Our parents thought it would be nice to have a larger size school.`,
	`Since our rooms were so crowded we needed to work side by side. The older kids were very helpful with the younger ones to finish their homework. In those times the room size just didn’t seem so little. When the students worked it is so extremely quiet in the room you would think they were sleeping.`,
	`But Mrs. Zelda Baxter was not a teacher to let you sleep in class. She would find ways to keep you busy at all times. If she did see your head on the desk she would quickly snap a ruler down. I think she greatly enjoyed doing that.`,
	`We learned the basics subjects, though she did make it exciting. When she was teaching us geography we felt like explorers. We zipped through every subject quickly but we did retain what we learned.`,
	`Every day was a different subject and we would play in the courtyard for a half an hour 3 times a week. Fridays after the weekly exams she would treat us with a BBQ. Once every month we’d go on a field trip to places like the museum, a park or the zoo.`,
	`On Saturday Vinny and I would go down to Satan's Quarry and go for a swim. I would bring a P&J sandwich, milk and a long stick with string on it. With a worm put on that string, I would drop it in this quarry. Vinny would hunt and find many quartz rocks with vivid colors and add this to his rock box.`,
	`As the days got warmer and the classes felt longer we knew summer would come soon and then no school for three months. But Mrs. Zelda Baxter made sure are heads were not clouded and told us to just read our books and learn from them. She also taught us to watch our Ps & Qs and to play very hard.`,
	`As we neared summer we learned that our teacher would not return the next year. She was going to retire after 55 years of teaching. Prior to the exams, we put together some money and purchased a card for her and we all signed it. It was given to her on our last day. She started to cry. We asked her what was wrong and she said she was just so full of joy. No other teacher could equal her zest for teaching and we would never forget her.`,
}

func main() {
	for _, para := range text {
		m := util.RuneCount(strings.ToLower(para))
		d := 0
		for c := 'a'; c <= 'z'; c++ {
			if m[c] == 0 {
				d = int(util.A1Z26(c))
				break
			}
		}
		fmt.Println(d)
	}
}
