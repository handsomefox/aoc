package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))
	totalCards := make(map[int]int)

	for currentCard := 1; sc.Scan(); currentCard++ {
		totalCards[currentCard]++
		matches := calculateMatches(sc.Text())
		for i := 1; i <= matches; i++ {
			totalCards[currentCard+i] += 1 * totalCards[currentCard]
		}
	}

	sum := 0
	for _, v := range totalCards {
		sum += v
	}

	return fmt.Sprint(sum)
}

func calculateMatches(s string) int {
	_, matches := parseCard(s)
	return matches
}
