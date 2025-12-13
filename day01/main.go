package main

import (
	input "aoc2015/util"
	"fmt"
)

func main() {
	data := input.ReadInputString("input/day01-real.txt")
	floor := 0
	index := -1
	for i, c := range data {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		}
		// did we enter the basement floor?
		// don't check if we already found the first one
		if index == -1 && floor == -1 {
			index = i + 1 // steps are indexed from 1
		}
	}
	fmt.Printf("part 1: %d\n", floor)
	fmt.Printf("part 2: %d\n", index)
}
