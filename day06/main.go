package main

import (
	"aoc2015/util"
	"fmt"
	"iter"
)

func main() {
	instructions := parseInput("input/day06-real.txt")
	fmt.Printf("part 1: %d\n", getEnabledLightCount(instructions))
	fmt.Printf("part 2: %d\n", getTotalBrightness(instructions))
}

func apply(r Rect, f func(c util.Coord)) {
	for c := range r.coords() {
		f(c)
	}
}

type Rect struct {
	p1, p2 util.Coord
}

func (r Rect) coords() iter.Seq[util.Coord] {
	return func(yield func(util.Coord) bool) {
		// we are relying on p1 < p2 in the input
		for row := r.p1.X; row <= r.p2.X; row++ {
			for col := r.p1.Y; col <= r.p2.Y; col++ {
				if !yield(util.Coord{X: row, Y: col}) {
					return
				}
			}
		}
	}
}
