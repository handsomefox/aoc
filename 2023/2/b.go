package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))
	games := make([]Game, 0)

	for sc.Scan() {
		txt := sc.Text()
		game := calculateGame(txt)
		games = append(games, game)
	}

	// What is the sum of the power of these sets?
	sum := 0
	for _, v := range games {
		sum += v.Power
	}

	return fmt.Sprint(sum)
}
