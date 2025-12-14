package main

import "testing"

func Test_IsNice(t *testing.T) {
	tests := []struct {
		str    string
		isNice bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, test := range tests {
		isNice := isNice(test.str)
		if isNice != test.isNice {
			t.Errorf("expected %v, got %v for %s\n", test.isNice, isNice, test.str)
		}
	}

}
