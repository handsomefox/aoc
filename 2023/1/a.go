package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func SolveA(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for sc.Scan() {
		sum += getNumsA(sc.Text())
	}

	return fmt.Sprint(sum)
}

func getNumsA(input string) int {
	var first, last string
	for _, r := range input {
		if unicode.IsDigit(r) {
			if first == "" {
				first = string(r)
			}
			last = string(r)
		}
	}

	return MustParse(first + last)
}
