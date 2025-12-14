package main

import "strings"

func isNice(s string) bool {
	vowelCount := 0
	duplicatedLetter := false
	for i, c := range s {
		if isVowel(c) {
			vowelCount += 1
		}
		if i > 0 {
			prev := rune(s[i-1])
			duplicatedLetter = duplicatedLetter || prev == c
		}
	}
	return vowelCount >= 3 && duplicatedLetter && !containsBannedWord(s)
}

const vowels = "aeiou"

func isVowel(c rune) bool {
	return strings.Contains(vowels, string(c))
}

var bannedWords = []string{"ab", "cd", "pq", "xy"}

func containsBannedWord(s string) bool {
	for _, w := range bannedWords {
		if strings.Contains(s, w) {
			return true
		}
	}
	return false
}
