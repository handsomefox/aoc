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

func calculatePoints(s string) int {
	points, _ := parseCard(s)
	return points
}
