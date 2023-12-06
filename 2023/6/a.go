package main

import (
	"fmt"
	"strings"
)

func SolveA(input string) string {
	split := strings.Split(input, "\n")
	time := strings.Fields(strings.Split(split[0], "Time: ")[1])
	distance := strings.Fields(strings.Split(split[1], "Distance: ")[1])

	totalWins := 1
	for i := range time {
		wins := len(WinningNums(MustParse(time[i]), MustParse(distance[i])))
		totalWins *= wins
	}

	return fmt.Sprint(totalWins)
}
