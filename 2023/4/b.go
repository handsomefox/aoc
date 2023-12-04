package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	totalCards := make(map[int]int)
	currentCard := 1
	for sc.Scan() {
		totalCards[currentCard]++
		matches := calculateMatches(sc.Text())
		for i := 1; i <= matches; i++ {
			totalCards[currentCard+i] += 1 * totalCards[currentCard]
		}
		currentCard++
	}

	sum := 0
	for _, v := range totalCards {
		sum += v
	}

	return fmt.Sprint(sum)
}

func calculateMatches(input string) (matches int) {
	cardData := strings.Split(input, ": ")
	i, _ := strings.CutPrefix(cardData[0], "Card ")
	numbers := strings.Split(cardData[1], " | ")

	card := Card{
		Number:  MustParse(strings.TrimSpace(i)),
		Winning: parseNumbers(numbers[0]),
		Actual:  parseNumbers(numbers[1]),
	}

	return card.Matches()
}
