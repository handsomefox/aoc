package main

import (
	"fmt"
	"strings"
)

func SolveA(input string) string {
	split := strings.Split(input, "\n")
	time := ParseFieldsA(split[0], "Time: ")
	distance := ParseFieldsA(split[1], "Distance: ")

	totalWins := 1
	for i := range time {
		totalWins *= AmountOfWinningNumbers(time[i], distance[i])
	}

	return fmt.Sprint(totalWins)
}

func ParseFieldsA(s string, prefix string) []int {
	strs := strings.Fields(strings.Split(s, prefix)[1])
	return IntSlice(strs)
}
