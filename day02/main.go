package main

import (
	"aoc2015/util"
	"fmt"
)

type Box struct {
	w, h, l int
}

func (b *Box) requiredPaper() int {
	s1 := b.l * b.w
	s2 := b.w * b.h
	s3 := b.h * b.l
	smallest := util.Min(util.Min(s1, s2), s3)
	return 2*(s1+s2+s3) + smallest
}

func (b *Box) requiredPaperForRibbon() int {
	s1 := util.Min(b.w, b.h)
	s2 := util.Min(b.l, util.Max(b.w, b.h))
	wrap := 2*s1 + 2*s2
	ribbon := b.w * b.h * b.l
	return wrap + ribbon
}

func main() {
	boxes := parseBoxes("input/day02-real.txt")
	paper := 0
	ribbon := 0
	for _, box := range boxes {
		paper += box.requiredPaper()
		ribbon += box.requiredPaperForRibbon()
	}

	fmt.Printf("part 1: %d\n", paper)
	fmt.Printf("part 2: %d\n", ribbon)
}
