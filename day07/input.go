package main

import (
	"aoc2015/util"
	"strings"
)

func parseInput(path string) map[string]string {
	wires := make(map[string]string)
	lines := util.ReadInputLines(path)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		wires[parts[1]] = parts[0]
	}
	return wires
}
