package main

import "testing"

func Test_EscapedLen(t *testing.T) {
	tests := []struct {
		s string
		l int
	}{
		{`""`, 0},
		{`"abc"`, 3},
		{`"aaa\"aaa"`, 7},
		{`"\x27"`, 1},
	}
	for _, test := range tests {
		actual := escapedLen(test.s)
		if test.l != actual {
			t.Errorf("expected %s => %d, got %d\n", test.s, test.l, actual)
		}
	}
}

func Test_EncodedLen(t *testing.T) {
	tests := []struct {
		s string
		l int
	}{
		{`""`, 6},
		{`"abc"`, 9},
		{`"aaa\"aaa"`, 16},
		{`"\x27"`, 11},
	}
	for _, test := range tests {
		actual := encodedLen(test.s)
		if test.l != actual {
			t.Errorf("expected %s => %d, got %d\n", test.s, test.l, actual)
		}
	}
}
