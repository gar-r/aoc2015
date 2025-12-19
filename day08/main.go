package main

import (
	"aoc2015/util"
	"fmt"
)

func main() {
	lines := util.ReadInputLines("input/day08-real.txt")
	sum := 0
	for _, line := range lines {
		sum += len(line)
		sum -= escapedLen(line)
	}
	fmt.Printf("part 1: %d\n", sum)

	sum = 0
	for _, line := range lines {
		sum += encodedLen(line)
		sum -= len(line)
	}
	fmt.Printf("part 2: %d\n", sum)
}

func escapedLen(raw string) int {
	escaping := 0
	result := 0
	for _, c := range raw[1 : len(raw)-1] { // skip surrounding double-quotes
		if escaping > 0 {
			if c == 'x' {
				escaping += 2
			}
			escaping -= 1
			if escaping == 0 {
				result += 1
			}
			continue
		}
		if c == '\\' {
			escaping += 1
			continue
		}
		result += 1
	}
	return result
}

func encodedLen(raw string) int {
	result := 0
	for _, c := range raw {
		if c == '\\' || c == '"' {
			result += 2
		} else {
			result += 1
		}
	}
	return result + 2 // add surrounding double-quotes
}
