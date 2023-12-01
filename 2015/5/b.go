package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))
	total := 0
	for sc.Scan() {
		text := sc.Text()
		if text == "" || text == " " || text == "\n" || text == "\r\n" {
			continue
		}
		if stringIsNice2(text) {
			total++
		}
	}
	return fmt.Sprint(total)
}

func stringIsNice2(s string) bool {
	if !hasTwoNonOverlappingLettersPairs(s) {
		return false
	}
	if !hasThreeLetterCombo(s) {
		return false
	}

	return true
}

func hasTwoNonOverlappingLettersPairs(s string) bool {
	if len(s) < 4 {
		return false
	}

	pairs := make(map[string]int)
	previous := string(s[0]) + string(s[1])
	pairs[previous]++

	for i := 2; i < len(s); i++ {
		current := string(s[i-1]) + string(s[i])
		if previous == current {
			return false
		}
		previous = current
		pairs[current]++
	}

	for _, v := range pairs {
		if v >= 2 {
			return true
		}
	}

	return false
}

func hasThreeLetterCombo(s string) bool {
	if len(s) < 3 {
		return false
	}

	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] && s[i] != s[i-1] {
			return true
		}
	}

	return false
}
