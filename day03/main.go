package main

import (
	"aoc2015/util"
	"fmt"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	directions := strings.TrimSpace(util.ReadInputString("input/day03-real.txt"))
	fmt.Printf("part 1: %d\n", part1(directions))
	fmt.Printf("part 2: %d\n", part2(directions))
}

func part1(directions string) int {
	visited := make(map[Coord]bool)
	pos := Coord{0, 0}
	visited[pos] = true
	for _, dir := range directions {
		move(&pos, dir)
		visited[pos] = true
	}
	return len(visited)
}

func part2(directions string) int {
	visited := make(map[Coord]bool)
	start := Coord{0, 0}
	visited[start] = true
	actors := []Coord{start, start} // 0: santa, 1: robo-santa
	for i, dir := range directions {
		coord := &actors[i%2]
		move(coord, dir)
		visited[*coord] = true
	}
	return len(visited)
}

func move(c *Coord, dir rune) {
	switch dir {
	case '>':
		c.y += 1
	case '<':
		c.y -= 1
	case 'v':
		c.x += 1
	case '^':
		c.x -= 1
	}
}
