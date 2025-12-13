package util

import (
	"os"
	"strings"
)

// ReadInputString reads the whole contents of the text file at path, or panics
func ReadInputString(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err) // panic if we couldn't read the input for any reason
	}
	return string(bytes)
}

// ReadInputLines reads the contents of the file at path line by line, or panics
func ReadInputLines(path string) []string {
	str := strings.TrimRight(ReadInputString(path), "\n")
	return strings.Split(str, "\n")
}
