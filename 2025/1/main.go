package main

import (
	"fmt"

	"aoc/shared"
)

const (
	year = 2025
	day  = 1
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string { return solve(input, false) }
func PartB(input string) string { return solve(input, true) }

func solve(input string, countIntermediate bool) string {
	dial, zeroes := 50, 0
	for line := range shared.Lines(input) {
		move, value := line[0], shared.MustParseInt(line[1:])
		for i := range value {
			last := i == value-1
			dial = stepDial(dial, move)
			if countIntermediate && dial == 0 && !last {
				zeroes++
			}
		}
		if dial == 0 {
			zeroes++
		}
	}
	return shared.String(zeroes)
}

func stepDial(dial int, move byte) int {
	delta := map[byte]int{'L': -1, 'R': 1}[move]
	return (dial + delta + 100) % 100
}
