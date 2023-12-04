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
		total += calculatePoints(sc.Text())
	}

	return fmt.Sprint(total)
}

func calculatePoints(input string) int {
	cardData := strings.Split(input, ": ")
	i, _ := strings.CutPrefix(cardData[0], "Card ")
	numbers := strings.Split(cardData[1], " | ")

	card := Card{
		Number:  MustParse(strings.TrimSpace(i)),
		Winning: parseNumbers(numbers[0]),
		Actual:  parseNumbers(numbers[1]),
	}

	return card.Points()
}
