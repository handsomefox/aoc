package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func IsUpper(r rune) bool {
	if !unicode.IsUpper(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}

func GetPriority(input rune) int {
	if !IsUpper(input) {
		return int(input) - 96
	}
	return int(input) - 38
}

func SolveA(input string) string {
	var (
		sc    = bufio.NewScanner(strings.NewReader(input))
		score = 0
	)
	for sc.Scan() {
		var (
			txt           = sc.Text()
			first, second = txt[:len(txt)/2], txt[len(txt)/2:]
		)
	outer:
		for _, item := range first {
			for _, itemTwo := range second {
				if item != itemTwo {
					continue
				}
				score += GetPriority(item)
				break outer
			}
		}
	}

	return fmt.Sprint(score)
}
