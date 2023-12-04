package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func pointsAndMatches(winning, actual []string) (points, matches int) {
	matched := 0
	for _, num := range actual {
		if slices.Contains(winning, num) {
			matched++
		}
	}
	return int(math.Pow(2, float64(matched-1))), matched
}

func parseCard(s string) (points, matches int) {
	numbers := strings.Split(strings.Split(s, ": ")[1], " | ")
	winning := strings.Fields(numbers[0])
	actual := strings.Fields(numbers[1])

	return pointsAndMatches(winning, actual)
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
