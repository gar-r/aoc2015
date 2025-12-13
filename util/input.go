package input

import (
	"os"
)

// ReadInputString reads the whole contents of the text file at path, or panics
func ReadInputString(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err) // panic if we couldn't read the input for any reason
	}
	return string(bytes)
}
