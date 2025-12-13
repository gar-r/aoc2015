package main

import (
	input "aoc2015/util"
	"strconv"
	"strings"
)

func parseBoxes(path string) []Box {
	lines := input.ReadInputLines(path)
	boxes := make([]Box, 0)
	for _, line := range lines {
		parts := strings.Split(line, "x")
		boxes = append(boxes, Box{
			mustParseInt(parts[0]),
			mustParseInt(parts[1]),
			mustParseInt(parts[2]),
		})
	}
	return boxes
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("cannot parse '" + s + "': " + err.Error())
	}
	return i
}
