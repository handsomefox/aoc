package main

import (
	"fmt"
	"strings"
)

func SolveB(input string) string {
	split := strings.Split(input, "\n")
	time := ParseFieldsB(split[0], "Time: ")
	distance := ParseFieldsB(split[1], "Distance: ")

	wins := AmountOfWinningNumbers(time, distance)

	return fmt.Sprint(wins)
}

func ParseFieldsB(s string, prefix string) int {
	nums := strings.Split(s, prefix)[1]
	nums = strings.ReplaceAll(nums, " ", "")
	return MustParse(nums)
}
