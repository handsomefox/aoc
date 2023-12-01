package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveA(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))
	total := 0
	for sc.Scan() {
		text := sc.Text()
		if text == "" || text == " " || text == "\n" || text == "\r\n" {
			continue
		}
		if stringIsNice(text) {
			total++
		}
	}
	return fmt.Sprint(total)
}

func stringIsNice(s string) bool {
	if !hasNVowels(s, 3) {
		return false
	}
	if !hasDoubleLetter(s) {
		return false
	}
	if hasProhibitedString(s) {
		return false
	}

	return true
}

func hasNVowels(s string, n int) bool {
	var (
		c      = 0
		vowels = map[rune]struct{}{
			'a': {},
			'e': {},
			'i': {},
			'o': {},
			'u': {},
		}
	)
	for _, r := range s {
		if _, ok := vowels[r]; ok {
			c++
		}
	}
	return c >= n
}

func hasDoubleLetter(s string) bool {
	if len(s) < 2 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}

func hasProhibitedString(s string) bool {
	prohibited := [...]string{"ab", "cd", "pq", "xy"}
	for _, p := range prohibited {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}
