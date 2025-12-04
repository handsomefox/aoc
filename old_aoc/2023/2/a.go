package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveA(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))
	games := make([]Game, 0)

	for sc.Scan() {
		txt := sc.Text()
		game := calculateGame(txt)
		games = append(games, game)
	}

	// What is the sum of the IDs of those games?
	sum := 0
	for _, v := range games {
		if v.Possible {
			sum += v.ID
		}
	}

	return fmt.Sprint(sum)
}
