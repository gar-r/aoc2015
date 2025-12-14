package util

import (
	"fmt"
	"strconv"
)

// MustParseInt converts s to an int, or panics.
// Use this function, when not being able to parse the number is fatal.
func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", s))
	}
	return i
}
