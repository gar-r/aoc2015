package main

import "aoc2015/util"

func getTotalBrightness(instuctions []Instr) int {
	lights := make(map[util.Coord]int)
	for _, instr := range instuctions {
		switch instr.operation {
		case TurnOn:
			apply(instr.target, func(c util.Coord) {
				val, ok := lights[c]
				if ok {
					lights[c] = val + 1
				} else {
					lights[c] = 1
				}
			})
		case TurnOff:
			apply(instr.target, func(c util.Coord) {
				val, ok := lights[c]
				if ok && val > 0 {
					lights[c] = val - 1
				}
			})
		case Toggle:
			apply(instr.target, func(c util.Coord) {
				val, ok := lights[c]
				if ok {
					lights[c] = val + 2
				} else {
					lights[c] = 2
				}
			})
		}
	}

	total := 0
	for _, v := range lights {
		total += v
	}
	return total
}
