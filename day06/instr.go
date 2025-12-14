package main

import (
	"aoc2015/util"
	"strings"
)

type Op int

const (
	TurnOn Op = iota
	TurnOff
	Toggle
)

type Instr struct {
	operation Op
	target    Rect
}

func parseInput(path string) []Instr {
	lines := util.ReadInputLines(path)
	instructions := make([]Instr, 0)
	for _, line := range lines {
		instructions = append(instructions, parseInstr(line))
	}
	return instructions
}

func parseInstr(s string) Instr {
	s = strings.Replace(s, "turn on", "turn_on", 1)
	s = strings.Replace(s, "turn off", "turn_off", 1)
	parts := strings.Split(s, " ")
	return Instr{
		operation: parseOp(parts[0]),
		target: Rect{
			p1: parseCoord(parts[1]),
			p2: parseCoord(parts[3]),
		},
	}
}

func parseOp(s string) Op {
	switch s {
	case "turn_on":
		return TurnOn
	case "turn_off":
		return TurnOff
	default:
		return Toggle
	}
}

func parseCoord(s string) util.Coord {
	parts := strings.Split(s, ",")
	return util.Coord{
		X: util.MustParseInt(parts[0]),
		Y: util.MustParseInt(parts[1]),
	}
}
