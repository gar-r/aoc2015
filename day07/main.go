package main

import (
	"fmt"
	"strconv"
	"strings"
)

var traceMemoized *memoizedTraceWire

func main() {
	wires := parseInput("input/day07-real.txt")

	traceMemoized = &memoizedTraceWire{
		fn:    traceWire,
		cache: make(map[string]uint16),
	}
	result := traceMemoized.call(wires, "a")

	fmt.Printf("part 1: %d\n", result)

	wires["b"] = strconv.Itoa(int(result)) // override b to the result
	traceMemoized.cache = make(map[string]uint16)

	result = traceMemoized.call(wires, "a")
	fmt.Printf("part 2: %d\n", result)
}

type traceWireFn func(map[string]string, string) uint16
type memoizedTraceWire struct {
	fn    traceWireFn
	cache map[string]uint16
}

func (m *memoizedTraceWire) call(wires map[string]string, out string) uint16 {
	if res, ok := m.cache[out]; ok {
		return res
	}
	res := m.fn(wires, out)
	m.cache[out] = res
	return res
}

func traceWire(wires map[string]string, out string) uint16 {
	n, err := strconv.Atoi(out)
	if err == nil {
		return uint16(n)
	}
	expr := wires[out]
	parts := strings.Split(expr, " ")
	if len(parts) == 1 { // simple value
		return traceMemoized.call(wires, parts[0])
	}
	if parts[0] == "NOT" {
		return ^(traceMemoized.call(wires, parts[1]))
	}
	switch parts[1] {
	case "AND":
		return traceMemoized.call(wires, parts[0]) & traceMemoized.call(wires, parts[2])
	case "OR":
		return traceMemoized.call(wires, parts[0]) | traceMemoized.call(wires, parts[2])
	case "LSHIFT":
		return traceMemoized.call(wires, parts[0]) << traceMemoized.call(wires, parts[2])
	case "RSHIFT":
		return traceMemoized.call(wires, parts[0]) >> traceMemoized.call(wires, parts[2])
	default:
		panic(expr)
	}
}

func isValue(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}
