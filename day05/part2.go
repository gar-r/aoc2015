package main

import (
	"fmt"
	"strings"
)

func isNiceV2(s string) bool {
	rule1 := false
	rule2 := false
	for i, c := range s {
		if i > 0 {
			pair := fmt.Sprintf("%c%c", s[i-1], c)
			if strings.LastIndex(s, pair) > i {
				rule1 = true
			}
		}
		if i > 1 && i < len(s)-1 && s[i-1] == s[i+1] {
			rule2 = true
		}
	}
	return rule1 && rule2
}
