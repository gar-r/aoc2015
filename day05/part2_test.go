package main

import "testing"

func Test_IsNiceV2(t *testing.T) {
	tests := []struct {
		str    string
		isNice bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}
	for _, test := range tests {
		actual := isNiceV2(test.str)
		if actual != test.isNice {
			t.Errorf("expected %v, got %v for %s\n", test.isNice, actual, test.str)
		}
	}
}
