package util

// Min returns the smaller of a, b integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of a, b integers
func Max(a, b int) int {
	if b < a {
		return a
	}
	return b
}
