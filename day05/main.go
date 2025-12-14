package main

import (
	"aoc2015/util"
	"fmt"
)

func main() {
	lines := util.ReadInputLines("input/day05-real.txt")
	nice_v1 := 0
	nice_v2 := 0
	for _, line := range lines {
		if isNice(line) {
			nice_v1 += 1
		}
		if isNiceV2(line) {
			nice_v2 += 1
		}
	}
	fmt.Printf("part 1: %d\n", nice_v1)
	fmt.Printf("part 2: %d\n", nice_v2)
}
