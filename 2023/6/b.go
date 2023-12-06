package main

import (
	"fmt"
	"strings"
)

func SolveB(input string) string {
	split := strings.Split(input, "\n")
	time := strings.ReplaceAll(strings.Split(split[0], "Time: ")[1], " ", "")
	distance := strings.ReplaceAll(strings.Split(split[1], "Distance: ")[1], " ", "")

	wins := len(WinningNums(MustParse(time), MustParse(distance)))

	return fmt.Sprint(wins)
}
