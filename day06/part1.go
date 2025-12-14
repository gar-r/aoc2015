package main

import (
	"aoc2015/util"
)

func getEnabledLightCount(instructions []Instr) int {
	lit := make(util.Set[util.Coord])
	for _, instr := range instructions {
		switch instr.operation {
		case TurnOn:
			apply(instr.target, func(c util.Coord) { lit.Insert(c) })
		case TurnOff:
			apply(instr.target, func(c util.Coord) { lit.Remove(c) })
		case Toggle:
			apply(instr.target, func(c util.Coord) {
				if lit.Contains(c) {
					lit.Remove(c)
				} else {
					lit.Insert(c)
				}
			})

		}
	}
	return len(lit)
}
